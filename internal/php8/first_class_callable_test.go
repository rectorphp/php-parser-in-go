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

// parseFirstClassCallable parses PHP 8 source and returns the root plus every
// parser error it collected.
func parseFirstClassCallable(source string) (ast.Vertex, []*errors.Error) {
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

// firstExpression returns the expression of the first statement.
func firstExpression(root ast.Vertex) ast.Vertex {
	return root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr
}

func TestFirstClassCallableFromFunction(test *testing.T) {
	root, parserErrors := parseFirstClassCallable(`<?php trim(...);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	call := firstExpression(root).(*ast.ExprFunctionCall)
	if len(call.Args) != 1 {
		test.Fatalf("expected 1 argument, got %d", len(call.Args))
	}

	argument := call.Args[0].(*ast.Argument)
	if argument.VariadicTkn == nil {
		test.Error("expected the placeholder argument to carry the ellipsis token")
	}
	if argument.Expr != nil {
		test.Errorf("expected the placeholder argument to have no expression, got %v", argument.Expr)
	}
}

func TestFirstClassCallableFromMethod(test *testing.T) {
	root, parserErrors := parseFirstClassCallable(`<?php $this->send(...);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	call := firstExpression(root).(*ast.ExprMethodCall)
	if len(call.Args) != 1 {
		test.Fatalf("expected 1 argument, got %d", len(call.Args))
	}
	if call.Args[0].(*ast.Argument).VariadicTkn == nil {
		test.Error("expected the placeholder argument to carry the ellipsis token")
	}
}

func TestFirstClassCallableFromStaticMethod(test *testing.T) {
	root, parserErrors := parseFirstClassCallable(`<?php Strings::trim(...);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	call := firstExpression(root).(*ast.ExprStaticCall)
	if len(call.Args) != 1 {
		test.Fatalf("expected 1 argument, got %d", len(call.Args))
	}
}

// A real variadic spread keeps its expression, so the two forms stay distinct.
func TestVariadicSpreadStillCarriesExpression(test *testing.T) {
	root, parserErrors := parseFirstClassCallable(`<?php trim(...$args);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	argument := firstExpression(root).(*ast.ExprFunctionCall).Args[0].(*ast.Argument)
	if argument.Expr == nil {
		test.Error("expected the spread argument to keep its expression")
	}
}

func TestFirstClassCallablePrintsBackUnchanged(test *testing.T) {
	source := `<?php
$emails = array_map(trim(...), explode(',', $addresses));
`

	root, parserErrors := parseFirstClassCallable(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
