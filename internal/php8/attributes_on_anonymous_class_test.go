package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestAttributesOnAnonymousClass(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php $command = new #[AsCommand(name: 'foo')] class extends Command {};`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	new, ok := assign.Expr.(*ast.ExprNew)
	if !ok {
		test.Fatalf("expected a new expression, got %T", assign.Expr)
	}

	class, ok := new.Class.(*ast.StmtClass)
	if !ok {
		test.Fatalf("expected an anonymous class, got %T", new.Class)
	}
	if len(class.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group on the anonymous class, got %d", len(class.AttrGroups))
	}
}

func TestAttributesOnAnonymousClassPrintBackUnchanged(test *testing.T) {
	source := `<?php
$command = new #[AsCommand(name: 'foo', description: 'define')] class extends Command {};
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
