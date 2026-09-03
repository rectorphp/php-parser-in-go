package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// Since PHP 8.4 a `new` with an argument list can be dereferenced without
// wrapping parentheses: `new Foo()->bar()` instead of `(new Foo())->bar()`.
func TestNewMethodCallWithoutParentheses(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php new Filesystem()->remove($paths);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	expression := root.(*ast.Root).Stmts[0].(*ast.StmtExpression)
	call, ok := expression.Expr.(*ast.ExprMethodCall)
	if !ok {
		test.Fatalf("expected *ast.ExprMethodCall, got %T", expression.Expr)
	}
	if _, ok := call.Var.(*ast.ExprNew); !ok {
		test.Errorf("expected method call receiver *ast.ExprNew, got %T", call.Var)
	}
}

func TestNewPropertyFetchWithoutParentheses(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php new Foo()->name;`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	expression := root.(*ast.Root).Stmts[0].(*ast.StmtExpression)
	fetch, ok := expression.Expr.(*ast.ExprPropertyFetch)
	if !ok {
		test.Fatalf("expected *ast.ExprPropertyFetch, got %T", expression.Expr)
	}
	if _, ok := fetch.Var.(*ast.ExprNew); !ok {
		test.Errorf("expected property fetch receiver *ast.ExprNew, got %T", fetch.Var)
	}
}

// The dereferenced `new` must print back byte-for-byte, and the plain forms that
// were valid before 8.4 (`new Foo`, `new Foo()`, `(new Foo())->bar()`) must be
// left exactly as written.
func TestNewDereferenceRoundTrips(test *testing.T) {
	for _, source := range []string{
		`<?php new Filesystem()->remove($paths);`,
		`<?php new \DateTime()->setTime(0, 0)->modify('+1 day');`,
		`<?php $mime = (string) new \finfo(FILEINFO_MIME_TYPE)->buffer($chunk);`,
		`<?php new Foo();`,
		`<?php new Foo;`,
		`<?php (new Foo())->bar();`,
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
