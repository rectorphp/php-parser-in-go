package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestVoidCastStatement(test *testing.T) {
	root, parserErrors := parsePHP85(`<?php (void) $repository->save($entity);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	statement := root.(*ast.Root).Stmts[0].(*ast.StmtExpression)
	cast, ok := statement.Expr.(*ast.ExprCastVoid)
	if !ok {
		test.Fatalf("expected a void cast, got %T", statement.Expr)
	}
	if _, ok := cast.Expr.(*ast.ExprMethodCall); !ok {
		test.Errorf("expected the cast to wrap the method call, got %T", cast.Expr)
	}
}

func TestVoidCastInForExpressions(test *testing.T) {
	root, parserErrors := parsePHP85(`<?php for ($i = 0; $i < 3; (void) $i++, (void) tick()) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	loop := root.(*ast.Root).Stmts[0].(*ast.StmtFor)
	if len(loop.Loop) != 2 {
		test.Fatalf("expected two loop expressions, got %d", len(loop.Loop))
	}
	for _, expression := range loop.Loop {
		if _, ok := expression.(*ast.ExprCastVoid); !ok {
			test.Errorf("expected a void cast, got %T", expression)
		}
	}
}

// (void) discards a result, so like in PHP it is not a value-producing expression.
func TestVoidCastIsNotAnExpression(test *testing.T) {
	_, parserErrors := parsePHP85(`<?php $result = (void) compute();`)
	if len(parserErrors) == 0 {
		test.Error("expected a syntax error for a void cast used as a value")
	}
}

func TestVoidCastPrintsBackUnchanged(test *testing.T) {
	source := `<?php
(void) $repository->save($entity);
( VOID )flush();
`

	root, parserErrors := parsePHP85(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
