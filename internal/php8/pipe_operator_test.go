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

// parsePHP85 parses PHP 8.5 source and returns the root plus every parser error
// it collected.
func parsePHP85(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 5},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

func TestPipeOperator(test *testing.T) {
	root, parserErrors := parsePHP85(`<?php $slug = $title |> trim(...) |> strtolower(...);`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	outer, ok := assign.Expr.(*ast.ExprBinaryPipe)
	if !ok {
		test.Fatalf("expected a pipe, got %T", assign.Expr)
	}
	// the pipe is left-associative: ($title |> trim(...)) |> strtolower(...)
	if _, ok := outer.Left.(*ast.ExprBinaryPipe); !ok {
		test.Errorf("expected the left side to be the inner pipe, got %T", outer.Left)
	}
	if _, ok := outer.Right.(*ast.ExprFunctionCall); !ok {
		test.Errorf("expected the right side to be a first-class callable, got %T", outer.Right)
	}
}

// The pipe binds tighter than comparison and looser than concatenation.
func TestPipeOperatorPrecedence(test *testing.T) {
	root, parserErrors := parsePHP85(`<?php $matches = $a . $b |> strtoupper(...) === 'AB';`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	assign := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign)
	identical, ok := assign.Expr.(*ast.ExprBinaryIdentical)
	if !ok {
		test.Fatalf("expected the comparison at the top, got %T", assign.Expr)
	}
	pipe, ok := identical.Left.(*ast.ExprBinaryPipe)
	if !ok {
		test.Fatalf("expected the pipe under the comparison, got %T", identical.Left)
	}
	if _, ok := pipe.Left.(*ast.ExprBinaryConcat); !ok {
		test.Errorf("expected the concatenation under the pipe, got %T", pipe.Left)
	}
}

func TestPipeOperatorPrintsBackUnchanged(test *testing.T) {
	source := `<?php
$slug = $title
    |> trim(...)
    |> strtolower(...);
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
