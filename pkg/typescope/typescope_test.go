package typescope_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/typescope"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/traverser"
)

// TestResolveTypeProperty covers what ResolveType adds over ResolveValue: the
// declared type of a `$this->property` fetch, whether the property is declared
// on its own or promoted in the constructor. It is the answer behind the "is
// this an object" question a reporting rule asks.
func TestResolveTypeProperty(test *testing.T) {
	tests := []struct {
		name         string
		source       string
		expectedType string
		resolved     bool
	}{
		{
			name: "a declared class-typed property",
			source: `<?php
class OrderService
{
    private LeadRepository $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`,
			expectedType: "LeadRepository",
			resolved:     true,
		},
		{
			name: "a promoted class-typed property",
			source: `<?php
class OrderService
{
    public function __construct(private readonly LeadRepository $repository)
    {
        $local = $this->repository;
    }
}`,
			expectedType: "LeadRepository",
			resolved:     true,
		},
		{
			name: "a nullable property keeps the inner type",
			source: `<?php
class OrderService
{
    private ?LeadRepository $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`,
			expectedType: "LeadRepository",
			resolved:     true,
		},
		{
			name: "a fully qualified type keeps its trailing segment",
			source: `<?php
class OrderService
{
    private \Mautic\LeadBundle\Entity\LeadRepository $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`,
			expectedType: "LeadRepository",
			resolved:     true,
		},
		{
			name: "a scalar property resolves to its built-in type",
			source: `<?php
class Email
{
    private string $body;

    public function boot()
    {
        $local = $this->body;
    }
}`,
			expectedType: "string",
			resolved:     true,
		},
		{
			name: "an untyped property is unresolved",
			source: `<?php
class OrderService
{
    private $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`,
			resolved: false,
		},
		{
			name: "a union type is unresolved",
			source: `<?php
class OrderService
{
    private LeadRepository|OrderRepository $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`,
			resolved: false,
		},
		{
			name: "a property of another object is unresolved",
			source: `<?php
class OrderService
{
    private LeadRepository $repository;

    public function boot($other)
    {
        $local = $other->repository;
    }
}`,
			resolved: false,
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			class, assign := parseClassAndAssign(test, testCase.source)
			scopes := typescope.Collect(class)

			typeName, resolved := scopes.ResolveType(assign.Expr, assign.Position)
			if resolved != testCase.resolved {
				test.Fatalf("resolved = %v, want %v (type %q)", resolved, testCase.resolved, typeName)
			}
			if resolved && typeName != testCase.expectedType {
				test.Errorf("type = %q, want %q", typeName, testCase.expectedType)
			}
		})
	}
}

// TestResolveValueLeavesPropertiesAlone pins the split: ResolveValue is what the
// rules copying a resolved type onto a signature use, and it keeps inferring from
// exactly what it did before the property lookup existed.
func TestResolveValueLeavesPropertiesAlone(test *testing.T) {
	class, assign := parseClassAndAssign(test, `<?php
class OrderService
{
    private LeadRepository $repository;

    public function boot()
    {
        $local = $this->repository;
    }
}`)

	if typeName, resolved := typescope.Collect(class).ResolveValue(assign.Expr, assign.Position); resolved {
		test.Errorf("ResolveValue resolved a property fetch to %q, want it left alone", typeName)
	}
}

// TestResolveValueBuiltinCall covers the internal-function return types
// ResolveValue reads: a call to a builtin with a single, unconditional scalar
// return resolves to that scalar, while a function that can return false in-band
// is left unresolved so no rule declares a type the value may not hold.
func TestResolveValueBuiltinCall(test *testing.T) {
	tests := []struct {
		name         string
		call         string
		expectedType string
		resolved     bool
	}{
		{name: "sprintf resolves to string", call: "sprintf('%s', $x)", expectedType: "string", resolved: true},
		{name: "trim resolves to string", call: "trim($x)", expectedType: "string", resolved: true},
		{name: "implode resolves to string", call: "implode(',', $x)", expectedType: "string", resolved: true},
		{name: "case-insensitive name resolves", call: "STRTOLOWER($x)", expectedType: "string", resolved: true},
		{name: "strlen resolves to int", call: "strlen($x)", expectedType: "int", resolved: true},
		{name: "count resolves to int", call: "count($x)", expectedType: "int", resolved: true},
		{name: "round resolves to float", call: "round($x)", expectedType: "float", resolved: true},
		{name: "sqrt resolves to float", call: "sqrt($x)", expectedType: "float", resolved: true},
		{name: "is_array resolves to bool", call: "is_array($x)", expectedType: "bool", resolved: true},
		{name: "str_contains resolves to bool", call: "str_contains($x, 'a')", expectedType: "bool", resolved: true},
		{name: "strpos is int|false, left alone", call: "strpos($x, 'a')", resolved: false},
		{name: "json_encode is string|false, left alone", call: "json_encode($x)", resolved: false},
		{name: "substr is version-dependent, left alone", call: "substr($x, 1)", resolved: false},
		{name: "abs is int|float, left alone", call: "abs($x)", resolved: false},
		{name: "an unknown function is left alone", call: "my_helper($x)", resolved: false},
		{name: "a namespaced call is left alone", call: "\\Foo\\sprintf('%s', $x)", resolved: false},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			source := "<?php\nclass Service\n{\n    public function boot($x)\n    {\n        $local = " + testCase.call + ";\n    }\n}"
			class, assign := parseClassAndAssign(test, source)

			typeName, resolved := typescope.Collect(class).ResolveValue(assign.Expr, assign.Position)
			if resolved != testCase.resolved {
				test.Fatalf("resolved = %v, want %v (type %q)", resolved, testCase.resolved, typeName)
			}
			if resolved && typeName != testCase.expectedType {
				test.Errorf("type = %q, want %q", typeName, testCase.expectedType)
			}
		})
	}
}

// parseClassAndAssign parses a source holding one class with one assignment and
// returns both.
func parseClassAndAssign(test *testing.T, source string) (ast.Vertex, *ast.ExprAssign) {
	test.Helper()

	root, err := parser.Parse([]byte(source), conf.Config{Version: &version.Version{Major: 8, Minor: 1}})
	if err != nil {
		test.Fatalf("parse: %v", err)
	}

	collector := &classAssignCollector{Null: &visitor.Null{}}
	traverser.NewTraverser(collector).Traverse(root)
	if collector.class == nil || collector.assign == nil {
		test.Fatalf("source must hold one class with one assignment")
	}
	return collector.class, collector.assign
}

type classAssignCollector struct {
	*visitor.Null
	class  ast.Vertex
	assign *ast.ExprAssign
}

func (collector *classAssignCollector) StmtClass(node *ast.StmtClass) {
	if collector.class == nil {
		collector.class = node
	}
}

func (collector *classAssignCollector) ExprAssign(node *ast.ExprAssign) {
	if collector.assign == nil {
		collector.assign = node
	}
}

// TestIsNullableDeclaration covers the distinction a caller asking for a certain
// type depends on: a value that may be null is not certainly anything, however
// well its class name resolves.
func TestIsNullableDeclaration(test *testing.T) {
	tests := []struct {
		name     string
		source   string
		nullable bool
	}{
		{
			name: "a nullable property",
			source: `<?php
class ApiMetadataDriver
{
    private ?ClassMetadata $metadata = null;

    public function load()
    {
        $local = $this->metadata;
    }
}`,
			nullable: true,
		},
		{
			name: "a plain property",
			source: `<?php
class ApiMetadataDriver
{
    private ClassMetadata $metadata;

    public function load()
    {
        $local = $this->metadata;
    }
}`,
			nullable: false,
		},
		{
			name: "a nullable parameter",
			source: `<?php
class OrderService
{
    public function load(?LeadRepository $repository)
    {
        $local = $repository;
    }
}`,
			nullable: true,
		},
		{
			name: "an untyped parameter",
			source: `<?php
class OrderService
{
    public function load($repository)
    {
        $local = $repository;
    }
}`,
			nullable: false,
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			class, assign := parseClassAndAssign(test, testCase.source)

			nullable := typescope.Collect(class).IsNullableDeclaration(assign.Expr, assign.Position)
			if nullable != testCase.nullable {
				test.Errorf("IsNullableDeclaration = %v, want %v", nullable, testCase.nullable)
			}
		})
	}
}
