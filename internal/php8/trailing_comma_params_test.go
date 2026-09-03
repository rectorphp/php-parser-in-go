package php8_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

// PHP 8.0 allows a trailing comma in a parameter list.
func TestTrailingCommaInParameterList(test *testing.T) {
	root, parserErrors := parsePhp8(`<?php function f(int $a, string $b,) {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	function := root.(*ast.Root).Stmts[0].(*ast.StmtFunction)
	if len(function.Params) != 2 {
		test.Errorf("expected 2 parameters, got %d", len(function.Params))
	}
}

func TestTrailingCommaInPromotedConstructor(test *testing.T) {
	if _, parserErrors := parsePhp8(`<?php class A { public function __construct(public int $x,) {} }`); len(parserErrors) != 0 {
		test.Errorf("trailing comma in promoted constructor should parse, got errors: %v", parserErrors)
	}
}

func TestParameterListWithoutTrailingCommaStillParses(test *testing.T) {
	for _, source := range []string{`<?php function f(int $a) {}`, `<?php function f() {}`} {
		if _, parserErrors := parsePhp8(source); len(parserErrors) != 0 {
			test.Errorf("%s should parse, got errors: %v", source, parserErrors)
		}
	}
}
