package php8_test

import (
	"slices"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func parsePropertyHooks(test *testing.T, source string) ast.Vertex {
	test.Helper()

	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 4},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()

	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors[0])
	}

	return parser.GetRootNode()
}

func firstClassStmt(test *testing.T, root ast.Vertex) ast.Vertex {
	test.Helper()

	switch class := root.(*ast.Root).Stmts[0].(type) {
	case *ast.StmtClass:
		return class.Stmts[0]
	case *ast.StmtInterface:
		return class.Stmts[0]
	}

	test.Fatalf("expected a class-like statement, got %T", root.(*ast.Root).Stmts[0])
	return nil
}

func hookName(hook ast.Vertex) string {
	return string(hook.(*ast.PropertyHook).Name.(*ast.Identifier).Value)
}

func TestPropertyHooksOnProperty(test *testing.T) {
	root := parsePropertyHooks(test, `<?php class A {
		public string $name {
			get => $this->name;
			set(string $value) { $this->name = $value; }
		}
	}`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	if len(propertyList.Props) != 1 || propertyList.SemiColonTkn != nil {
		test.Fatalf("expected a single property without a semicolon, got %d props", len(propertyList.Props))
	}
	if propertyList.OpenCurlyBracketTkn == nil || propertyList.CloseCurlyBracketTkn == nil {
		test.Fatal("expected the hook list curly brackets to be kept")
	}
	if len(propertyList.Hooks) != 2 {
		test.Fatalf("expected 2 hooks, got %d", len(propertyList.Hooks))
	}

	getHook := propertyList.Hooks[0].(*ast.PropertyHook)
	if hookName(getHook) != "get" || getHook.DoubleArrowTkn == nil || getHook.Expr == nil || getHook.SemiColonTkn == nil {
		test.Errorf("expected an arrow get hook, got %+v", getHook)
	}

	setHook := propertyList.Hooks[1].(*ast.PropertyHook)
	if hookName(setHook) != "set" || len(setHook.Params) != 1 {
		test.Fatalf("expected a set hook with one parameter, got %+v", setHook)
	}
	if _, ok := setHook.Stmt.(*ast.StmtStmtList); !ok {
		test.Errorf("expected a statement list body, got %T", setHook.Stmt)
	}
}

func TestPropertyHookForms(test *testing.T) {
	root := parsePropertyHooks(test, `<?php abstract class A {
		public int $count = 0 {
			#[Pure] final get { return 1; }
			set => $value * 2;
		}
	}`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	if propertyList.Props[0].(*ast.StmtProperty).Expr == nil {
		test.Error("expected the default value to be kept")
	}

	getHook := propertyList.Hooks[0].(*ast.PropertyHook)
	if len(getHook.AttrGroups) != 1 {
		test.Errorf("expected one attribute group, got %d", len(getHook.AttrGroups))
	}
	if !slices.Equal(modifierValues(getHook.Modifiers), []string{"final"}) {
		test.Errorf("expected final modifier, got %v", modifierValues(getHook.Modifiers))
	}

	setHook := propertyList.Hooks[1].(*ast.PropertyHook)
	if setHook.OpenParenthesisTkn != nil || setHook.Params != nil || setHook.Expr == nil {
		test.Errorf("expected a short set hook without parameters, got %+v", setHook)
	}
}

func TestPropertyHookByReferenceAndAbstract(test *testing.T) {
	root := parsePropertyHooks(test, `<?php abstract class A {
		abstract public array $items { &get; set; }
	}`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	getHook := propertyList.Hooks[0].(*ast.PropertyHook)
	if getHook.AmpersandTkn == nil || getHook.SemiColonTkn == nil || getHook.Stmt != nil || getHook.Expr != nil {
		test.Errorf("expected an abstract by-reference get hook, got %+v", getHook)
	}
}

func TestPropertyHookEmptyList(test *testing.T) {
	root := parsePropertyHooks(test, `<?php class A { public $a {} }`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	if propertyList.Hooks == nil || len(propertyList.Hooks) != 0 {
		test.Errorf("expected an empty, non-nil hook list, got %v", propertyList.Hooks)
	}
}

func TestPropertyHookInInterface(test *testing.T) {
	root := parsePropertyHooks(test, `<?php interface HasName { public string $name { get; } }`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	if len(propertyList.Hooks) != 1 || hookName(propertyList.Hooks[0]) != "get" {
		test.Errorf("expected a get hook, got %v", propertyList.Hooks)
	}
}

func TestPropertyHookOnPromotedParameter(test *testing.T) {
	root := parsePropertyHooks(test, `<?php class A {
		public function __construct(public int $i { get => 1; }, public int $j = 2 { set => $value; }) {}
	}`)

	method := firstClassStmt(test, root).(*ast.StmtClassMethod)
	for i, param := range method.Params {
		parameter := param.(*ast.Parameter)
		if len(parameter.Hooks) != 1 || parameter.OpenCurlyBracketTkn == nil || parameter.CloseCurlyBracketTkn == nil {
			test.Errorf("param %d: expected one hook, got %v", i, parameter.Hooks)
		}
		if parameter.Position.EndPos != parameter.CloseCurlyBracketTkn.Position.EndPos {
			test.Errorf("param %d: expected the position to end at the hook list", i)
		}
	}
}

func TestPlainPropertyHasNoHooks(test *testing.T) {
	root := parsePropertyHooks(test, `<?php class A { public $a, $b; }`)

	propertyList := firstClassStmt(test, root).(*ast.StmtPropertyList)
	if propertyList.Hooks != nil || propertyList.SemiColonTkn == nil {
		test.Errorf("expected a plain property list, got %+v", propertyList)
	}
}

func TestAsymmetricVisibility(test *testing.T) {
	root := parsePropertyHooks(test, `<?php class A {
		public private(set) string $a;
		protected(set) int $b;
		public function __construct(public PUBLIC(SET) int $c) {}
	}`)

	stmts := root.(*ast.Root).Stmts[0].(*ast.StmtClass).Stmts
	if got := modifierValues(stmts[0].(*ast.StmtPropertyList).Modifiers); !slices.Equal(got, []string{"public", "private(set)"}) {
		test.Errorf("unexpected modifiers %v", got)
	}
	if got := modifierValues(stmts[1].(*ast.StmtPropertyList).Modifiers); !slices.Equal(got, []string{"protected(set)"}) {
		test.Errorf("unexpected modifiers %v", got)
	}

	parameter := stmts[2].(*ast.StmtClassMethod).Params[0].(*ast.Parameter)
	if got := modifierValues(parameter.Modifiers); !slices.Equal(got, []string{"public", "PUBLIC(SET)"}) {
		test.Errorf("unexpected modifiers %v", got)
	}
	if parameter.Modifiers[1].(*ast.Identifier).IdentifierTkn.ID != token.T_PUBLIC_SET {
		test.Errorf("expected T_PUBLIC_SET token")
	}
}

// The curly brace offset syntax was a compile error throughout PHP 8 and PHP
// 8.4 dropped it from the grammar to make room for property hooks.
func TestCurlyBraceOffsetIsSyntaxError(test *testing.T) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 4},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	parser := php8.NewParser(php8.NewLexer([]byte(`<?php $a{0};`), config), config)
	parser.Parse()

	if len(parserErrors) == 0 {
		test.Error("expected a syntax error for $a{0}")
	}
}
