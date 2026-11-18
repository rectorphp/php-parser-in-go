package formatter_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/formatter"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// An if whose body ends with an if without an else gets its body braced
// when it has an else or an elseif itself: printed without braces, the
// else would belong to the inner if.
func TestFormatter_BracesDanglingElse(test *testing.T) {
	echo := func(value string) ast.Vertex {
		return &ast.StmtEcho{Exprs: []ast.Vertex{&ast.ScalarLnumber{Value: []byte(value)}}}
	}
	variable := func(name string) ast.Vertex {
		return &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("$" + name)}}
	}
	inner := &ast.StmtIf{Cond: variable("a"), Stmt: echo("1")}
	loop := &ast.StmtWhile{Cond: variable("b"), Stmt: &ast.StmtIf{Cond: variable("c"), Stmt: echo("2")}}
	root := &ast.Root{Stmts: []ast.Vertex{
		&ast.StmtIf{Cond: variable("p"), Stmt: inner, Else: &ast.StmtElse{Stmt: echo("3")}},
		&ast.StmtIf{
			Cond:   variable("q"),
			Stmt:   echo("4"),
			ElseIf: []ast.Vertex{&ast.StmtElseIf{Cond: variable("r"), Stmt: loop}},
			Else:   &ast.StmtElse{Stmt: echo("5")},
		},
	}}
	root.Accept(formatter.NewFormatter())
	var buffer bytes.Buffer
	root.Accept(printer.NewPrinter(&buffer))

	// Without the open tag, and with elseif spaced either way.
	actual := strings.TrimLeft(strings.TrimPrefix(buffer.String(), "<?php"), " \n")
	actual = strings.ReplaceAll(actual, "elseif (", "elseif(")
	expected := "if ($p) {\n    if ($a) echo 1;\n} else echo 3;\nif ($q) echo 4; elseif($r) {\n    while($b) if ($c) echo 2;\n} else echo 5;"
	if actual != expected {
		test.Errorf("\nexpected: %q\nactual: %q\n", expected, actual)
	}
}
