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

// parseIntersectionTypes parses PHP 8 source and returns the root plus every
// parser error it collected.
func parseIntersectionTypes(source string) (ast.Vertex, []*errors.Error) {
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

// firstMethodOfFirstClass returns the first method of the first class.
func firstMethodOfFirstClass(root ast.Vertex) *ast.StmtClassMethod {
	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	return class.Stmts[0].(*ast.StmtClassMethod)
}

func TestIntersectionPropertyType(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { private MockObject&CorePermissions $permissions; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	property := class.Stmts[0].(*ast.StmtPropertyList)
	intersection, ok := property.Type.(*ast.Intersection)
	if !ok {
		test.Fatalf("expected an intersection type, got %T", property.Type)
	}
	if len(intersection.Types) != 2 {
		test.Errorf("expected 2 member types, got %d", len(intersection.Types))
	}
}

func TestIntersectionReturnType(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { private function get(): Email&MockObject {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if _, ok := firstMethodOfFirstClass(root).ReturnType.(*ast.Intersection); !ok {
		test.Errorf("expected an intersection return type, got %T", firstMethodOfFirstClass(root).ReturnType)
	}
}

func TestIntersectionParameterType(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { public function set(Countable&Traversable $items) {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	parameter := firstMethodOfFirstClass(root).Params[0].(*ast.Parameter)
	if _, ok := parameter.Type.(*ast.Intersection); !ok {
		test.Errorf("expected an intersection parameter type, got %T", parameter.Type)
	}
}

func TestIntersectionOfThreeTypes(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { private X&Y&Z $value; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	intersection := class.Stmts[0].(*ast.StmtPropertyList).Type.(*ast.Intersection)
	if len(intersection.Types) != 3 {
		test.Errorf("expected 3 member types, got %d", len(intersection.Types))
	}
	if len(intersection.SeparatorTkns) != 2 {
		test.Errorf("expected 2 separators, got %d", len(intersection.SeparatorTkns))
	}
}

// The `&` disambiguation must not disturb the other meanings of the character.
func TestAmpersandKeepsItsOtherMeanings(test *testing.T) {
	sources := []string{
		`<?php function sort(array &$items) {}`,
		`<?php function &getReference() {}`,
		`<?php $mask = $flags & MASK;`,
		`<?php $mask = $flags & $other;`,
		`<?php $alias = &$original;`,
		`<?php foreach ($items as &$item) {}`,
		`<?php $closure = function () use (&$counter) {};`,
		`<?php $pairs = [&$first, 'key' => &$second];`,
		`<?php function collect(...$args) {}`,
		`<?php class A { public function __construct(private readonly Logger&Stringable $logger) {} }`,
	}

	for _, source := range sources {
		if _, parserErrors := parseIntersectionTypes(source); len(parserErrors) != 0 {
			test.Errorf("unexpected parser errors for %q: %v", source, parserErrors)
		}
	}
}

func TestIntersectionTypesPrintBackUnchanged(test *testing.T) {
	source := `<?php
class A
{
    private MockObject&CorePermissions $permissions;

    private function getMockEmail(bool $published = true): Email&MockObject
    {
        return $this->flags & MASK;
    }
}
`

	root, parserErrors := parseIntersectionTypes(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
