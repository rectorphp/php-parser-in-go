// Package typescope resolves the static type of a value expression by tracing it
// through the function scopes of a class subtree.
//
// It answers one question — "what type is this expression, at this position?" —
// and is the reusable core behind the inferred-param-type rule and any future rule
// that must resolve indirection rather than read a literal:
//
//	scopes := typescope.Collect(classNode)
//	typeName, ok := scopes.ResolveValue(argumentExpr, callPosition)
//
// A variable resolves only when it has exactly one definition in the position's
// nearest enclosing function scope — a typed parameter or a single assignment. Two
// assignments, a reassigned parameter, an untyped parameter, or a definition in a
// different scope all leave it unresolved. Arithmetic and concat expressions
// resolve through their operands: concat is always string, modulo always int, and
// the other arithmetic operators combine their operand types (float if either is
// float, otherwise int).
//
// ResolveType additionally follows a `$this->property` fetch to the property's
// declared type, whether written as a property declaration or promoted in the
// constructor. It is kept apart from ResolveValue so the rules copying a resolved
// type onto a signature keep inferring from exactly what they did before.
package typescope

import (
	"strings"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/traverser"
)

// Scopes is what was gathered from a class subtree — its function scopes and the
// declared type of each of its properties — ready to resolve expression types
// against.
type Scopes struct {
	functionScopes []*functionScope
	propertyTypes  map[string]ast.Vertex
}

// Collect gathers every function-like scope in a class subtree, attributes each
// simple assignment to its enclosing scope, so later resolution sees a variable's
// definitions in the scope they actually belong to, and records the declared type
// of every property.
func Collect(classNode ast.Vertex) Scopes {
	collector := &scopeCollector{Null: &visitor.Null{}, propertyTypes: map[string]ast.Vertex{}}
	traverser.NewTraverser(collector).Traverse(classNode)

	scopes := Scopes{functionScopes: collector.scopes, propertyTypes: collector.propertyTypes}
	for _, assignment := range collector.assignments {
		scope := enclosingScope(assignment.position, scopes)
		if scope == nil {
			continue
		}
		scope.assignedValues[assignment.name] = append(scope.assignedValues[assignment.name], assignment.value)
	}
	return scopes
}

// ResolveType resolves an expression's type like ResolveValue does, and follows a
// `$this->property` fetch to the property's declared type on top of it.
func (scopes Scopes) ResolveType(expr ast.Vertex, atPosition *position.Position) (string, bool) {
	if propertyType, ok := scopes.resolvePropertyType(expr); ok {
		return propertyType, true
	}
	return scopes.ResolveValue(expr, atPosition)
}

// IsNullableDeclaration reports whether the type an expression is declared with
// is written nullable — a `?LeadRepository` property or parameter. It answers
// the difference between "holds a repository" and "is definitely a repository",
// which is what a caller asking for a certain type has to know: a value that may
// be null is not certainly anything.
//
// An expression with no declaration to read reports false, as does a type that
// simply is not nullable.
func (scopes Scopes) IsNullableDeclaration(expr ast.Vertex, atPosition *position.Position) bool {
	declaredType, ok := scopes.declarationOf(expr, atPosition)
	if !ok {
		return false
	}
	_, nullable := declaredType.(*ast.Nullable)
	return nullable
}

// declarationOf returns the type node an expression's type is declared by: a
// property declaration for a `$this->property` fetch, a parameter declaration
// for a variable holding one.
func (scopes Scopes) declarationOf(expr ast.Vertex, atPosition *position.Position) (ast.Vertex, bool) {
	if propertyFetch, ok := expr.(*ast.ExprPropertyFetch); ok {
		property, named := identifierName(propertyFetch.Prop)
		if !named {
			return nil, false
		}
		declaredType, ok := scopes.propertyTypes[property]
		return declaredType, ok && declaredType != nil
	}

	variable, isVariable := expr.(*ast.ExprVariable)
	if !isVariable {
		return nil, false
	}
	// the parameter map is keyed by the variable name as written, `$` included
	name, named := identifierName(variable.Name)
	if !named {
		return nil, false
	}
	scope := enclosingScope(atPosition, scopes)
	if scope == nil {
		return nil, false
	}
	declaredType, ok := scope.parameterTypes[name]
	return declaredType, ok && declaredType != nil
}

// resolvePropertyType reads the declared type of a `$this->someProperty` fetch. A
// property with no type, or one fetched on anything but $this, is unresolved —
// there is nothing to read the type from.
func (scopes Scopes) resolvePropertyType(expr ast.Vertex) (string, bool) {
	propertyFetch, ok := expr.(*ast.ExprPropertyFetch)
	if !ok {
		return "", false
	}
	receiver, isVariable := propertyFetch.Var.(*ast.ExprVariable)
	if !isVariable {
		return "", false
	}
	receiverName, named := identifierName(receiver.Name)
	if !named || receiverName != "$this" {
		return "", false
	}
	property, ok := identifierName(propertyFetch.Prop)
	if !ok {
		return "", false
	}
	declaredType, ok := scopes.propertyTypes[property]
	if !ok || declaredType == nil {
		return "", false
	}
	return declaredTypeName(declaredType)
}

// ResolveValue resolves the type of a value expression at the given position: a
// literal, a resolvable variable, or a type-bearing binary operator. It reports
// false when the type cannot be pinned to a single known name.
func (scopes Scopes) ResolveValue(expr ast.Vertex, atPosition *position.Position) (string, bool) {
	return scopes.inferValueType(expr, atPosition, map[string]bool{})
}

// ResolveClass resolves an expression to a class name at the given position — a
// `new Foo()` or a variable that traces to one. It reports false when the
// expression resolves to a built-in type or cannot be resolved at all, so callers
// see only genuine class types.
func (scopes Scopes) ResolveClass(expr ast.Vertex, atPosition *position.Position) (string, bool) {
	typeName, ok := scopes.ResolveValue(expr, atPosition)
	if !ok || IsBuiltinType(typeName) {
		return "", false
	}
	return typeName, true
}

// resolveVariableType traces a variable to its single definition in the position's
// nearest enclosing function scope. A typed parameter yields its declared type; a
// single assignment yields its right-hand side's type. Anything ambiguous — no
// definition, more than one, or an untyped parameter — is unresolved. The seen set
// breaks cycles such as `$a = $a + 1`.
func (scopes Scopes) resolveVariableType(name string, atPosition *position.Position, seen map[string]bool) (string, bool) {
	if seen[name] {
		return "", false
	}

	scope := enclosingScope(atPosition, scopes)
	if scope == nil {
		return "", false
	}

	// a variable's type is a pure function of its scope, so the same variable
	// resolved again in the same scope — e.g. one passed to several methods —
	// reuses the first result instead of walking its definition anew
	if cached, ok := scope.resolvedVariables[name]; ok {
		return cached.name, cached.resolved
	}

	typeName, resolved := scopes.resolveVariableInScope(name, scope, atPosition, seen)
	scope.resolvedVariables[name] = resolvedVariable{name: typeName, resolved: resolved}
	return typeName, resolved
}

// resolveVariableInScope resolves a variable against the scope it belongs to: a
// typed parameter yields its declared type, a single assignment yields its
// right-hand side's type, and anything ambiguous is unresolved.
func (scopes Scopes) resolveVariableInScope(name string, scope *functionScope, atPosition *position.Position, seen map[string]bool) (string, bool) {
	parameterType, isParameter := scope.parameterTypes[name]
	assignedValues := scope.assignedValues[name]

	definitionCount := len(assignedValues)
	if isParameter {
		definitionCount++
	}
	if definitionCount != 1 {
		return "", false
	}

	if isParameter {
		return builtinTypeName(parameterType)
	}

	nextSeen := map[string]bool{name: true}
	for seenName := range seen {
		nextSeen[seenName] = true
	}
	return scopes.inferValueType(assignedValues[0], atPosition, nextSeen)
}

// inferValueType resolves the type of a value expression: a literal, a resolvable
// variable, or a type-bearing binary operator. Concatenation is always string,
// modulo is always int, and the other arithmetic operators combine their operand
// types. Anything else is unresolved.
func (scopes Scopes) inferValueType(expr ast.Vertex, atPosition *position.Position, seen map[string]bool) (string, bool) {
	switch node := expr.(type) {
	case *ast.ScalarString:
		return "string", true
	case *ast.ScalarLnumber:
		return "int", true
	case *ast.ScalarDnumber:
		return "float", true
	case *ast.ExprArray:
		return "array", true
	case *ast.ExprConstFetch:
		if isBooleanConstant(node) {
			return "bool", true
		}
	case *ast.ExprVariable:
		name, ok := identifierName(node.Name)
		if !ok {
			return "", false
		}
		return scopes.resolveVariableType(name, atPosition, seen)
	case *ast.ExprBinaryConcat:
		return "string", true
	case *ast.ExprBinaryMod:
		return "int", true
	case *ast.ExprBinaryPlus:
		return scopes.combineNumericTypes(node.Left, node.Right, atPosition, seen)
	case *ast.ExprBinaryMinus:
		return scopes.combineNumericTypes(node.Left, node.Right, atPosition, seen)
	case *ast.ExprBinaryMul:
		return scopes.combineNumericTypes(node.Left, node.Right, atPosition, seen)
	case *ast.ExprBinaryPow:
		return scopes.combineNumericTypes(node.Left, node.Right, atPosition, seen)
	case *ast.ExprNew:
		return classNameFromNew(node)
	case *ast.ExprFunctionCall:
		return builtinCallReturnType(node)
	}
	return "", false
}

// builtinCallReturnType resolves a call to an internal PHP function whose return
// type is a single, unconditional scalar — `string`, `int`, `float`, or `bool` —
// to that type. Only a plain single-part name is read, matched case-insensitively
// as PHP function names are; a `\sprintf` or namespaced call is left alone.
func builtinCallReturnType(node *ast.ExprFunctionCall) (string, bool) {
	name, ok := node.Function.(*ast.Name)
	if !ok || len(name.Parts) != 1 {
		return "", false
	}
	part, ok := name.Parts[0].(*ast.NamePart)
	if !ok {
		return "", false
	}
	typeName, ok := builtinReturnTypes[strings.ToLower(string(part.Value))]
	return typeName, ok
}

// builtinReturnTypes maps internal PHP functions with a single, unconditional
// scalar return type to that type. A function is listed only when it returns the
// same scalar in every case across PHP 7 and 8 — never a union, nullable, or a
// `false`/`null` on failure. That excludes the large families that signal failure
// in-band: `strpos`/`preg_match` (int|false), `json_encode`/`base64_decode`
// (string|false), `substr`/`mb_strlen` (version-dependent false), and the
// numeric-or-float pair `abs`/`pow`. Adding a false-returning function here would
// let the param-type rules declare a type the value does not always hold.
var builtinReturnTypes = map[string]string{
	// string
	"sprintf": "string", "vsprintf": "string",
	"strtolower": "string", "strtoupper": "string",
	"lcfirst": "string", "ucfirst": "string", "ucwords": "string",
	"trim": "string", "ltrim": "string", "rtrim": "string",
	"str_repeat": "string", "str_pad": "string", "strrev": "string", "nl2br": "string",
	"addslashes": "string", "stripslashes": "string",
	"addcslashes": "string", "stripcslashes": "string", "quotemeta": "string",
	"htmlspecialchars": "string", "htmlentities": "string",
	"html_entity_decode": "string", "htmlspecialchars_decode": "string", "strip_tags": "string",
	"dechex": "string", "decoct": "string", "decbin": "string",
	"chr": "string", "str_rot13": "string", "number_format": "string",
	"serialize": "string", "base64_encode": "string", "bin2hex": "string",
	"urlencode": "string", "urldecode": "string",
	"rawurlencode": "string", "rawurldecode": "string",
	"implode": "string", "join": "string",
	"basename": "string", "dirname": "string", "uniqid": "string", "preg_quote": "string",
	"gettype": "string", "date": "string", "gmdate": "string", "md5": "string", "sha1": "string",
	"mb_strtolower": "string", "mb_strtoupper": "string", "mb_convert_case": "string",

	// int
	"strlen": "int", "count": "int", "sizeof": "int",
	"strcmp": "int", "strcasecmp": "int", "strncmp": "int", "strncasecmp": "int",
	"strnatcmp": "int", "strnatcasecmp": "int",
	"ord": "int", "crc32": "int", "time": "int", "intdiv": "int",
	"levenshtein": "int", "similar_text": "int", "substr_count": "int",
	"strspn": "int", "strcspn": "int",
	"rand": "int", "mt_rand": "int", "random_int": "int",
	"getrandmax": "int", "mt_getrandmax": "int",
	"memory_get_usage": "int", "memory_get_peak_usage": "int",

	// float
	"floatval": "float", "sqrt": "float",
	"sin": "float", "cos": "float", "tan": "float",
	"asin": "float", "acos": "float", "atan": "float", "atan2": "float",
	"sinh": "float", "cosh": "float", "tanh": "float",
	"exp": "float", "log": "float", "log10": "float",
	"deg2rad": "float", "rad2deg": "float", "hypot": "float", "fmod": "float",
	"pi": "float", "lcg_value": "float",
	"ceil": "float", "floor": "float", "round": "float",

	// bool
	"ctype_alpha": "bool", "ctype_alnum": "bool", "ctype_digit": "bool", "ctype_space": "bool",
	"ctype_upper": "bool", "ctype_lower": "bool", "ctype_punct": "bool", "ctype_xdigit": "bool",
	"ctype_cntrl": "bool", "ctype_graph": "bool", "ctype_print": "bool",
	"is_array": "bool", "is_string": "bool", "is_int": "bool", "is_integer": "bool",
	"is_long": "bool", "is_float": "bool", "is_double": "bool", "is_bool": "bool",
	"is_null": "bool", "is_object": "bool", "is_callable": "bool", "is_iterable": "bool",
	"is_scalar": "bool", "is_numeric": "bool", "is_countable": "bool",
	"is_a": "bool", "is_subclass_of": "bool",
	"boolval": "bool", "in_array": "bool", "array_key_exists": "bool", "key_exists": "bool",
	"str_contains": "bool", "str_starts_with": "bool", "str_ends_with": "bool",
	"function_exists": "bool", "method_exists": "bool", "property_exists": "bool",
	"class_exists": "bool", "interface_exists": "bool", "trait_exists": "bool", "enum_exists": "bool",
	"defined": "bool", "extension_loaded": "bool", "hash_equals": "bool", "fnmatch": "bool",
}

// classNameFromNew returns the class of a `new Foo()` expression, accepting only a
// plain single-part class name so it can be copied onto another signature as
// written. A qualified, fully-qualified, or dynamic (`new $class()`) class is
// unresolved.
func classNameFromNew(node *ast.ExprNew) (string, bool) {
	name, ok := node.Class.(*ast.Name)
	if !ok || len(name.Parts) != 1 {
		return "", false
	}
	part, ok := name.Parts[0].(*ast.NamePart)
	if !ok {
		return "", false
	}
	return string(part.Value), true
}

// IsBuiltinType reports whether a resolved type name is a PHP built-in type rather
// than a class name, so the object and class rules can tell the two apart.
func IsBuiltinType(name string) bool {
	switch name {
	case "int", "float", "string", "bool", "array", "object", "iterable",
		"callable", "void", "null", "false", "true", "mixed", "never",
		"self", "static", "parent":
		return true
	}
	return false
}

// combineNumericTypes resolves the result type of an arithmetic operator. Both
// operands must resolve to a numeric type; the result is float when either operand
// is float, otherwise int.
func (scopes Scopes) combineNumericTypes(left, right ast.Vertex, atPosition *position.Position, seen map[string]bool) (string, bool) {
	leftType, leftOk := scopes.inferValueType(left, atPosition, seen)
	rightType, rightOk := scopes.inferValueType(right, atPosition, seen)
	if !leftOk || !rightOk {
		return "", false
	}
	if !isNumericType(leftType) || !isNumericType(rightType) {
		return "", false
	}
	if leftType == "float" || rightType == "float" {
		return "float", true
	}
	return "int", true
}

func isNumericType(typeName string) bool {
	return typeName == "int" || typeName == "float"
}

// builtinTypeName returns a parameter's declared type as a bare name, accepting
// only a plain identifier (`int`, `array`) or a single-part name (`Foo`). A
// nullable, union, or intersection type is unresolved, since copying it verbatim
// onto another signature is not always sound.
func builtinTypeName(typeNode ast.Vertex) (string, bool) {
	switch node := typeNode.(type) {
	case *ast.Identifier:
		return string(node.Value), true
	case *ast.Name:
		if len(node.Parts) != 1 {
			return "", false
		}
		part, ok := node.Parts[0].(*ast.NamePart)
		if !ok {
			return "", false
		}
		return string(part.Value), true
	}
	return "", false
}

// functionScope is one function-like body — a class method, closure, arrow
// function, or nested function — with the parameters it declares and the
// assignments made directly within it. Assignments made inside a nested scope are
// attributed to that nested scope, not this one.
type functionScope struct {
	position       *position.Position
	parameterTypes map[string]ast.Vertex
	assignedValues map[string][]ast.Vertex

	// resolvedVariables memoises each variable's resolved type within this scope,
	// so repeated resolution of the same variable does no repeated work.
	resolvedVariables map[string]resolvedVariable
}

// resolvedVariable is a memoised variable resolution — the type name and whether
// it resolved at all.
type resolvedVariable struct {
	name     string
	resolved bool
}

// enclosingScope returns the innermost scope whose byte range contains the given
// position — the smallest one, so a position inside a closure resolves against the
// closure rather than the method around it.
func enclosingScope(target *position.Position, scopes Scopes) *functionScope {
	if target == nil {
		return nil
	}

	var innermost *functionScope
	for _, scope := range scopes.functionScopes {
		if scope.position.StartPos > target.StartPos || scope.position.EndPos < target.EndPos {
			continue
		}
		if innermost == nil || scopeWidth(scope) < scopeWidth(innermost) {
			innermost = scope
		}
	}
	return innermost
}

func scopeWidth(scope *functionScope) int {
	return scope.position.EndPos - scope.position.StartPos
}

// scopeCollector gathers every function-like node and every simple assignment in a
// class subtree. Assignments are attributed to their enclosing scope after the
// whole subtree is walked, so nested-scope assignments land in the nested scope.
type scopeCollector struct {
	*visitor.Null
	scopes        []*functionScope
	assignments   []*rawAssignment
	propertyTypes map[string]ast.Vertex
}

// StmtPropertyList records the declared type of every property it declares, so a
// `$this->someProperty` fetch can be resolved to it.
func (collector *scopeCollector) StmtPropertyList(node *ast.StmtPropertyList) {
	for _, propertyVertex := range node.Props {
		property, ok := propertyVertex.(*ast.StmtProperty)
		if !ok {
			continue
		}
		name, ok := variableName(property.Var)
		if !ok {
			continue
		}
		collector.propertyTypes[name] = node.Type
	}
}

// Parameter records a constructor-promoted property — a parameter carrying a
// visibility modifier declares a property of the same name and type.
func (collector *scopeCollector) Parameter(node *ast.Parameter) {
	if len(node.Modifiers) == 0 {
		return
	}
	name, ok := variableName(node.Var)
	if !ok {
		return
	}
	collector.propertyTypes[name] = node.Type
}

type rawAssignment struct {
	name     string
	value    ast.Vertex
	position *position.Position
}

func (collector *scopeCollector) StmtClassMethod(node *ast.StmtClassMethod) {
	collector.addScope(node.Position, node.Params)
}

func (collector *scopeCollector) ExprClosure(node *ast.ExprClosure) {
	collector.addScope(node.Position, node.Params)
}

func (collector *scopeCollector) ExprArrowFunction(node *ast.ExprArrowFunction) {
	collector.addScope(node.Position, node.Params)
}

func (collector *scopeCollector) StmtFunction(node *ast.StmtFunction) {
	collector.addScope(node.Position, node.Params)
}

func (collector *scopeCollector) ExprAssign(node *ast.ExprAssign) {
	if node.Position == nil {
		return
	}
	variable, ok := node.Var.(*ast.ExprVariable)
	if !ok {
		return
	}
	name, ok := identifierName(variable.Name)
	if !ok {
		return
	}
	collector.assignments = append(collector.assignments, &rawAssignment{name: name, value: node.Expr, position: node.Position})
}

func (collector *scopeCollector) addScope(scopePosition *position.Position, params []ast.Vertex) {
	if scopePosition == nil {
		return
	}
	collector.scopes = append(collector.scopes, &functionScope{
		position:          scopePosition,
		parameterTypes:    parameterTypeMap(params),
		assignedValues:    map[string][]ast.Vertex{},
		resolvedVariables: map[string]resolvedVariable{},
	})
}

// parameterTypeMap maps every parameter name to its declared type node, keeping
// untyped parameters as keys with a nil type so they still count as definitions.
func parameterTypeMap(params []ast.Vertex) map[string]ast.Vertex {
	types := map[string]ast.Vertex{}
	for _, parameterVertex := range params {
		parameter, ok := parameterVertex.(*ast.Parameter)
		if !ok {
			continue
		}
		variable, ok := parameter.Var.(*ast.ExprVariable)
		if !ok {
			continue
		}
		name, ok := identifierName(variable.Name)
		if !ok {
			continue
		}
		types[name] = parameter.Type
	}
	return types
}

// declaredTypeName reads a declared type as a bare name: a plain identifier
// (`int`), a class name (its trailing segment, so an imported and a fully
// qualified name agree), or the inner type of a nullable one — `?LeadRepository`
// still says the property holds a repository. A union or intersection type is
// unresolved, there being no single name to report.
func declaredTypeName(typeNode ast.Vertex) (string, bool) {
	switch node := typeNode.(type) {
	case *ast.Identifier:
		return string(node.Value), true
	case *ast.Nullable:
		return declaredTypeName(node.Expr)
	case *ast.Name:
		return trailingNamePart(node.Parts)
	case *ast.NameFullyQualified:
		return trailingNamePart(node.Parts)
	case *ast.NameRelative:
		return trailingNamePart(node.Parts)
	}
	return "", false
}

func trailingNamePart(parts []ast.Vertex) (string, bool) {
	if len(parts) == 0 {
		return "", false
	}
	part, ok := parts[len(parts)-1].(*ast.NamePart)
	if !ok {
		return "", false
	}
	return string(part.Value), true
}

// variableName returns a variable node's name without the leading dollar, so a
// declared `$repository` property and a `$this->repository` fetch agree.
func variableName(node ast.Vertex) (string, bool) {
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

// identifierName returns an identifier node's name, reporting false for any other
// node so callers can use it as a typed guard on an ast.Vertex.
func identifierName(node ast.Vertex) (string, bool) {
	identifier, ok := node.(*ast.Identifier)
	if !ok {
		return "", false
	}
	return string(identifier.Value), true
}

// isBooleanConstant reports whether a constant fetch is `true` or `false`.
func isBooleanConstant(node *ast.ExprConstFetch) bool {
	name, ok := trailingName(node.Const)
	if !ok {
		return false
	}
	return strings.EqualFold(name, "true") || strings.EqualFold(name, "false")
}

// trailingName returns the last segment of a name node.
func trailingName(node ast.Vertex) (string, bool) {
	name, ok := node.(*ast.Name)
	if !ok {
		return "", false
	}
	return trailingNamePart(name.Parts)
}
