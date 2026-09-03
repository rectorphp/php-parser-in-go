package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestUnionReturnType(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php function f(): int|string {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	union, ok := function.ReturnType.(*ast.Union)
	if !ok {
		test.Fatalf("expected return type *ast.Union, got %T", function.ReturnType)
	}
	if len(union.Types) != 2 {
		test.Errorf("expected 2 union members, got %d", len(union.Types))
	}
	if len(union.SeparatorTkns) != 1 {
		test.Errorf("expected 1 separator token, got %d", len(union.SeparatorTkns))
	}
}

func TestUnionParameterType(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php function f(int|string|null $x) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	parameter := function.Params[0].(*ast.Parameter)
	union, ok := parameter.Type.(*ast.Union)
	if !ok {
		test.Fatalf("expected parameter type *ast.Union, got %T", parameter.Type)
	}
	if len(union.Types) != 3 {
		test.Errorf("expected 3 union members, got %d", len(union.Types))
	}
}

// A union type must print back byte-for-byte, separators included.
func TestUnionTypeRoundTrips(test *testing.T) {
	source := `<?php class A { public A\B|C|null $x; }`
	root, parserErrors := parsePhp8(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	var buffer bytes.Buffer
	root.Accept(printer.NewPrinter(&buffer))
	if buffer.String() != source {
		test.Errorf("round-trip mismatch:\n  want: %s\n  got:  %s", source, buffer.String())
	}
}
