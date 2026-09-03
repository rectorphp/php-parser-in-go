package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// parseEnum parses PHP 8 source and returns the root plus every parser error it
// collected.
func parseEnum(source string) (ast.Vertex, []*errors.Error) {
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

// firstEnum returns the first statement as an enum declaration.
func firstEnum(root ast.Vertex) *ast.StmtEnum {
	return root.(*ast.Root).Stmts[0].(*ast.StmtEnum)
}

func TestPureEnum(test *testing.T) {
	root, parserErrors := parseEnum(`<?php enum Suit { case Hearts; case Spades; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	enum := firstEnum(root)
	if string(enum.Name.(*ast.Identifier).Value) != "Suit" {
		test.Errorf("expected enum name Suit, got %s", enum.Name.(*ast.Identifier).Value)
	}
	if enum.Type != nil {
		test.Errorf("expected no backing type, got %T", enum.Type)
	}
	if len(enum.Stmts) != 2 {
		test.Fatalf("expected 2 cases, got %d", len(enum.Stmts))
	}

	enumCase := enum.Stmts[0].(*ast.StmtEnumCase)
	if enumCase.Expr != nil {
		test.Errorf("expected a pure case to have no value, got %T", enumCase.Expr)
	}
}

func TestBackedEnum(test *testing.T) {
	root, parserErrors := parseEnum(`<?php enum RedirectUrlToken: string { case PageLink = 'pagelink'; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	enum := firstEnum(root)
	backingType, ok := enum.Type.(*ast.Name)
	if !ok {
		test.Fatalf("expected a name as backing type, got %T", enum.Type)
	}
	if string(backingType.Parts[0].(*ast.NamePart).Value) != "string" {
		test.Errorf("expected backing type string, got %s", backingType.Parts[0].(*ast.NamePart).Value)
	}

	enumCase := enum.Stmts[0].(*ast.StmtEnumCase)
	if enumCase.Expr == nil {
		test.Error("expected the case to carry a value")
	}
}

func TestEnumImplementsInterface(test *testing.T) {
	root, parserErrors := parseEnum(`<?php enum Suit: string implements HasColor, JsonSerializable { case Hearts = 'H'; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if len(firstEnum(root).Implements) != 2 {
		test.Errorf("expected 2 implemented interfaces, got %d", len(firstEnum(root).Implements))
	}
}

// An enum body holds ordinary members next to its cases.
func TestEnumWithConstantsAndMethods(test *testing.T) {
	root, parserErrors := parseEnum(`<?php enum Suit: string {
		case Hearts = 'H';
		const Wild = self::Hearts;
		public function color(): string { return 'red'; }
		public static function fromLabel(string $label): self {}
	}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	enum := firstEnum(root)
	if len(enum.Stmts) != 4 {
		test.Fatalf("expected 4 statements in the enum body, got %d", len(enum.Stmts))
	}
	if _, ok := enum.Stmts[1].(*ast.StmtClassConstList); !ok {
		test.Errorf("expected a class constant list, got %T", enum.Stmts[1])
	}
	if _, ok := enum.Stmts[2].(*ast.StmtClassMethod); !ok {
		test.Errorf("expected a class method, got %T", enum.Stmts[2])
	}
}

func TestAttributedEnumAndCase(test *testing.T) {
	root, parserErrors := parseEnum(`<?php #[Serializable] enum Suit { #[Deprecated] case Hearts; }`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	enum := firstEnum(root)
	if len(enum.AttrGroups) != 1 {
		test.Errorf("expected 1 attribute group on the enum, got %d", len(enum.AttrGroups))
	}
	if len(enum.Stmts[0].(*ast.StmtEnumCase).AttrGroups) != 1 {
		test.Error("expected 1 attribute group on the case")
	}
}

func TestEnumScansAsTokenOnlyBeforeAName(test *testing.T) {
	root, parserErrors := parseEnum(`<?php enum Suit {}`)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
	if firstEnum(root).EnumTkn.ID != token.T_STRING {
		test.Errorf("expected the enum keyword to keep its scanned token, got %v", firstEnum(root).EnumTkn.ID)
	}
}

// enum is contextual, not reserved: code written before PHP 8.1 keeps working.
func TestEnumStaysUsableAsIdentifier(test *testing.T) {
	sources := []string{
		`<?php function enum() {}`,
		`<?php class A { public function enum() {} }`,
		`<?php $value = $row->enum;`,
		`<?php $value = Config::enum;`,
		`<?php $value = enum($input);`,
	}

	for _, source := range sources {
		if _, parserErrors := parseEnum(source); len(parserErrors) != 0 {
			test.Errorf("unexpected parser errors for %q: %v", source, parserErrors)
		}
	}
}

func TestEnumPrintsBackUnchanged(test *testing.T) {
	source := `<?php
enum RedirectUrlToken: string implements HasLabel
{
    case PageLink     = 'pagelink';
    case FormField    = 'formfield';

    public static function names(): array
    {
        return array_column(self::cases(), 'value');
    }
}
`

	root, parserErrors := parseEnum(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	if output.String() != source {
		test.Errorf("printed output differs from source:\n%s", output.String())
	}
}
