package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// Partial function application turns `?` into a placeholder for one argument,
// `name: ?` into a named one and a bare `...` into the rest: `trim(?)`.
func TestPartialFunctionApplicationPlaceholders(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php f($x, ?, y: ?, ...);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	call := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprFunctionCall)
	if len(call.Args) != 4 {
		test.Fatalf("expected 4 arguments, got %d", len(call.Args))
	}

	if call.Args[0].(*ast.Argument).Expr == nil {
		test.Error("expected the first argument to keep its expression")
	}

	positional := call.Args[1].(*ast.Argument)
	if positional.PlaceholderTkn == nil || positional.Expr != nil || positional.Name != nil {
		test.Errorf("expected a positional placeholder, got %+v", positional)
	}

	named := call.Args[2].(*ast.Argument)
	if named.PlaceholderTkn == nil || named.Expr != nil {
		test.Errorf("expected a named placeholder, got %+v", named)
	}
	if name, ok := named.Name.(*ast.Identifier); !ok || string(name.Value) != "y" {
		test.Errorf("expected the placeholder to be named y, got %v", named.Name)
	}

	rest := call.Args[3].(*ast.Argument)
	if rest.VariadicTkn == nil || rest.Expr != nil {
		test.Errorf("expected a variadic placeholder, got %+v", rest)
	}
}

func TestPartialFunctionApplicationRoundTrips(test *testing.T) {
	for _, source := range []string{
		`<?php array_map(trim(?), [' x ']);`,
		`<?php f($x, ?, y: ?, ...);`,
		`<?php $object->format(?, 'Y-m-d');`,
		`<?php Strings::pad(?, 10, ...);`,
		`<?php trim(...);`,
		`<?php f($isSet ? $a : $b);`,
	} {
		root, parserErrors := parsePhp8(source)
		if len(parserErrors) != 0 {
			test.Fatalf("%s: unexpected parser errors: %v", source, parserErrors)
		}
		var buffer bytes.Buffer
		root.Accept(printer.NewPrinter(&buffer))
		if buffer.String() != source {
			test.Errorf("round-trip mismatch:\n  want: %s\n  got:  %s", source, buffer.String())
		}
	}
}
