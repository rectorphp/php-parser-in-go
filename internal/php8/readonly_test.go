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

// parseReadonly parses PHP 8 source and returns the root plus every parser error
// it collected, so a test can assert both that a construct parses cleanly and
// what it parsed to.
func parseReadonly(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 2},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

// modifierValues returns the string value of each modifier identifier.
func modifierValues(modifiers []ast.Vertex) []string {
	values := make([]string, 0, len(modifiers))
	for _, modifier := range modifiers {
		if identifier, ok := modifier.(*ast.Identifier); ok {
			values = append(values, string(identifier.Value))
		}
	}
	return values
}

func TestReadonlyScansAsToken(test *testing.T) {
	lexer := php8.NewLexer([]byte(`<?php readonly `), conf.Config{})

	var ids []token.ID
	for {
		tkn := lexer.Lex()
		if tkn == nil || tkn.ID == 0 {
			break
		}
		ids = append(ids, tkn.ID)
	}

	if len(ids) == 0 || ids[0] != token.T_READONLY {
		test.Fatalf("expected first token T_READONLY, got %v", ids)
	}
}

func TestReadonlyTypedProperty(test *testing.T) {
	root, parserErrors := parseReadonly(`<?php class A { public readonly int $x; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	property := class.Stmts[0].(*ast.StmtPropertyList)
	values := modifierValues(property.Modifiers)
	if !slices.Contains(values, "public") || !slices.Contains(values, "readonly") {
		test.Errorf("expected modifiers [public readonly], got %v", values)
	}
}

func TestReadonlyPromotedParameter(test *testing.T) {
	root, parserErrors := parseReadonly(`<?php class A { public function __construct(public readonly int $x) {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	method := class.Stmts[0].(*ast.StmtClassMethod)
	parameter := method.Params[0].(*ast.Parameter)
	values := modifierValues(parameter.Modifiers)
	if !slices.Contains(values, "readonly") {
		test.Errorf("expected promoted parameter to carry readonly modifier, got %v", values)
	}
}

// readonly is contextual, not reserved: it must still parse as an ordinary
// method name so existing code that predates PHP 8.1 keeps working.
func TestReadonlyStaysUsableAsMethodName(test *testing.T) {
	_, parserErrors := parseReadonly(`<?php class A { public function readonly() {} }`)
	if len(parserErrors) != 0 {
		test.Errorf("readonly should be usable as a method name, got errors: %v", parserErrors)
	}
}

// PHP 8.2 allows readonly as a class modifier, alone or next to final.
func TestReadonlyClass(test *testing.T) {
	root, parserErrors := parseReadonly(`<?php readonly class A {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	values := modifierValues(class.Modifiers)
	if !slices.Contains(values, "readonly") {
		test.Errorf("expected class modifiers to contain readonly, got %v", values)
	}
}

func TestFinalReadonlyClass(test *testing.T) {
	root, parserErrors := parseReadonly(`<?php final readonly class A {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	values := modifierValues(class.Modifiers)
	if len(values) != 2 || values[0] != "final" || values[1] != "readonly" {
		test.Errorf("expected modifiers [final readonly], got %v", values)
	}
}

func TestReadonlyAbstractClass(test *testing.T) {
	root, parserErrors := parseReadonly(`<?php abstract readonly class A {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	values := modifierValues(class.Modifiers)
	if len(values) != 2 || values[0] != "abstract" || values[1] != "readonly" {
		test.Errorf("expected modifiers [abstract readonly], got %v", values)
	}
}
