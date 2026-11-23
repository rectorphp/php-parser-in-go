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

// parseDynamicClassConstantFetch parses PHP 8.3 source and returns the root plus
// every parser error it collected.
func parseDynamicClassConstantFetch(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 3},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

func classConstantFetchOf(test *testing.T, root ast.Vertex) *ast.ExprClassConstFetch {
	test.Helper()
	returnStmt := root.(*ast.Root).Stmts[0].(*ast.StmtReturn)
	fetch, ok := returnStmt.Expr.(*ast.ExprClassConstFetch)
	if !ok {
		test.Fatalf("expected a class constant fetch, got %T", returnStmt.Expr)
	}
	return fetch
}

func TestDynamicClassConstantFetchOnClassName(test *testing.T) {
	root, parserErrors := parseDynamicClassConstantFetch(`<?php return Status::{$name};`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	fetch := classConstantFetchOf(test, root)
	if _, ok := fetch.Const.(*ast.ExprVariable); !ok {
		test.Errorf("expected the constant name to be a variable, got %T", fetch.Const)
	}
	if fetch.OpenCurlyBracketTkn == nil || fetch.CloseCurlyBracketTkn == nil {
		test.Error("expected the curly bracket tokens to be kept")
	}
}

func TestDynamicClassConstantFetchOnVariableClass(test *testing.T) {
	root, parserErrors := parseDynamicClassConstantFetch(`<?php return $this::{'PREFIX_' . $name};`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	fetch := classConstantFetchOf(test, root)
	if _, ok := fetch.Const.(*ast.ExprBinaryConcat); !ok {
		test.Errorf("expected the constant name to be a concat, got %T", fetch.Const)
	}
}

// A plain class constant fetch keeps an identifier name and no curly brackets.
func TestPlainClassConstantFetchHasNoCurlyBrackets(test *testing.T) {
	root, parserErrors := parseDynamicClassConstantFetch(`<?php return Status::ACTIVE;`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	fetch := classConstantFetchOf(test, root)
	if _, ok := fetch.Const.(*ast.Identifier); !ok {
		test.Errorf("expected the constant name to be an identifier, got %T", fetch.Const)
	}
	if fetch.OpenCurlyBracketTkn != nil || fetch.CloseCurlyBracketTkn != nil {
		test.Error("expected no curly bracket tokens")
	}
}

func TestDynamicClassConstantFetchPrintsBackUnchanged(test *testing.T) {
	source := `<?php
$value = Status::{ $name };
$other = $this::{'PREFIX_' . $name}->value;
`

	root, parserErrors := parseDynamicClassConstantFetch(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
