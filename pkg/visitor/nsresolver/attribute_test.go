package nsresolver_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/nsresolver"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/traverser"
)

type attributeCollector struct {
	visitor.Null
	nodes []*ast.Attribute
}

func (collector *attributeCollector) Attribute(node *ast.Attribute) {
	collector.nodes = append(collector.nodes, node)
}

func TestResolveAttributeName(test *testing.T) {
	src := []byte(`<?php
namespace App;
use App\Attribute\AsThing;
#[AsThing]
final class Controller
{
}
`)

	phpVersion, _ := version.New("8.4")
	root, err := parser.Parse(src, conf.Config{Version: phpVersion})
	assert.NilError(test, err)

	resolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(resolver).Traverse(root)

	collector := &attributeCollector{}
	traverser.NewTraverser(collector).Traverse(root)
	assert.Equal(test, 1, len(collector.nodes))

	resolved := resolver.ResolvedNames[collector.nodes[0].Name]
	assert.Equal(test, "App\\Attribute\\AsThing", resolved)
}
