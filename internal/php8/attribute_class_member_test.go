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

// parseClassMemberAttributes parses PHP 8 source and returns the root plus every
// parser error it collected.
func parseClassMemberAttributes(source string) (ast.Vertex, []*errors.Error) {
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

// attributeNames returns the name of the first attribute in each group.
func attributeNames(attrGroups []ast.Vertex) []string {
	names := make([]string, 0, len(attrGroups))
	for _, attrGroup := range attrGroups {
		for _, attr := range attrGroup.(*ast.AttributeGroup).Attrs {
			name := attr.(*ast.Attribute).Name.(*ast.Name)
			for _, part := range name.Parts {
				names = append(names, string(part.(*ast.NamePart).Value))
			}
		}
	}
	return names
}

func TestAttributeOnClassMethod(test *testing.T) {
	root, parserErrors := parseClassMemberAttributes(`<?php class A { #[Route] public function index() {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	method := class.Stmts[0].(*ast.StmtClassMethod)
	names := attributeNames(method.AttrGroups)
	if len(names) != 1 || names[0] != "Route" {
		test.Errorf("expected method attribute [Route], got %v", names)
	}
}

func TestAttributeOnProperty(test *testing.T) {
	root, parserErrors := parseClassMemberAttributes(`<?php class A { #[Column(type: 'string')] private string $name; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	property := class.Stmts[0].(*ast.StmtPropertyList)
	names := attributeNames(property.AttrGroups)
	if len(names) != 1 || names[0] != "Column" {
		test.Errorf("expected property attribute [Column], got %v", names)
	}
}

func TestAttributeOnClassConst(test *testing.T) {
	root, parserErrors := parseClassMemberAttributes(`<?php class A { #[Deprecated] public const B = 1; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	constList := class.Stmts[0].(*ast.StmtClassConstList)
	names := attributeNames(constList.AttrGroups)
	if len(names) != 1 || names[0] != "Deprecated" {
		test.Errorf("expected constant attribute [Deprecated], got %v", names)
	}
}

// Several groups may stack on one member, and each group may hold several
// attributes, just like on a class declaration.
func TestMultipleAttributeGroupsOnMethod(test *testing.T) {
	root, parserErrors := parseClassMemberAttributes(`<?php class A { #[First, Second] #[Third] public function run() {} }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	class := root.(*ast.Root).Stmts[0].(*ast.StmtClass)
	method := class.Stmts[0].(*ast.StmtClassMethod)
	if len(method.AttrGroups) != 2 {
		test.Fatalf("expected 2 attribute groups, got %d", len(method.AttrGroups))
	}

	names := attributeNames(method.AttrGroups)
	if len(names) != 3 || names[0] != "First" || names[1] != "Second" || names[2] != "Third" {
		test.Errorf("expected [First Second Third], got %v", names)
	}
}

// The printer must reproduce the source byte-for-byte, attributes included.
func TestAttributedClassMembersPrintBackUnchanged(test *testing.T) {
	source := `<?php
class A
{
    #[Deprecated]
    public const B = 1;

    #[Column(type: 'string'), Index]
    private string $name;

    #[\Attributes\RunInSeparateProcess]
    public function test(): void
    {
    }
}
`

	root, parserErrors := parseClassMemberAttributes(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
