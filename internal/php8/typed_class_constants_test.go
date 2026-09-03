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

// parseTypedConstant parses PHP 8.3 source and returns the root plus every
// parser error it collected.
func parseTypedConstant(source string) (ast.Vertex, []*errors.Error) {
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

// firstClassConstList returns the constant list of the first statement of the
// first class-like statement in the source.
func firstClassConstList(test *testing.T, source string) *ast.StmtClassConstList {
	test.Helper()

	root, parserErrors := parseTypedConstant(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	switch classLike := root.(*ast.Root).Stmts[0].(type) {
	case *ast.StmtClass:
		return classLike.Stmts[0].(*ast.StmtClassConstList)
	case *ast.StmtInterface:
		return classLike.Stmts[0].(*ast.StmtClassConstList)
	case *ast.StmtTrait:
		return classLike.Stmts[0].(*ast.StmtClassConstList)
	default:
		test.Fatalf("unexpected first statement %T", classLike)
		return nil
	}
}

// constantName returns the name of the constant at the given index.
func constantName(constList *ast.StmtClassConstList, index int) string {
	constant := constList.Consts[index].(*ast.StmtConstant)
	return string(constant.Name.(*ast.Identifier).Value)
}

func TestTypedClassConstant(test *testing.T) {
	constList := firstClassConstList(test, `<?php class A { public const string NAME = 'a'; }`)

	name, ok := constList.Type.(*ast.Name)
	if !ok {
		test.Fatalf("expected constant type *ast.Name, got %T", constList.Type)
	}
	if value := string(name.Parts[0].(*ast.NamePart).Value); value != "string" {
		test.Errorf("expected type string, got %s", value)
	}
	if got := constantName(constList, 0); got != "NAME" {
		test.Errorf("expected constant NAME, got %s", got)
	}
}

// The type is optional; a constant without one keeps a nil Type.
func TestUntypedClassConstantHasNoType(test *testing.T) {
	constList := firstClassConstList(test, `<?php class A { public const NAME = 'a'; }`)

	if constList.Type != nil {
		test.Errorf("expected no type, got %T", constList.Type)
	}
}

func TestTypedClassConstantNullableAndUnion(test *testing.T) {
	nullable := firstClassConstList(test, `<?php class A { const ?Bar C = null; }`)
	if _, ok := nullable.Type.(*ast.Nullable); !ok {
		test.Errorf("expected *ast.Nullable, got %T", nullable.Type)
	}

	union := firstClassConstList(test, `<?php class A { const int|float B = 1; }`)
	if _, ok := union.Type.(*ast.Union); !ok {
		test.Errorf("expected *ast.Union, got %T", union.Type)
	}
}

// One type covers every constant of the group.
func TestTypedClassConstantList(test *testing.T) {
	constList := firstClassConstList(test, `<?php class A { const int F = 4, G = 5; }`)

	if constList.Type == nil {
		test.Fatal("expected the group to carry a type")
	}
	if len(constList.Consts) != 2 {
		test.Fatalf("expected 2 constants, got %d", len(constList.Consts))
	}
	if got := constantName(constList, 1); got != "G" {
		test.Errorf("expected second constant G, got %s", got)
	}
}

func TestTypedClassConstantInInterfaceTraitAndEnum(test *testing.T) {
	for _, source := range []string{
		`<?php interface I { const string NAME = 'a'; }`,
		`<?php trait T { const string NAME = 'a'; }`,
		`<?php enum E: string { const string NAME = 'a'; case One = 'one'; }`,
	} {
		if _, parserErrors := parseTypedConstant(source); len(parserErrors) != 0 {
			test.Errorf("unexpected parser errors for %s: %v", source, parserErrors)
		}
	}
}

// The printer round-trips the source byte for byte, so a file that only holds
// typed constants is never rewritten.
func TestTypedClassConstantRoundTrip(test *testing.T) {
	source := "<?php\n\nfinal class A\n{\n    public const string NAME = 'a';\n\n    protected const ?Bar C = null;\n\n    const int F = 4, G = 5;\n}\n"

	root, parserErrors := parseTypedConstant(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	var output bytes.Buffer
	root.Accept(printer.NewPrinter(&output))

	if output.String() != source {
		test.Errorf("expected round trip, got:\n%s", output.String())
	}
}
