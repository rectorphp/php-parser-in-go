package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestDNFParameterType(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { public function run((Countable&Traversable)|array $items) {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	parameter := firstMethodOfFirstClass(root).Params[0].(*ast.Parameter)
	union, ok := parameter.Type.(*ast.Union)
	if !ok {
		test.Fatalf("expected a union type, got %T", parameter.Type)
	}
	if len(union.Types) != 2 {
		test.Fatalf("expected 2 union members, got %d", len(union.Types))
	}

	intersection, ok := union.Types[0].(*ast.Intersection)
	if !ok {
		test.Fatalf("expected the first union member to be an intersection, got %T", union.Types[0])
	}
	if intersection.OpenParenthesisTkn == nil || intersection.CloseParenthesisTkn == nil {
		test.Error("expected the parenthesized intersection to keep its parenthesis tokens")
	}
}

func TestDNFReturnType(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { private function get(): (X&Y)|null {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if _, ok := firstMethodOfFirstClass(root).ReturnType.(*ast.Union); !ok {
		test.Errorf("expected a union return type, got %T", firstMethodOfFirstClass(root).ReturnType)
	}
}

func TestDNFPropertyWithTwoParenthesizedIntersections(test *testing.T) {
	root, parserErrors := parseIntersectionTypes(`<?php class A { private (X&Y)|(A&B) $value; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	union := class.Stmts[0].(*ast.StmtPropertyList).Type.(*ast.Union)
	if len(union.Types) != 2 {
		test.Fatalf("expected 2 union members, got %d", len(union.Types))
	}
	for index, member := range union.Types {
		if _, ok := member.(*ast.Intersection); !ok {
			test.Errorf("expected union member %d to be an intersection, got %T", index, member)
		}
	}
}

func TestDNFTypesPrintBackUnchanged(test *testing.T) {
	source := `<?php
class A
{
    private (X&Y)|null $value;

    public function run((Countable&Traversable)|array $items): (A&B)|(C&D)
    {
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
