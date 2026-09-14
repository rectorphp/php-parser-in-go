// Package visitor contains walker.visitor implementations
package nsresolver

import (
	"errors"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
	"strings"
)

// NamespaceResolver visitor
type NamespaceResolver struct {
	visitor.Null
	Namespace     *Namespace
	ResolvedNames map[ast.Vertex]string

	goDeep bool
}

// NewNamespaceResolver NamespaceResolver type constructor
func NewNamespaceResolver() *NamespaceResolver {
	return &NamespaceResolver{
		Namespace:     NewNamespace(""),
		ResolvedNames: map[ast.Vertex]string{},
		goDeep:        true,
	}
}

func (namespaceResolver *NamespaceResolver) EnterNode(node ast.Vertex) bool {
	node.Accept(namespaceResolver)

	if !namespaceResolver.goDeep {
		namespaceResolver.goDeep = true
		return false
	}

	return true
}

func (namespaceResolver *NamespaceResolver) StmtNamespace(node *ast.StmtNamespace) {
	if node.Name == nil {
		namespaceResolver.Namespace = NewNamespace("")
	} else {
		nameParts := node.Name.(*ast.Name).Parts
		namespaceResolver.Namespace = NewNamespace(concatNameParts(nameParts))
	}
}

func (namespaceResolver *NamespaceResolver) StmtUse(node *ast.StmtUseList) {
	useType := ""
	if node.Type != nil {
		useType = string(node.Type.(*ast.Identifier).Value)
	}

	for _, use := range node.Uses {
		namespaceResolver.AddAlias(useType, use, nil)
	}

	namespaceResolver.goDeep = false
}

func (namespaceResolver *NamespaceResolver) StmtGroupUse(node *ast.StmtGroupUseList) {
	useType := ""
	if node.Type != nil {
		useType = string(node.Type.(*ast.Identifier).Value)
	}

	for _, use := range node.Uses {
		namespaceResolver.AddAlias(useType, use, node.Prefix.(*ast.Name).Parts)
	}

	namespaceResolver.goDeep = false
}

func (namespaceResolver *NamespaceResolver) StmtClass(node *ast.StmtClass) {
	if node.Extends != nil {
		namespaceResolver.ResolveName(node.Extends, "")
	}

	if node.Implements != nil {
		for _, interfaceName := range node.Implements {
			namespaceResolver.ResolveName(interfaceName, "")
		}
	}

	if node.Name != nil {
		namespaceResolver.AddNamespacedName(node, string(node.Name.(*ast.Identifier).Value))
	}
}

func (namespaceResolver *NamespaceResolver) StmtEnum(node *ast.StmtEnum) {
	if node.Implements != nil {
		for _, interfaceName := range node.Implements {
			namespaceResolver.ResolveName(interfaceName, "")
		}
	}

	namespaceResolver.AddNamespacedName(node, string(node.Name.(*ast.Identifier).Value))
}

func (namespaceResolver *NamespaceResolver) StmtInterface(node *ast.StmtInterface) {
	if node.Extends != nil {
		for _, interfaceName := range node.Extends {
			namespaceResolver.ResolveName(interfaceName, "")
		}
	}

	namespaceResolver.AddNamespacedName(node, string(node.Name.(*ast.Identifier).Value))
}

func (namespaceResolver *NamespaceResolver) StmtTrait(node *ast.StmtTrait) {
	namespaceResolver.AddNamespacedName(node, string(node.Name.(*ast.Identifier).Value))
}

func (namespaceResolver *NamespaceResolver) StmtFunction(node *ast.StmtFunction) {
	namespaceResolver.AddNamespacedName(node, string(node.Name.(*ast.Identifier).Value))

	for _, parameter := range node.Params {
		namespaceResolver.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if node.ReturnType != nil {
		namespaceResolver.ResolveType(node.ReturnType)
	}
}

func (namespaceResolver *NamespaceResolver) StmtClassMethod(node *ast.StmtClassMethod) {
	for _, parameter := range node.Params {
		namespaceResolver.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if node.ReturnType != nil {
		namespaceResolver.ResolveType(node.ReturnType)
	}
}

func (namespaceResolver *NamespaceResolver) ExprClosure(node *ast.ExprClosure) {
	for _, parameter := range node.Params {
		namespaceResolver.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if node.ReturnType != nil {
		namespaceResolver.ResolveType(node.ReturnType)
	}
}

func (namespaceResolver *NamespaceResolver) ExprArrowFunction(node *ast.ExprArrowFunction) {
	for _, parameter := range node.Params {
		namespaceResolver.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if node.ReturnType != nil {
		namespaceResolver.ResolveType(node.ReturnType)
	}
}

func (namespaceResolver *NamespaceResolver) StmtPropertyList(node *ast.StmtPropertyList) {
	if node.Type != nil {
		namespaceResolver.ResolveType(node.Type)
	}
}

func (namespaceResolver *NamespaceResolver) StmtClassConstList(node *ast.StmtClassConstList) {
	if node.Type != nil {
		namespaceResolver.ResolveType(node.Type)
	}
}

func (namespaceResolver *NamespaceResolver) StmtConstList(node *ast.StmtConstList) {
	for _, constant := range node.Consts {
		namespaceResolver.AddNamespacedName(constant, string(constant.(*ast.StmtConstant).Name.(*ast.Identifier).Value))
	}
}

func (namespaceResolver *NamespaceResolver) ExprStaticCall(node *ast.ExprStaticCall) {
	namespaceResolver.ResolveName(node.Class, "")
}

func (namespaceResolver *NamespaceResolver) ExprStaticPropertyFetch(node *ast.ExprStaticPropertyFetch) {
	namespaceResolver.ResolveName(node.Class, "")
}

func (namespaceResolver *NamespaceResolver) ExprClassConstFetch(node *ast.ExprClassConstFetch) {
	namespaceResolver.ResolveName(node.Class, "")
}

func (namespaceResolver *NamespaceResolver) ExprNew(node *ast.ExprNew) {
	namespaceResolver.ResolveName(node.Class, "")
}

func (namespaceResolver *NamespaceResolver) ExprInstanceOf(node *ast.ExprInstanceOf) {
	namespaceResolver.ResolveName(node.Class, "")
}

func (namespaceResolver *NamespaceResolver) StmtCatch(node *ast.StmtCatch) {
	for _, catchType := range node.Types {
		namespaceResolver.ResolveName(catchType, "")
	}
}

func (namespaceResolver *NamespaceResolver) ExprFunctionCall(node *ast.ExprFunctionCall) {
	namespaceResolver.ResolveName(node.Function, "function")
}

func (namespaceResolver *NamespaceResolver) ExprConstFetch(node *ast.ExprConstFetch) {
	namespaceResolver.ResolveName(node.Const, "const")
}

func (namespaceResolver *NamespaceResolver) Attribute(node *ast.Attribute) {
	namespaceResolver.ResolveName(node.Name, "")
}

func (namespaceResolver *NamespaceResolver) StmtTraitUse(node *ast.StmtTraitUse) {
	for _, trait := range node.Traits {
		namespaceResolver.ResolveName(trait, "")
	}

	for _, adaptation := range node.Adaptations {
		switch traitUseAdaptation := adaptation.(type) {
		case *ast.StmtTraitUsePrecedence:
			refTrait := traitUseAdaptation.Trait
			if refTrait != nil {
				namespaceResolver.ResolveName(refTrait, "")
			}
			for _, insteadOf := range traitUseAdaptation.Insteadof {
				namespaceResolver.ResolveName(insteadOf, "")
			}

		case *ast.StmtTraitUseAlias:
			refTrait := traitUseAdaptation.Trait
			if refTrait != nil {
				namespaceResolver.ResolveName(refTrait, "")
			}
		}
	}
}

// LeaveNode is invoked after node process
func (namespaceResolver *NamespaceResolver) LeaveNode(node ast.Vertex) {
	switch stmtNamespace := node.(type) {
	case *ast.StmtNamespace:
		if stmtNamespace.Stmts != nil {
			namespaceResolver.Namespace = NewNamespace("")
		}
	}
}

// AddAlias adds a new alias
func (namespaceResolver *NamespaceResolver) AddAlias(useType string, useNode ast.Vertex, prefix []ast.Vertex) {
	switch use := useNode.(type) {
	case *ast.StmtUse:
		if use.Type != nil {
			useType = string(use.Type.(*ast.Identifier).Value)
		}

		// The imported name is a Name for `use \App\Mailer;`, a NameRelative for
		// the far more common unqualified `use App\Mailer;`, or a
		// NameFullyQualified — accept all three. Anything else carries no parts to
		// alias, so it is skipped rather than panicked on.
		useNameParts := nameNodeParts(use.Use)
		if len(useNameParts) == 0 {
			return
		}

		var alias string
		if use.Alias == nil {
			alias = string(useNameParts[len(useNameParts)-1].(*ast.NamePart).Value)
		} else {
			alias = string(use.Alias.(*ast.Identifier).Value)
		}

		namespaceResolver.Namespace.AddAlias(useType, concatNameParts(prefix, useNameParts), alias)
	}
}

// nameNodeParts returns the name parts of a Name, NameRelative or
// NameFullyQualified node, and nil for anything else.
func nameNodeParts(node ast.Vertex) []ast.Vertex {
	switch name := node.(type) {
	case *ast.Name:
		return name.Parts
	case *ast.NameRelative:
		return name.Parts
	case *ast.NameFullyQualified:
		return name.Parts
	default:
		return nil
	}
}

// AddNamespacedName adds namespaced name by node
func (namespaceResolver *NamespaceResolver) AddNamespacedName(node ast.Vertex, nodeName string) {
	if namespaceResolver.Namespace.Namespace == "" {
		namespaceResolver.ResolvedNames[node] = nodeName
	} else {
		namespaceResolver.ResolvedNames[node] = namespaceResolver.Namespace.Namespace + "\\" + nodeName
	}
}

// ResolveName adds a resolved fully qualified name by node
func (namespaceResolver *NamespaceResolver) ResolveName(nameNode ast.Vertex, aliasType string) {
	resolved, err := namespaceResolver.Namespace.ResolveName(nameNode, aliasType)
	if err == nil {
		namespaceResolver.ResolvedNames[nameNode] = resolved
	}
}

// ResolveType adds a resolved fully qualified type name
func (namespaceResolver *NamespaceResolver) ResolveType(node ast.Vertex) {
	switch typeNode := node.(type) {
	case *ast.Nullable:
		namespaceResolver.ResolveType(typeNode.Expr)
	case *ast.Union:
		for _, memberType := range typeNode.Types {
			namespaceResolver.ResolveType(memberType)
		}
	case *ast.Intersection:
		for _, memberType := range typeNode.Types {
			namespaceResolver.ResolveType(memberType)
		}
	case *ast.Name:
		namespaceResolver.ResolveName(node, "")
	case *ast.NameRelative:
		namespaceResolver.ResolveName(node, "")
	case *ast.NameFullyQualified:
		namespaceResolver.ResolveName(node, "")
	}
}

// Namespace context
type Namespace struct {
	Namespace string
	Aliases   map[string]map[string]string
}

// NewNamespace constructor
func NewNamespace(namespaceName string) *Namespace {
	return &Namespace{
		Namespace: namespaceName,
		Aliases: map[string]map[string]string{
			"":         {},
			"const":    {},
			"function": {},
		},
	}
}

// AddAlias adds a new alias
func (namespace *Namespace) AddAlias(aliasType string, aliasName string, alias string) {
	aliasType = strings.ToLower(aliasType)

	if aliasType == "const" {
		namespace.Aliases[aliasType][alias] = aliasName
	} else {
		namespace.Aliases[aliasType][strings.ToLower(alias)] = aliasName
	}
}

// ResolveName returns a resolved fully qualified name
func (namespace *Namespace) ResolveName(nameNode ast.Vertex, aliasType string) (string, error) {
	switch name := nameNode.(type) {
	case *ast.NameFullyQualified:
		// Fully qualifid name is already resolved
		return concatNameParts(name.Parts), nil

	case *ast.NameRelative:
		if namespace.Namespace == "" {
			return concatNameParts(name.Parts), nil
		}
		return namespace.Namespace + "\\" + concatNameParts(name.Parts), nil

	case *ast.Name:
		if aliasType == "const" && len(name.Parts) == 1 {
			part := strings.ToLower(string(name.Parts[0].(*ast.NamePart).Value))
			if part == "true" || part == "false" || part == "null" {
				return part, nil
			}
		}

		if aliasType == "" && len(name.Parts) == 1 {
			part := strings.ToLower(string(name.Parts[0].(*ast.NamePart).Value))

			switch part {
			case "self":
				fallthrough
			case "static":
				fallthrough
			case "parent":
				fallthrough
			case "int":
				fallthrough
			case "float":
				fallthrough
			case "bool":
				fallthrough
			case "string":
				fallthrough
			case "void":
				fallthrough
			case "iterable":
				fallthrough
			case "object":
				return part, nil
			}
		}

		aliasName, err := namespace.ResolveAlias(nameNode, aliasType)
		if err != nil {
			// resolve as relative name if alias not found
			if namespace.Namespace == "" {
				return concatNameParts(name.Parts), nil
			}
			return namespace.Namespace + "\\" + concatNameParts(name.Parts), nil
		}

		if len(name.Parts) > 1 {
			// if name qualified, replace first part by alias
			return aliasName + "\\" + concatNameParts(name.Parts[1:]), nil
		}

		return aliasName, nil
	}

	return "", errors.New("must be instance of name.Names")
}

// ResolveAlias returns alias or error if not found
func (namespace *Namespace) ResolveAlias(nameNode ast.Vertex, aliasType string) (string, error) {
	aliasType = strings.ToLower(aliasType)
	nameParts := nameNode.(*ast.Name).Parts

	firstPartStr := string(nameParts[0].(*ast.NamePart).Value)

	if len(nameParts) > 1 { // resolve aliases for qualified names, always against class alias type
		firstPartStr = strings.ToLower(firstPartStr)
		aliasType = ""
	} else {
		if aliasType != "const" { // constants are case-sensitive
			firstPartStr = strings.ToLower(firstPartStr)
		}
	}

	aliasName, ok := namespace.Aliases[aliasType][firstPartStr]
	if !ok {
		return "", errors.New("not found")
	}

	return aliasName, nil
}

func concatNameParts(parts ...[]ast.Vertex) string {
	result := ""

	for _, part := range parts {
		for _, namePart := range part {
			if result == "" {
				result = string(namePart.(*ast.NamePart).Value)
			} else {
				result = result + "\\" + string(namePart.(*ast.NamePart).Value)
			}
		}
	}

	return result
}
