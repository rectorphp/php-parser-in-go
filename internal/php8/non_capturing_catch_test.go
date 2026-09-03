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

// parseNonCapturingCatch parses PHP 8 source and returns the root plus every
// parser error it collected.
func parseNonCapturingCatch(source string) (ast.Vertex, []*errors.Error) {
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

// catchOf returns the first catch block of the first try statement.
func catchOf(root ast.Vertex) *ast.StmtCatch {
	try := root.(*ast.Root).Stmts[0].(*ast.StmtTry)
	return try.Catches[0].(*ast.StmtCatch)
}

func TestCatchWithoutVariable(test *testing.T) {
	root, parserErrors := parseNonCapturingCatch(`<?php try {} catch (Exception) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	catch := catchOf(root)
	if catch.Var != nil {
		test.Errorf("expected no catch variable, got %v", catch.Var)
	}
	if len(catch.Types) != 1 {
		test.Errorf("expected 1 caught type, got %d", len(catch.Types))
	}
}

func TestCatchWithoutVariableMultipleTypes(test *testing.T) {
	root, parserErrors := parseNonCapturingCatch(`<?php try {} catch (NonUniqueResultException|EntityNotFoundException) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	catch := catchOf(root)
	if len(catch.Types) != 2 {
		test.Errorf("expected 2 caught types, got %d", len(catch.Types))
	}
}

// The capturing form must keep working next to the new one.
func TestCatchWithVariableStillParses(test *testing.T) {
	root, parserErrors := parseNonCapturingCatch(`<?php try {} catch (Exception $exception) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if catchOf(root).Var == nil {
		test.Error("expected catch variable to be parsed")
	}
}

func TestNonCapturingCatchPrintsBackUnchanged(test *testing.T) {
	source := `<?php
try {
    doSomething();
} catch (NonUniqueResultException|EntityNotFoundException) {
    return null;
} catch (Throwable $throwable) {
    throw $throwable;
}
`

	root, parserErrors := parseNonCapturingCatch(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
