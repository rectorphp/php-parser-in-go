// Package classresolver resolves PHP class hierarchies and method signatures
// into a registry keyed by fully qualified class name.
//
// It builds on the namespace resolver, so class declarations and their parents
// are keyed by their fully qualified name — a class referenced through a `use`
// alias or a relative name resolves to the same key as its declaration. This
// lets a consumer look up a method declared on a class or any of its ancestors,
// even across files, by merging the registry of every file first.
//
// When two declarations share a fully qualified name (for example two global
// classes with the same short name in unrelated scripts), the name is marked
// ambiguous and lookups against it fail, so callers can stay on the safe side.
package classresolver

import (
	"strings"
	"sync"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/nsresolver"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/traverser"
)

// classData holds the resolved information for a single class name.
type classData struct {
	parent      string                          // fully qualified parent name, empty when none
	interfaces  []string                        // fully qualified names implemented, or extended by an interface
	isInterface bool                            // set for an interface declaration, which is no class to extend
	methods     map[string][]string             // method name -> declared parameter names (with leading $)
	methodNodes map[string]*ast.StmtClassMethod // method name -> declaring node, for body inspection
	visibility  map[string]string               // method name -> "public"/"protected"/"private"
	traits      []string                        // fully qualified names of the traits used here
	properties  map[string]bool                 // property names declared here (no leading $), promotion included
	ambiguous   bool                            // set when the name was declared more than once
}

// Registry holds resolved class signatures and hierarchy, keyed by fully
// qualified class name, gathered from one or more ASTs. It also records which
// names appear as the target of a `new` expression across the collected files.
type Registry struct {
	classes      map[string]*classData
	instantiated map[string]bool

	// extendedMutex guards the memoized ExtendedClassNames set and the
	// parent-to-children index, which callers ask for once per file while the
	// registry itself stops changing after the collect pass, and the
	// anonymousParents set feeding both.
	extendedMutex    sync.Mutex
	extendedNames    map[string]bool
	childrenByParent map[string][]string
	anonymousParents map[string]bool
}

func NewRegistry() *Registry {
	return &Registry{
		classes:          make(map[string]*classData),
		instantiated:     make(map[string]bool),
		anonymousParents: make(map[string]bool),
	}
}

// Collect records every class declared under root. Call it for each file whose
// classes should be resolvable; later lookups see the union. Variadic methods
// are omitted so calls to them are never matched.
func (registry *Registry) Collect(root ast.Vertex) {
	registry.CollectWithResolvedNames(root, ResolveNames(root))
}

// CollectWithResolvedNames is Collect for a caller that already holds root's
// resolved names, letting several collectors over one file share a single walk
// of the namespace resolver rather than each paying for its own.
func (registry *Registry) CollectWithResolvedNames(root ast.Vertex, resolvedNames map[ast.Vertex]string) {
	collector := &classCollector{
		Null:          &visitor.Null{},
		registry:      registry,
		resolvedNames: resolvedNames,
	}
	traverser.NewTraverser(collector).Traverse(root)
}

// CollectTraitsWithResolvedNames records only the trait declarations under root,
// leaving its classes and interfaces out of the registry. It lets a caller fold
// in traits from files it will never rewrite — a vendored trait, say — so a class
// that uses one can be checked against the trait's real properties instead of
// treated as unresolved, without the file's classes affecting class-level lookups
// such as uniqueness or the extended set.
func (registry *Registry) CollectTraitsWithResolvedNames(root ast.Vertex, resolvedNames map[ast.Vertex]string) {
	collector := &classCollector{
		Null:          &visitor.Null{},
		registry:      registry,
		resolvedNames: resolvedNames,
		traitsOnly:    true,
	}
	traverser.NewTraverser(collector).Traverse(root)
}

// HasUnresolvedTraitReference reports whether any collected class or trait uses a
// trait whose declaration is not itself in the registry. It lets a caller decide
// whether folding in traits from an unscanned tree (vendor/) could resolve
// anything, so that extra scan runs only when it might help. An ambiguous trait
// name counts as resolved here: a second declaration would not disambiguate it,
// so scanning for one is fruitless.
func (registry *Registry) HasUnresolvedTraitReference() bool {
	for _, data := range registry.classes {
		for _, traitFQN := range data.traits {
			if _, ok := registry.classes[traitFQN]; !ok {
				return true
			}
		}
	}
	return false
}

// LookupMethodParameterCount resolves methodName starting at classFQN and
// walking up the parent chain, returning the declared parameter count. It fails
// (false) when a class in the chain is unknown, ambiguous, or method not found.
// The seen set guards against cyclic extends declarations.
func (registry *Registry) LookupMethodParameterCount(classFQN, methodName string) (int, bool) {
	parameterNames, ok := registry.LookupMethodParameterNames(classFQN, methodName)
	if !ok {
		return 0, false
	}
	return len(parameterNames), true
}

// LookupMethodParameterNames resolves methodName starting at classFQN and
// walking up the parent chain, returning the declared parameter names (each with
// its leading $). It fails (false) when a class in the chain is unknown,
// ambiguous, or the method is not found. The seen set guards against cyclic
// extends declarations.
func (registry *Registry) LookupMethodParameterNames(classFQN, methodName string) ([]string, bool) {
	seen := make(map[string]bool)

	for classFQN != "" && !seen[classFQN] {
		seen[classFQN] = true

		data, ok := registry.classes[classFQN]
		if !ok || data.ambiguous {
			return nil, false
		}

		if parameterNames, ok := data.methods[methodName]; ok {
			return parameterNames, true
		}

		classFQN = data.parent
	}

	return nil, false
}

// TraitDeclaresProperty reports whether any trait used by classFQN — directly or
// through a trait that itself uses another — declares a property named
// propertyName (given without its leading $). It lets a caller avoid marking a
// class property `readonly` when a trait contributes a same-named property, which
// PHP rejects unless both sides agree. The check is best effort: a trait outside
// the collected files is invisible here, so absence is not proof of absence.
func (registry *Registry) TraitDeclaresProperty(classFQN, propertyName string) bool {
	seen := make(map[string]bool)

	var usedByDeclares func(fqn string) bool
	usedByDeclares = func(fqn string) bool {
		data, ok := registry.classes[fqn]
		if !ok || data.ambiguous {
			return false
		}
		for _, traitFQN := range data.traits {
			if seen[traitFQN] {
				continue
			}
			seen[traitFQN] = true

			traitData, ok := registry.classes[traitFQN]
			if ok && !traitData.ambiguous && traitData.properties[propertyName] {
				return true
			}
			if usedByDeclares(traitFQN) {
				return true
			}
		}
		return false
	}

	return usedByDeclares(classFQN)
}

// UsesTraitDeclaringProperty reports whether classFQN draws in a trait — directly
// or through a trait that itself uses another — that declares any property at all.
// A readonly class flattens every trait property into itself, and PHP rejects the
// class unless all of them are readonly; the registry records a trait property's
// name but not its readonly modifier, so a caller weighing the class-level
// `readonly` modifier treats any trait property as reason enough to hold off. The
// check is best effort: a trait outside the collected files is invisible here, so
// UsesUnresolvedTrait covers that gap.
func (registry *Registry) UsesTraitDeclaringProperty(classFQN string) bool {
	seen := make(map[string]bool)

	var walk func(fqn string) bool
	walk = func(fqn string) bool {
		data, ok := registry.classes[fqn]
		if !ok || data.ambiguous {
			return false
		}
		for _, traitFQN := range data.traits {
			if seen[traitFQN] {
				continue
			}
			seen[traitFQN] = true

			traitData, ok := registry.classes[traitFQN]
			if ok && !traitData.ambiguous && len(traitData.properties) > 0 {
				return true
			}
			if walk(traitFQN) {
				return true
			}
		}
		return false
	}

	return walk(classFQN)
}

// UsesUnresolvedTrait reports whether classFQN draws in a trait — directly, or
// through a trait that itself uses another — whose declaration is absent from the
// collected files. A `use` records the trait's name even when its body was never
// scanned, so the name is known while its members are not. When that happens a
// property-level check such as TraitDeclaresProperty cannot see what the trait
// contributes, and its false answer is no proof of absence; a caller that must
// stay safe treats an unresolved trait as possibly declaring any member. A class
// that is itself unknown or ambiguous is unresolved too, so it reports true.
func (registry *Registry) UsesUnresolvedTrait(classFQN string) bool {
	seen := make(map[string]bool)

	var walk func(fqn string) bool
	walk = func(fqn string) bool {
		data, ok := registry.classes[fqn]
		if !ok || data.ambiguous {
			return true
		}
		for _, traitFQN := range data.traits {
			if seen[traitFQN] {
				continue
			}
			seen[traitFQN] = true
			if walk(traitFQN) {
				return true
			}
		}
		return false
	}

	return walk(classFQN)
}

// InheritsVisibleMethod resolves methodName starting at startFQN and walking up
// the parent chain, reporting whether an ancestor declares it protected or
// public — the visibilities that a subclass override may not reduce below. Pass
// the parent of the class in question as startFQN, so the class's own
// declaration is excluded.
//
// found is true once such a declaration is met; a private ancestor declaration
// imposes no constraint, so the walk continues past it. resolvable is false when
// a class in the chain is unknown or ambiguous, meaning the ancestors cannot be
// proven — a caller that must stay safe treats that like a constraint. The seen
// set guards against cyclic extends declarations.
func (registry *Registry) InheritsVisibleMethod(startFQN, methodName string) (found bool, resolvable bool) {
	seen := make(map[string]bool)

	for startFQN != "" && !seen[startFQN] {
		seen[startFQN] = true

		data, ok := registry.classes[startFQN]
		if !ok || data.ambiguous {
			return false, false
		}

		if visibility := data.visibility[methodName]; visibility == "protected" || visibility == "public" {
			return true, true
		}

		startFQN = data.parent
	}

	return false, true
}

// ExtendedClassNames returns the set of fully qualified names that at least one
// collected class extends, anonymous classes included. A name in this set has a
// child, so it is not a leaf of the class hierarchy. Names that could not be
// resolved contribute nothing. The result is memoized and shared between
// callers, so treat it as read-only; collecting another file discards the memo.
func (registry *Registry) ExtendedClassNames() map[string]bool {
	registry.extendedMutex.Lock()
	defer registry.extendedMutex.Unlock()

	if registry.extendedNames != nil {
		return registry.extendedNames
	}

	extended := make(map[string]bool)
	for _, data := range registry.classes {
		// an interface extends interfaces, which says nothing about a class
		// having a child
		if data.isInterface {
			continue
		}
		if data.parent != "" {
			extended[data.parent] = true
		}
	}
	for parentFQN := range registry.anonymousParents {
		extended[parentFQN] = true
	}
	registry.extendedNames = extended

	return extended
}

// DescendantDeclaresMethod reports whether any collected class below classFQN
// declares a method named methodName — the condition under which narrowing the
// method on classFQN would break an override.
//
// The walk goes down the hierarchy, so it answers a question InheritsVisibleMethod
// cannot: that one looks up at the ancestors of one class, this one looks down at
// every class that inherits from it.
//
// resolvable is false when the answer cannot be proven: classFQN is unknown or
// ambiguous, a descendant is ambiguous, or an anonymous class extends classFQN or
// one of its descendants — an anonymous class has no name to key its methods by,
// so its overrides stay invisible. A caller that must stay safe treats an
// unresolvable answer like a declared override. The seen set guards against
// cyclic extends declarations.
func (registry *Registry) DescendantDeclaresMethod(classFQN, methodName string) (declared bool, resolvable bool) {
	data, ok := registry.classes[classFQN]
	if !ok || data.ambiguous {
		return false, false
	}

	children := registry.childNames()

	seen := map[string]bool{classFQN: true}
	pending := []string{classFQN}
	for len(pending) > 0 {
		currentFQN := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if registry.hasAnonymousChild(currentFQN) {
			return false, false
		}

		for _, childFQN := range children[currentFQN] {
			if seen[childFQN] {
				continue
			}
			seen[childFQN] = true

			childData, ok := registry.classes[childFQN]
			if !ok || childData.ambiguous {
				return false, false
			}

			if _, ok := childData.methods[methodName]; ok {
				return true, true
			}

			pending = append(pending, childFQN)
		}
	}

	return false, true
}

// DescendantOverrides returns the declaring node of every collected class below
// classFQN that declares a method named methodName. It answers the finer question
// DescendantDeclaresMethod cannot: not merely whether an override exists, but which
// nodes they are, so a caller can inspect each body — an override that returns a
// value blocks narrowing the ancestor to `void`, one with a value-less body does
// not, since the same run types it `void` too and the pair stays compatible.
//
// resolvable is false under the same unprovable conditions as
// DescendantDeclaresMethod: classFQN is unknown or ambiguous, a descendant is
// ambiguous, or an anonymous class extends classFQN or one of its descendants. A
// caller that must stay safe treats an unresolvable answer like a value-returning
// override. The seen set guards against cyclic extends declarations.
func (registry *Registry) DescendantOverrides(classFQN, methodName string) (overrides []*ast.StmtClassMethod, resolvable bool) {
	data, ok := registry.classes[classFQN]
	if !ok || data.ambiguous {
		return nil, false
	}

	children := registry.childNames()

	seen := map[string]bool{classFQN: true}
	pending := []string{classFQN}
	for len(pending) > 0 {
		currentFQN := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if registry.hasAnonymousChild(currentFQN) {
			return nil, false
		}

		for _, childFQN := range children[currentFQN] {
			if seen[childFQN] {
				continue
			}
			seen[childFQN] = true

			childData, ok := registry.classes[childFQN]
			if !ok || childData.ambiguous {
				return nil, false
			}

			if method, ok := childData.methodNodes[methodName]; ok {
				overrides = append(overrides, method)
			}

			pending = append(pending, childFQN)
		}
	}

	return overrides, true
}

// childNames indexes the collected classes by their parent, so a walk down the
// hierarchy costs one pass instead of one per level. It is memoized alongside
// ExtendedClassNames and shares its invalidation, so treat the result as
// read-only.
func (registry *Registry) childNames() map[string][]string {
	registry.extendedMutex.Lock()
	defer registry.extendedMutex.Unlock()

	if registry.childrenByParent != nil {
		return registry.childrenByParent
	}

	children := make(map[string][]string)
	for classFQN, data := range registry.classes {
		if data.parent != "" {
			children[data.parent] = append(children[data.parent], classFQN)
		}
	}
	registry.childrenByParent = children

	return children
}

// hasAnonymousChild reports whether an anonymous class extends classFQN.
func (registry *Registry) hasAnonymousChild(classFQN string) bool {
	registry.extendedMutex.Lock()
	defer registry.extendedMutex.Unlock()

	return registry.anonymousParents[classFQN]
}

// IsSubtypeOf reports whether classFQN is ancestorFQN itself, extends it, or
// implements it — directly or through any number of steps. Both names are fully
// qualified and carry no leading backslash.
//
// The walk climbs only through declarations the registry holds, so a name it
// never collected is a leaf: a class extending a vendor class is a subtype of
// that very name, but not of what the vendor class extends in turn. That makes
// the answer a lower bound — true is certain, false only means "not provable
// from the collected files" — which is what a caller reporting on the answer
// wants. An ambiguous name stops its branch for the same reason.
func (registry *Registry) IsSubtypeOf(classFQN, ancestorFQN string) bool {
	if classFQN == "" || ancestorFQN == "" {
		return false
	}
	return registry.isSubtypeOf(classFQN, ancestorFQN, make(map[string]bool))
}

// isSubtypeOf is IsSubtypeOf with the set of names already visited, guarding
// against a cyclic extends or implements declaration.
func (registry *Registry) isSubtypeOf(classFQN, ancestorFQN string, seen map[string]bool) bool {
	if classFQN == ancestorFQN {
		return true
	}
	if seen[classFQN] {
		return false
	}
	seen[classFQN] = true

	data, ok := registry.classes[classFQN]
	if !ok || data.ambiguous {
		return false
	}
	if data.parent != "" && registry.isSubtypeOf(data.parent, ancestorFQN, seen) {
		return true
	}
	for _, interfaceFQN := range data.interfaces {
		if registry.isSubtypeOf(interfaceFQN, ancestorFQN, seen) {
			return true
		}
	}
	return false
}

// InterfacesOf returns the fully qualified names of every interface classFQN
// implements — directly, through a parent class, or through interface
// inheritance. The walk climbs only declarations the registry holds, so an
// interface from an unscanned vendor tree contributes its own name but not what
// it extends in turn; a name that resolves to more than one declaration stops
// its branch. The order is unspecified, so treat the result as a set.
func (registry *Registry) InterfacesOf(classFQN string) []string {
	interfaces := make(map[string]bool)
	seen := make(map[string]bool)

	var collectFrom func(fqn string)
	collectFrom = func(fqn string) {
		if fqn == "" || seen[fqn] {
			return
		}
		seen[fqn] = true

		data, ok := registry.classes[fqn]
		if !ok || data.ambiguous {
			return
		}
		for _, interfaceFQN := range data.interfaces {
			interfaces[interfaceFQN] = true
			collectFrom(interfaceFQN) // pull in the interfaces it extends
		}
		collectFrom(data.parent) // inherit the parent class's interfaces
	}
	collectFrom(classFQN)

	result := make([]string, 0, len(interfaces))
	for interfaceFQN := range interfaces {
		result = append(result, interfaceFQN)
	}
	return result
}

// IsInterface reports whether the name was collected as an interface
// declaration. A name the run never collected reports false, so a caller asking
// "does this interface exist" is told no rather than maybe.
func (registry *Registry) IsInterface(interfaceFQN string) bool {
	data, ok := registry.classes[interfaceFQN]
	return ok && !data.ambiguous && data.isInterface
}

// IsUnique reports whether classFQN was declared exactly once, so it is known
// and not ambiguous. Callers that mutate a class use this to stay off names
// that resolve to more than one declaration.
func (registry *Registry) IsUnique(classFQN string) bool {
	data, ok := registry.classes[classFQN]
	return ok && !data.ambiguous
}

// IsInstantiated reports whether classFQN appears as the target of a `new`
// expression anywhere in the collected files. A `new self` or `new static`
// written inside a class counts that class as instantiated. Instantiation
// through reflection, a container, or deserialization is invisible here.
func (registry *Registry) IsInstantiated(classFQN string) bool {
	return registry.instantiated[classFQN]
}

// addAnonymousParent records a name extended by an anonymous class. Such a
// class has no name to key the classes map by, yet its parent still has a
// child, so it must not be treated as a leaf.
func (registry *Registry) addAnonymousParent(parentFQN string) {
	registry.extendedMutex.Lock()
	defer registry.extendedMutex.Unlock()

	registry.extendedNames = nil
	registry.childrenByParent = nil
	registry.anonymousParents[parentFQN] = true
}

func (registry *Registry) add(classFQN string, data *classData) {
	registry.extendedMutex.Lock()
	registry.extendedNames = nil
	registry.childrenByParent = nil
	registry.extendedMutex.Unlock()

	if existing, ok := registry.classes[classFQN]; ok {
		existing.ambiguous = true
		return
	}
	registry.classes[classFQN] = data
}

// ResolveNames runs the namespace resolver over root and returns its map of AST
// node to fully qualified name, covering class declarations and the names they
// extend.
func ResolveNames(root ast.Vertex) map[ast.Vertex]string {
	resolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(resolver).Traverse(root)
	return resolver.ResolvedNames
}

// classCollector records each class's parent and non-variadic method signatures.
// When traitsOnly is set it records trait declarations alone, leaving classes and
// interfaces out of the registry.
type classCollector struct {
	*visitor.Null
	registry      *Registry
	resolvedNames map[ast.Vertex]string
	traitsOnly    bool
}

func (collector *classCollector) StmtClass(node *ast.StmtClass) {
	if collector.traitsOnly {
		return
	}
	parentFQN := ""
	if node.Extends != nil {
		if resolvedParentFQN, ok := collector.resolvedNames[node.Extends]; ok {
			parentFQN = resolvedParentFQN
		}
	}

	classFQN, ok := collector.resolvedNames[node]
	if !ok {
		// An anonymous class carries no name, so it never enters the classes
		// map — only what it extends is worth keeping.
		if parentFQN != "" {
			collector.registry.addAnonymousParent(parentFQN)
		}
		return
	}

	if hasSelfInstantiation(node, collector.resolvedNames) {
		collector.registry.instantiated[classFQN] = true
	}

	data := &classData{
		methods:     make(map[string][]string),
		methodNodes: make(map[string]*ast.StmtClassMethod),
		visibility:  make(map[string]string),
		parent:      parentFQN,
		interfaces:  collector.resolvedList(node.Implements),
		traits:      collector.traitFQNs(node.Stmts),
		properties:  propertyNamesFromStmts(node.Stmts),
	}

	for _, stmt := range node.Stmts {
		method, ok := stmt.(*ast.StmtClassMethod)
		if !ok {
			continue
		}

		methodName, ok := identifierName(method.Name)
		if !ok {
			continue
		}

		// Visibility is recorded for every method, variadic included, so a
		// caller can tell whether an inherited method forces a subclass to keep
		// a wider visibility than private.
		data.visibility[methodName] = methodVisibility(method)

		// The declaring node is kept for every method, variadic included, so a
		// caller can inspect an override's body — whether it returns a value —
		// which decides if narrowing an ancestor to `void` is safe.
		data.methodNodes[methodName] = method

		// Parameter names are matched against call sites, where a variadic
		// swallows any argument count, so a variadic method is left unrecorded.
		if isVariadic(method.Params) {
			continue
		}

		data.methods[methodName] = parameterNames(method.Params)
	}

	collector.registry.add(classFQN, data)
}

// StmtInterface records an interface declaration and what it extends, so a
// class implementing it is a subtype of every interface above it too. An
// interface declares no method bodies to look up, so only the hierarchy is kept.
func (collector *classCollector) StmtInterface(node *ast.StmtInterface) {
	if collector.traitsOnly {
		return
	}
	interfaceFQN, ok := collector.resolvedNames[node]
	if !ok {
		return
	}
	collector.registry.add(interfaceFQN, &classData{
		isInterface: true,
		interfaces:  collector.resolvedList(node.Extends),
		methods:     make(map[string][]string),
		visibility:  make(map[string]string),
	})
}

// StmtTrait records a trait declaration: its property names, so a class that
// uses it can tell whether marking one of its own properties `readonly` would
// clash with an inherited one, and the traits it uses in turn, so that check
// reaches nested traits.
func (collector *classCollector) StmtTrait(node *ast.StmtTrait) {
	traitFQN, ok := collector.resolvedNames[node]
	if !ok {
		return
	}
	collector.registry.add(traitFQN, &classData{
		methods:    make(map[string][]string),
		visibility: make(map[string]string),
		traits:     collector.traitFQNs(node.Stmts),
		properties: propertyNamesFromStmts(node.Stmts),
	})
}

// traitFQNs returns the fully qualified names of every trait a `use` statement in
// stmts pulls in, dropping any that did not resolve to a name.
func (collector *classCollector) traitFQNs(stmts []ast.Vertex) []string {
	var traitFQNs []string
	for _, stmt := range stmts {
		traitUse, ok := stmt.(*ast.StmtTraitUse)
		if !ok {
			continue
		}
		for _, traitName := range traitUse.Traits {
			if traitFQN, ok := collector.resolvedNames[traitName]; ok {
				traitFQNs = append(traitFQNs, traitFQN)
			}
		}
	}
	return traitFQNs
}

// propertyNamesFromStmts collects the property names a class or trait body
// declares, without the leading $, counting both plain property declarations and
// constructor-promoted parameters. Visibility is not filtered: a same-named
// property of any visibility is enough to clash.
func propertyNamesFromStmts(stmts []ast.Vertex) map[string]bool {
	names := make(map[string]bool)
	for _, stmt := range stmts {
		switch typed := stmt.(type) {
		case *ast.StmtPropertyList:
			for _, property := range typed.Props {
				declared, ok := property.(*ast.StmtProperty)
				if !ok {
					continue
				}
				if name, ok := propertyVariableName(declared.Var); ok {
					names[name] = true
				}
			}
		case *ast.StmtClassMethod:
			methodName, ok := identifierName(typed.Name)
			if !ok || !strings.EqualFold(methodName, "__construct") {
				continue
			}
			for _, parameter := range typed.Params {
				promoted, ok := parameter.(*ast.Parameter)
				if !ok || len(promoted.Modifiers) == 0 {
					continue // a parameter with no modifier promotes nothing
				}
				if name, ok := propertyVariableName(promoted.Var); ok {
					names[name] = true
				}
			}
		}
	}
	if len(names) == 0 {
		return nil
	}
	return names
}

// propertyVariableName returns a property variable's name without its $, so it
// compares equal to the names the readonly rule works with.
func propertyVariableName(node ast.Vertex) (string, bool) {
	variable, ok := node.(*ast.ExprVariable)
	if !ok {
		return "", false
	}
	name, ok := identifierName(variable.Name)
	if !ok {
		return "", false
	}
	return strings.TrimPrefix(name, "$"), true
}

// resolvedList returns the fully qualified name of each name node that resolved
// to one, dropping the rest.
func (collector *classCollector) resolvedList(nameNodes []ast.Vertex) []string {
	names := make([]string, 0, len(nameNodes))
	for _, nameNode := range nameNodes {
		if resolved, ok := collector.resolvedNames[nameNode]; ok {
			names = append(names, resolved)
		}
	}
	if len(names) == 0 {
		return nil
	}
	return names
}

// ExprNew records the fully qualified name of every class instantiated with
// `new`. A `new self`/`new static` resolves to the literal "self"/"static" here;
// those are attributed to the enclosing class by hasSelfInstantiation instead.
func (collector *classCollector) ExprNew(node *ast.ExprNew) {
	if instantiatedFQN, ok := collector.resolvedNames[node.Class]; ok {
		collector.registry.instantiated[instantiatedFQN] = true
	}
}

// hasSelfInstantiation reports whether the class body instantiates itself with
// `new self` or `new static`, so the class must not be treated as never
// instantiated. Nested classes contribute too, which only makes the answer more
// conservative.
func hasSelfInstantiation(classNode *ast.StmtClass, resolvedNames map[ast.Vertex]string) bool {
	detector := &selfInstantiationDetector{Null: &visitor.Null{}, resolvedNames: resolvedNames}
	traverser.NewTraverser(detector).Traverse(classNode)
	return detector.found
}

// selfInstantiationDetector flags a `new self` or `new static` under the node it
// traverses.
type selfInstantiationDetector struct {
	*visitor.Null
	resolvedNames map[ast.Vertex]string
	found         bool
}

func (detector *selfInstantiationDetector) ExprNew(node *ast.ExprNew) {
	// `new self` reaches here as a resolved *ast.Name, while the `static`
	// keyword in `new static` is an *ast.Identifier the namespace resolver
	// leaves untouched, so both forms are checked.
	if identifier, ok := node.Class.(*ast.Identifier); ok {
		value := strings.ToLower(string(identifier.Value))
		if value == "self" || value == "static" {
			detector.found = true
		}
		return
	}
	if name, ok := detector.resolvedNames[node.Class]; ok && (name == "self" || name == "static") {
		detector.found = true
	}
}

// parameterNames returns the variable name of each parameter, keeping declared
// order. A parameter whose variable name cannot be resolved contributes an empty
// string, so the slice length always matches the declared parameter count.
func parameterNames(parameters []ast.Vertex) []string {
	names := make([]string, 0, len(parameters))
	for _, parameter := range parameters {
		name := ""
		if typedParameter, ok := parameter.(*ast.Parameter); ok {
			if variable, ok := typedParameter.Var.(*ast.ExprVariable); ok {
				if variableName, ok := identifierName(variable.Name); ok {
					name = variableName
				}
			}
		}
		names = append(names, name)
	}
	return names
}

// methodVisibility returns a method's declared visibility, defaulting to
// "public" when no visibility modifier is present, as PHP does.
func methodVisibility(method *ast.StmtClassMethod) string {
	for _, modifier := range method.Modifiers {
		identifier, ok := modifier.(*ast.Identifier)
		if !ok {
			continue
		}
		switch strings.ToLower(string(identifier.Value)) {
		case "public":
			return "public"
		case "protected":
			return "protected"
		case "private":
			return "private"
		}
	}
	return "public"
}

func identifierName(node ast.Vertex) (string, bool) {
	identifier, ok := node.(*ast.Identifier)
	if !ok {
		return "", false
	}
	return string(identifier.Value), true
}

func isVariadic(parameters []ast.Vertex) bool {
	for _, parameter := range parameters {
		if typedParameter, ok := parameter.(*ast.Parameter); ok && typedParameter.VariadicTkn != nil {
			return true
		}
	}
	return false
}
