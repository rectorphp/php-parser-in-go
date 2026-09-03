package visitor_test

import (
	"strings"
	"testing"

	"gotest.tools/assert"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor"
)

func firstClass(test *testing.T, src string) *ast.StmtClass {
	phpVersion, _ := version.New("8.4")
	root, err := parser.Parse([]byte(src), conf.Config{Version: phpVersion})
	assert.NilError(test, err)

	for _, stmt := range root.(*ast.Root).Stmts {
		if class, ok := stmt.(*ast.StmtClass); ok {
			return class
		}
		if namespace, ok := stmt.(*ast.StmtNamespace); ok {
			for _, inner := range namespace.Stmts {
				if class, ok := inner.(*ast.StmtClass); ok {
					return class
				}
			}
		}
	}
	test.Fatal("no class found")
	return nil
}

func TestGetDocComment(test *testing.T) {
	class := firstClass(test, `<?php
/**
 * @api
 */
final class Foo
{
}
`)

	doc := visitor.GetDocComment(class)
	assert.Assert(test, doc != nil)
	assert.Assert(test, strings.Contains(string(doc.Value), "@api"))
	assert.Assert(test, strings.Contains(visitor.GetDocCommentText(class), "@api"))
}

func TestGetDocCommentWithAttribute(test *testing.T) {
	class := firstClass(test, `<?php
/** @deprecated */
#[SomeAttribute]
final class Bar
{
}
`)

	assert.Assert(test, strings.Contains(visitor.GetDocCommentText(class), "@deprecated"))
}

func TestGetDocCommentNone(test *testing.T) {
	class := firstClass(test, `<?php
final class Baz
{
}
`)

	assert.Assert(test, visitor.GetDocComment(class) == nil)
	assert.Equal(test, "", visitor.GetDocCommentText(class))
}
