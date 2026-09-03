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

// parseParameterAttributes parses PHP 8 source and returns the root plus every
// parser error it collected.
func parseParameterAttributes(source string) (ast.Vertex, []*errors.Error) {
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

// firstParameterOfFirstFunction returns the first parameter of the first function.
func firstParameterOfFirstFunction(root ast.Vertex) *ast.Parameter {
	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	return function.Params[0].(*ast.Parameter)
}

func TestAttributeOnParameter(test *testing.T) {
	root, parserErrors := parseParameterAttributes(`<?php function boot(#[Autowire(env: 'json:CACHE')] array $servers) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	parameter := firstParameterOfFirstFunction(root)
	if len(parameter.AttrGroups) != 1 {
		test.Fatalf("expected 1 attribute group, got %d", len(parameter.AttrGroups))
	}
	if parameter.Type == nil {
		test.Error("expected the parameter type to survive the attribute")
	}
}

func TestAttributeOnPromotedParameter(test *testing.T) {
	root, parserErrors := parseParameterAttributes(`<?php class A { public function __construct(#[Autowire] private Logger $logger) {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	method := class.Stmts[0].(*ast.StmtClassMethod)
	parameter := method.Params[0].(*ast.Parameter)
	if len(parameter.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group, got %d", len(parameter.AttrGroups))
	}
	if len(parameter.Modifiers) != 1 {
		test.Errorf("expected the private modifier to survive the attribute, got %v", parameter.Modifiers)
	}
}

// Only the attributed parameter carries the group; its neighbours stay clean.
func TestAttributeOnSecondParameterOnly(test *testing.T) {
	root, parserErrors := parseParameterAttributes(`<?php function boot(string $name, #[Autowire] array $servers) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	if len(function.Params[0].(*ast.Parameter).AttrGroups) != 0 {
		test.Error("expected the first parameter to have no attributes")
	}
	if len(function.Params[1].(*ast.Parameter).AttrGroups) != 1 {
		test.Error("expected the second parameter to carry one attribute group")
	}
}

func TestAttributedParametersPrintBackUnchanged(test *testing.T) {
	source := `<?php
class RedisAdapter
{
    public function __construct(
        #[Autowire(env: 'json:MAUTIC_CACHE_ADAPTER_REDIS')]
        array $servers,
        #[First, Second] private string $namespace,
    ) {
    }
}
`

	root, parserErrors := parseParameterAttributes(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
