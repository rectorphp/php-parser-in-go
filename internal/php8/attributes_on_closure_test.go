package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func parseClosureExpression(test *testing.T, source string) ast.Vertex {
	root, parserErrors := parseIntersectionTypes(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	return root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr
}

func TestAttributesOnClosure(test *testing.T) {
	expression := parseClosureExpression(test, `<?php #[Pure] function () {};`)

	closure, ok := expression.(*ast.ExprClosure)
	if !ok {
		test.Fatalf("expected a closure, got %T", expression)
	}
	if len(closure.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group on the closure, got %d", len(closure.AttrGroups))
	}
	if closure.StaticTkn != nil {
		test.Errorf("expected no static token")
	}
	if closure.Position.StartPos != 6 {
		test.Errorf("expected the closure to start at the attribute (6), got %d", closure.Position.StartPos)
	}
}

func TestAttributesOnStaticClosure(test *testing.T) {
	expression := parseClosureExpression(test, `<?php #[A] #[B] static function () {};`)

	closure, ok := expression.(*ast.ExprClosure)
	if !ok {
		test.Fatalf("expected a closure, got %T", expression)
	}
	if len(closure.AttrGroups) != 2 {
		test.Errorf("expected 2 attribute groups on the closure, got %d", len(closure.AttrGroups))
	}
	if closure.StaticTkn == nil {
		test.Errorf("expected a static token")
	}
	if closure.Position.StartPos != 6 {
		test.Errorf("expected the closure to start at the first attribute (6), got %d", closure.Position.StartPos)
	}
}

func TestAttributesOnArrowFunction(test *testing.T) {
	expression := parseClosureExpression(test, `<?php #[A] #[B] fn() => 1;`)

	arrowFunction, ok := expression.(*ast.ExprArrowFunction)
	if !ok {
		test.Fatalf("expected an arrow function, got %T", expression)
	}
	if len(arrowFunction.AttrGroups) != 2 {
		test.Errorf("expected 2 attribute groups on the arrow function, got %d", len(arrowFunction.AttrGroups))
	}
	if arrowFunction.Position.StartPos != 6 {
		test.Errorf("expected the arrow function to start at the first attribute (6), got %d", arrowFunction.Position.StartPos)
	}
}

func TestAttributesOnStaticArrowFunction(test *testing.T) {
	expression := parseClosureExpression(test, `<?php #[A] static fn() => 1;`)

	arrowFunction, ok := expression.(*ast.ExprArrowFunction)
	if !ok {
		test.Fatalf("expected an arrow function, got %T", expression)
	}
	if len(arrowFunction.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group on the arrow function, got %d", len(arrowFunction.AttrGroups))
	}
	if arrowFunction.StaticTkn == nil {
		test.Errorf("expected a static token")
	}
}

func TestAttributesOnClosureAsArgument(test *testing.T) {
	expression := parseClosureExpression(test, `<?php array_map(#[Pure] fn($item) => $item * 2, $items);`)

	call, ok := expression.(*ast.ExprFunctionCall)
	if !ok {
		test.Fatalf("expected a function call, got %T", expression)
	}

	argument := call.Args[0].(*ast.Argument)
	arrowFunction, ok := argument.Expr.(*ast.ExprArrowFunction)
	if !ok {
		test.Fatalf("expected an arrow function argument, got %T", argument.Expr)
	}
	if len(arrowFunction.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group on the arrow function, got %d", len(arrowFunction.AttrGroups))
	}
}

func TestAttributesOnClosurePrintBackUnchanged(test *testing.T) {
	source := `<?php
$pure = #[Pure] function () {};
$both = /* note */ #[A]
    #[B( 1 )]   static   function ( $a ) use ( $b ) : int { return $a; };
$arrow = #[A] fn() => 1;
$staticArrow = #[A] static fn() => 1;
array_map(#[Pure] fn($item) => $item * 2, $items);
#[Attribute(callback: #[Inner] static fn() => null)]
function foo() {}
`

	root, parserErrors := parseIntersectionTypes(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
