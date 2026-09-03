package ast_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

func TestToString(test *testing.T) {
	name := &ast.Name{Parts: []ast.Vertex{
		&ast.NamePart{Value: []byte("App")},
		&ast.NamePart{Value: []byte("Entity")},
		&ast.NamePart{Value: []byte("User")},
	}}
	assert.Equal(test, "App\\Entity\\User", ast.ToString(name))

	fullyQualified := &ast.NameFullyQualified{Parts: []ast.Vertex{
		&ast.NamePart{Value: []byte("App")},
		&ast.NamePart{Value: []byte("User")},
	}}
	assert.Equal(test, "App\\User", ast.ToString(fullyQualified))

	assert.Equal(test, "User", ast.ToString(&ast.NamePart{Value: []byte("User")}))
	assert.Equal(test, "string", ast.ToString(&ast.Identifier{Value: []byte("string")}))
	assert.Equal(test, "", ast.ToString(&ast.Root{}))
}

func TestIsReservedType(test *testing.T) {
	for _, reserved := range []string{"int", "STRING", "Bool", "void", "iterable", "self", "static", "mixed", "never"} {
		assert.Assert(test, ast.IsReservedType(reserved), "expected %q reserved", reserved)
	}
	for _, className := range []string{"App\\User", "Closure", "Iterator", "Stringable"} {
		assert.Assert(test, !ast.IsReservedType(className), "expected %q not reserved", className)
	}
}
