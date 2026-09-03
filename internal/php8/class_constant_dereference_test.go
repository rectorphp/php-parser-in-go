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

// parseClassConstantDereference parses PHP 8 source and returns the root plus
// every parser error it collected.
func parseClassConstantDereference(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 2},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

func TestPropertyFetchOnClassConstant(test *testing.T) {
	root, parserErrors := parseClassConstantDereference(`<?php $direction = Order::Descending->value;`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	propertyFetch, ok := assign.Expr.(*ast.ExprPropertyFetch)
	if !ok {
		test.Fatalf("expected a property fetch, got %T", assign.Expr)
	}
	if _, ok := propertyFetch.Var.(*ast.ExprClassConstFetch); !ok {
		test.Errorf("expected the property to be fetched off a class constant, got %T", propertyFetch.Var)
	}
}

func TestArrayFetchOnClassConstant(test *testing.T) {
	_, parserErrors := parseClassConstantDereference(`<?php $first = Config::DEFAULTS[0];`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
}

func TestMethodCallOnClassConstant(test *testing.T) {
	_, parserErrors := parseClassConstantDereference(`<?php $label = Suit::Hearts->label();`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
}

// A plain class constant fetch must still parse as one, not as a dereference.
func TestPlainClassConstantStillParses(test *testing.T) {
	root, parserErrors := parseClassConstantDereference(`<?php $limit = self::MAX_COUNT;`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	if _, ok := assign.Expr.(*ast.ExprClassConstFetch); !ok {
		test.Errorf("expected a class constant fetch, got %T", assign.Expr)
	}
}

func TestClassConstantDereferencePrintsBackUnchanged(test *testing.T) {
	source := `<?php
$query->orderBy($this->getTableAlias() . '.dateAdded', Order::Descending->value);
`

	root, parserErrors := parseClassConstantDereference(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
