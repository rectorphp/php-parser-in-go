package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// parseNewExpression parses PHP 8.0 source and returns the root plus every
// parser error it collected.
func parseNewExpression(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 0},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

func TestNewWithExpressionClassName(test *testing.T) {
	root, parserErrors := parseNewExpression(`<?php $object = new ($prefix . 'Handler')($config);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	newExpr, ok := assign.Expr.(*ast.ExprNew)
	if !ok {
		test.Fatalf("expected a new expression, got %T", assign.Expr)
	}
	brackets, ok := newExpr.Class.(*ast.ExprBrackets)
	if !ok {
		test.Fatalf("expected the class to be a bracketed expression, got %T", newExpr.Class)
	}
	if _, ok := brackets.Expr.(*ast.ExprBinaryConcat); !ok {
		test.Errorf("expected the bracketed class expression to be a concat, got %T", brackets.Expr)
	}
	if len(newExpr.Args) != 1 {
		test.Errorf("expected one constructor argument, got %d", len(newExpr.Args))
	}
}

func TestNewWithExpressionClassNameWithoutArguments(test *testing.T) {
	_, parserErrors := parseNewExpression(`<?php $object = new (static::getModel());`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
}

func TestNewWithExpressionClassNameDereferenced(test *testing.T) {
	_, parserErrors := parseNewExpression(`<?php return (new (static::getModel()))->newQuery();`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
}

func TestNewWithExpressionClassNamePrintsBackUnchanged(test *testing.T) {
	source := `<?php
$model = new ( static::getModel() );
$handler = new ($prefix . 'Handler')($config);
`

	root, parserErrors := parseNewExpression(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
