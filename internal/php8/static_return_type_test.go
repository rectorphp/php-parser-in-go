package php8_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func parsePhp8(source string) (ast.Vertex, []*errors.Error) {
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

func TestStaticReturnType(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php function make(): static {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	returnType, ok := function.ReturnType.(*ast.Identifier)
	if !ok {
		test.Fatalf("expected return type *ast.Identifier, got %T", function.ReturnType)
	}
	if string(returnType.Value) != "static" {
		test.Errorf("expected return type value \"static\", got %q", returnType.Value)
	}
}

func TestNullableStaticReturnType(test *testing.T) {
	_, parserErrors := parsePhp8(`<?php class A { public function self(): ?static {} }`)
	if len(parserErrors) != 0 {
		test.Errorf("nullable static return type should parse, got errors: %v", parserErrors)
	}
}

// static keeps working in its other roles once it is also a valid type name.
func TestStaticStillWorksInOtherRoles(test *testing.T) {
	sources := []string{
		`<?php $x = new static();`,
		`<?php echo static::VALUE;`,
		`<?php class A { public static int $count; }`,
		`<?php $fn = static function () {};`,
	}
	for _, source := range sources {
		if _, parserErrors := parsePhp8(source); len(parserErrors) != 0 {
			test.Errorf("%s should parse, got errors: %v", source, parserErrors)
		}
	}
}
