package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// PHP 8.0 allows a trailing comma in a closure use list.
func TestTrailingCommaInClosureUseList(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php $f = function () use ($a, &$b,) {};`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	closure := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign).Expr.(*ast.ExprClosure)
	if len(closure.Uses) != 2 {
		test.Errorf("expected 2 closure uses, got %d", len(closure.Uses))
	}
	if len(closure.UseSeparatorTkns) != 2 {
		test.Errorf("expected 2 use separator tokens including the trailing comma, got %d", len(closure.UseSeparatorTkns))
	}
}

func TestTrailingCommaInClosureUseListPrintsBack(test *testing.T) {
	source := `<?php $f = function ( ) use ( $a , & $b , ) { } ;`
	root, parserErrors := parsePhp8(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := &bytes.Buffer{}
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("\nexpected: %s\ngot: %s", source, output.String())
	}
}

func TestClosureUseListWithoutTrailingCommaStillParses(test *testing.T) {
	for _, source := range []string{
		`<?php $f = function () use ($a) {};`,
		`<?php $f = function () use ($a, $b) {};`,
		`<?php $f = function () {};`,
	} {
		if _, parserErrors := parsePhp8(source); len(parserErrors) != 0 {
			test.Errorf("%s should parse, got errors: %v", source, parserErrors)
		}
	}
}
