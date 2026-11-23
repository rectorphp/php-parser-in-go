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

// parseCallTimeReference parses source and returns the root plus every parser error it collected.
func parseCallTimeReference(source string) (ast.Vertex, []*errors.Error) {
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

// referenceArguments parses a single call statement and returns its arguments.
func referenceArguments(test *testing.T, source string) []ast.Vertex {
	root, parserErrors := parseCallTimeReference(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	switch call := root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(type) {
	case *ast.ExprFunctionCall:
		return call.Args
	case *ast.ExprMethodCall:
		return call.Args
	case *ast.ExprStaticCall:
		return call.Args
	default:
		test.Fatalf("unexpected call node %T", call)
		return nil
	}
}

func assertReferenceArgument(test *testing.T, argument ast.Vertex, variableName string) {
	referenceArgument := argument.(*ast.Argument)
	if referenceArgument.AmpersandTkn == nil {
		test.Fatal("expected the argument to carry the ampersand token")
	}
	variable := referenceArgument.Expr.(*ast.ExprVariable)
	if name := string(variable.Name.(*ast.Identifier).Value); name != variableName {
		test.Errorf("expected variable %s, got %s", variableName, name)
	}
}

func TestCallTimeReferenceInFunctionCall(test *testing.T) {
	arguments := referenceArguments(test, `<?php foo($a, &$b);`)
	if len(arguments) != 2 {
		test.Fatalf("expected 2 arguments, got %d", len(arguments))
	}
	if arguments[0].(*ast.Argument).AmpersandTkn != nil {
		test.Error("expected the plain argument to carry no ampersand token")
	}
	assertReferenceArgument(test, arguments[1], "$b")
}

func TestCallTimeReferenceInMethodCall(test *testing.T) {
	arguments := referenceArguments(test, `<?php $object->send(&$message);`)
	assertReferenceArgument(test, arguments[0], "$message")
}

func TestCallTimeReferenceInStaticCall(test *testing.T) {
	arguments := referenceArguments(test, `<?php Mailer::send(&$message);`)
	assertReferenceArgument(test, arguments[0], "$message")
}

func TestCallTimeReferenceKeepsBitwiseAnd(test *testing.T) {
	arguments := referenceArguments(test, `<?php foo($a & $b);`)
	argument := arguments[0].(*ast.Argument)
	if argument.AmpersandTkn != nil {
		test.Error("expected a bitwise and argument to carry no ampersand token")
	}
	if _, isBitwiseAnd := argument.Expr.(*ast.ExprBinaryBitwiseAnd); !isBitwiseAnd {
		test.Errorf("expected a bitwise and expression, got %T", argument.Expr)
	}
}

func TestCallTimeReferencePrintsBackUnchanged(test *testing.T) {
	source := `<?php
foo( & $a, $b);
$object->send(&$message);
Mailer::send(&$list[0], & $this->items);
`
	root, parserErrors := parseCallTimeReference(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	buffer := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(buffer))

	if printed := buffer.String(); printed != source {
		test.Errorf("\nexpected: %s\ngot: %s\n", source, printed)
	}
}
