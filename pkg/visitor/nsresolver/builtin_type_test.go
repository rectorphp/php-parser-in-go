package nsresolver_test

import (
	"slices"
	"testing"

	"gotest.tools/assert"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/nsresolver"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/traverser"
)

// resolveTypeNames returns the sorted resolved names of every name node in src.
func resolveTypeNames(test *testing.T, src string) []string {
	phpVersion, _ := version.New("8.4")
	root, err := parser.Parse([]byte(src), conf.Config{Version: phpVersion})
	assert.NilError(test, err)

	resolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(resolver).Traverse(root)

	var resolvedNames []string
	for node, resolvedName := range resolver.ResolvedNames {
		if _, isName := node.(*ast.Name); isName {
			resolvedNames = append(resolvedNames, resolvedName)
		}
	}
	slices.Sort(resolvedNames)

	return resolvedNames
}

func TestDoNotResolveBuiltinTypeInParam(test *testing.T) {
	resolvedNames := resolveTypeNames(test, `<?php
namespace App;
function run(null $a, false $b, true $c, mixed $d) {}
`)

	assert.DeepEqual(test, []string{"false", "mixed", "null", "true"}, resolvedNames)
}

func TestDoNotResolveBuiltinTypeInReturn(test *testing.T) {
	resolvedNames := resolveTypeNames(test, `<?php
namespace App;
function first(): null {}
function second(): false {}
function third(): true {}
function fourth(): mixed {}
function fifth(): never {}
`)

	assert.DeepEqual(test, []string{"false", "mixed", "never", "null", "true"}, resolvedNames)
}

func TestDoNotResolveBuiltinTypeInProperty(test *testing.T) {
	resolvedNames := resolveTypeNames(test, `<?php
namespace App;
final class Holder
{
    public null $a;
    public false $b;
    public true $c;
    public mixed $d;
}
`)

	assert.DeepEqual(test, []string{"false", "mixed", "null", "true"}, resolvedNames)
}

func TestDoNotResolveBuiltinTypeInUnionAndNullable(test *testing.T) {
	resolvedNames := resolveTypeNames(test, `<?php
namespace App;
final class Holder
{
    public function run(Item|NULL $item, ?Mixed $value): Item|False {}
}
`)

	assert.DeepEqual(test, []string{"App\\Item", "App\\Item", "false", "mixed", "null"}, resolvedNames)
}

func TestResolveClassNamedLikeBuiltinType(test *testing.T) {
	resolvedNames := resolveTypeNames(test, `<?php
namespace App;
function run(Nullable $a, Mixedness $b): Never\Type {}
`)

	assert.DeepEqual(test, []string{"App\\Mixedness", "App\\Never\\Type", "App\\Nullable"}, resolvedNames)
}

func TestResolveMixedAndNeverAsClassOutsideTypePosition(test *testing.T) {
	phpVersion, _ := version.New("7.4")
	root, err := parser.Parse([]byte(`<?php
namespace App;
class Mixed {}
class Never {}
class Child extends Mixed {}
function run($value) {
    new Never();
    return $value instanceof Mixed;
}
`), conf.Config{Version: phpVersion})
	assert.NilError(test, err)

	resolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(resolver).Traverse(root)

	var resolvedNames []string
	for node, resolvedName := range resolver.ResolvedNames {
		if _, isName := node.(*ast.Name); isName {
			resolvedNames = append(resolvedNames, resolvedName)
		}
	}
	slices.Sort(resolvedNames)

	assert.DeepEqual(test, []string{"App\\Mixed", "App\\Mixed", "App\\Never"}, resolvedNames)
}
