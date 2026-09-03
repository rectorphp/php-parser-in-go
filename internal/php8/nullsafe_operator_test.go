package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestNullsafePropertyFetch(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php $user?->name;`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	expression := root.(*ast.Root).Stmts[0].(*ast.StmtExpression)
	if _, ok := expression.Expr.(*ast.ExprNullsafePropertyFetch); !ok {
		test.Errorf("expected *ast.ExprNullsafePropertyFetch, got %T", expression.Expr)
	}
}

func TestNullsafeMethodCall(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php $user?->getName($a, $b);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	expression := root.(*ast.Root).Stmts[0].(*ast.StmtExpression)
	call, ok := expression.Expr.(*ast.ExprNullsafeMethodCall)
	if !ok {
		test.Fatalf("expected *ast.ExprNullsafeMethodCall, got %T", expression.Expr)
	}
	if len(call.Args) != 2 {
		test.Errorf("expected 2 arguments, got %d", len(call.Args))
	}
}

// A nullsafe chain must print back byte-for-byte, and plain -> must stay plain.
func TestNullsafeRoundTrips(test *testing.T) {
	for _, source := range []string{
		`<?php $a?->b()?->c;`,
		`<?php $a->b()?->c();`,
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
