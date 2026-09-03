package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// A throw used in expression position (here as a coalesce fallback) becomes an
// ast.ExprThrow.
func TestThrowExpression(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php $value = $input ?? throw new InvalidArgumentException();`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	coalesce := assign.Expr.(*ast.ExprBinaryCoalesce)
	if _, ok := coalesce.Right.(*ast.ExprThrow); !ok {
		test.Errorf("expected coalesce right side *ast.ExprThrow, got %T", coalesce.Right)
	}
}

func TestThrowInArrowFunction(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php $f = fn() => throw new E();`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	arrow := assign.Expr.(*ast.ExprArrowFunction)
	if _, ok := arrow.Expr.(*ast.ExprThrow); !ok {
		test.Errorf("expected arrow body *ast.ExprThrow, got %T", arrow.Expr)
	}
}

// A statement-position throw stays an ast.StmtThrow — the expression form must
// not change it.
func TestThrowStatementStaysStatement(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php throw new E();`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
	if _, ok := root.(*ast.Root).Stmts[0].(*ast.StmtThrow); !ok {
		test.Errorf("expected *ast.StmtThrow, got %T", root.(*ast.Root).Stmts[0])
	}
}

func TestThrowExpressionRoundTrips(test *testing.T) {
	source := `<?php $x = $c ? throw new E() : $y;`
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
