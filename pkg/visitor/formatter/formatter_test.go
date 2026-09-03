package formatter_test

import (
	"bytes"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/formatter"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

func TestFormatter_Root(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter()
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	expected := `<?php 

;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Nullable(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Nullable{
		Expr: &ast.Identifier{
			Value: []byte("array"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `?array`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Parameter{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Parameter{
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `&$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Variadic(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Parameter{
		VariadicTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `...$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Type(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Parameter{
		Type: &ast.Identifier{
			Value: []byte("array"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `array $var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Default(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Parameter{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		DefaultValue: &ast.ScalarString{
			Value: []byte("'default'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$var = 'default'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Identifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Identifier{
		Value: []byte("foo"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Argument{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Argument{
		AmpersandTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `&$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument_Variadic(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Argument{
		VariadicTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `...$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtBreak(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtBreak{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `break;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtBreak_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtBreak{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `break $var;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Case(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `case $var:
        ;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Catch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtCatch{
		Types: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$baz"),
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `catch (foo | bar $baz) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `class foo {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Modifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClass{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("final"),
			},
		},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `final class foo {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Anonymous(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `class foo($a, $b) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Extends(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Extends: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `class foo extends bar {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Implements(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Implements: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `class foo implements bar {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtClassConstList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `const foo = 'foo', bar = 'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtClassConstList_Modifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassConstList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
		},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `public const foo = 'foo', bar = 'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtNop{},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo() ;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Modifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
		},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `public function foo() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassMethod{
		AmpersandTkn: &token.Token{},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function &foo() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Params(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo($a, $b) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_ReturnType(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo(): bar {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtConstList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `const foo = 'foo', bar = 'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtConstant(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtConstant{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Expr: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo = 'bar'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtContinue(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtContinue{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `continue;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtContinue_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtContinue{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `continue $var;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDeclare(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtDeclare{
		Stmt: &ast.StmtNop{},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `declare() ;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDeclare_Consts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `declare(foo = 'foo', bar = 'bar') {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDefault(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `default:
        ;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDo(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtDo{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `do {
        ;
    } while($var);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtEcho(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ScalarString{
				Value: []byte("'foo'"),
			},
			&ast.ScalarString{
				Value: []byte("'bar'"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `echo 'foo', 'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtElse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `else {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtElseIf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `elseif($var) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtExpression(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtExpression{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$var;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFinally(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFinally{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `finally {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `for($foo, $bar; $foo, $bar; $foo, $bar) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foreach($foo as $val) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foreach($foo as &$val) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach_Key(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$key"),
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foreach($foo as $key => $val) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFunction{
		AmpersandTkn: &token.Token{},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function &foo() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_Params(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo($a, $b) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_ReturnType(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo(): bar {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGlobal(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `global $a, $b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGoto(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtGoto{
		Label: &ast.Identifier{
			Value: []byte("FOO"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `goto FOO;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtHaltCompiler(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtHaltCompiler{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `__halt_compiler();`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `if ($foo) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf_ElseIf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$baz"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `if ($foo) {
        ;
    } elseif($bar) {
        ;
    } elseif($baz) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf_Else(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		Else: &ast.StmtElse{
			Stmt: &ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `if ($foo) {
        ;
    } else {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInlineHtml(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
					&ast.StmtInlineHtml{
						Value: []byte("<div></div>"),
					},
					&ast.StmtEcho{
						Exprs: []ast.Vertex{
							&ast.ExprVariable{
								Name: &ast.Identifier{
									Value: []byte("$foo"),
								},
							},
						},
					},
					&ast.StmtInlineHtml{
						Value: []byte("<div></div>"),
					},
					&ast.StmtNop{},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter()
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	expected := `<?php 

{
    ;?><div></div><?php 
    echo $foo;?><div></div><?php 
    ;
}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInterface(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtInterface{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `interface foo {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInterface_Extends(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtInterface{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Extends: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `interface foo extends bar {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtLabel(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtLabel{
		Name: &ast.Identifier{
			Value: []byte("FOO"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `FOO:`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNamespace_Name(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtNamespace{
		Name: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `namespace foo;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNamespace_Stmts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `namespace {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtNop{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtProperty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtProperty{
		Var: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtProperty_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtProperty{
		Var: &ast.Identifier{
			Value: []byte("$foo"),
		},
		Expr: &ast.Identifier{
			Value: []byte("$bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo = $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtPropertyList{
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo, $bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList_Modifiers(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
			&ast.Identifier{
				Value: []byte("static"),
			},
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `public static $foo, $bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList_Type(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtPropertyList{
		Type: &ast.Identifier{
			Value: []byte("array"),
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `array $foo, $bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtReturn(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtReturn{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `return;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtReturn_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtReturn{
		Expr: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `return $foo;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStatic(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `static $a, $b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStaticVar(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStaticVar_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.Identifier{
			Value: []byte("$bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo = $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStmtList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtNop{},
						},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter()
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	expected := `<?php 

{
    {
        ;
    }
}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtSwitch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtBreak{},
				},
			},
			&ast.StmtDefault{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `switch($foo) {
        case 'bar':
            break;
        default:
            ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtThrow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtThrow{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `throw $foo;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTrait(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTrait{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `trait foo {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use foo, bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUse_Adaptations(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Adaptations: []ast.Vertex{
			&ast.StmtTraitUseAlias{
				Method: &ast.Identifier{
					Value: []byte("foo"),
				},
				Alias: &ast.Identifier{
					Value: []byte("baz"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use foo, bar {
        foo as baz;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo as public bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Trait(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUseAlias{
		Trait: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
		Alias: &ast.Identifier{
			Value: []byte("baz"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::bar as public baz;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Alias(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo as bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Modifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo as public;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUsePrecedence(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUsePrecedence{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Insteadof: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("baz"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo insteadof bar, baz;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUsePrecedence_Trait(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTraitUsePrecedence{
		Trait: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Insteadof: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("baz"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::bar insteadof baz;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `try {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry_Catch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		Catches: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Value: []byte("foo"),
							},
						},
					},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Value: []byte("foo"),
							},
						},
					},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `try {
        ;
    } catch (foo $bar) {
        ;
    } catch (foo $bar) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry_Finally(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		Finally: &ast.StmtFinally{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `try {
        ;
    } finally {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUnset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `unset($a, $b);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUseList{
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("foo"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use foo, bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUse_Type(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUseList{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("foo"),
						},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use function foo;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGroupUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtGroupUseList{
		Prefix: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("baz"),
						},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use foo\{bar, baz};`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGroupUse_Type(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtGroupUseList{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Prefix: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("baz"),
						},
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `use function foo\{bar, baz};`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUse{
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration_Type(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUse{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration_Alias(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtUse{
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo as bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtWhile(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtNop{},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `while($foo) ;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArray(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `array($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayDimFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Dim: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo[$bar]`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrayItem{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem_Key(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrayItem{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo => $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem_Variadic(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrayItem{
		EllipsisTkn: &token.Token{},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `...$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrowFunction{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `fn() => $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrowFunction{
		AmpersandTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `fn&() => $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_Params(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrowFunction{
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `fn($a, $b) => $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_ReturnType(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprArrowFunction{
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `fn(): foo => $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBitwiseNot(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `~$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBooleanNot(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `!$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBrackets(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBrackets{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `($foo)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClassConstFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClassConstFetch{
		Class: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Const: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo::bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClone(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClone{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `clone $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosure{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Ref(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosure{
		AmpersandTkn: &token.Token{},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function&() {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Params(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosure{
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function($a, $b) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_ReturnType(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosure{
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function(): foo {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Use(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosure{
		Uses: []ast.Vertex{
			&ast.ExprClosureUse{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$foo"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `function() use($foo) {
        ;
    }`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosureUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosureUse{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$a"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$a`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosureUse_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprClosureUse{
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$a"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `&$a`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprConstFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprConstFetch{
		Const: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("FOO"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `FOO`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprEmpty{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `empty($foo)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprErrorSuppress(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprErrorSuppress{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `@$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprEval(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprEval{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `eval($foo)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprExit(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprExit{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `exit`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprExit_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprExit{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `exit($foo)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprFunctionCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprFunctionCall{
		Function: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo()`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprFunctionCall_Arguments(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprFunctionCall{
		Function: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo($bar)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprInclude(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprInclude{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `include 'foo.php'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprIncludeOnce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprIncludeOnce{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `include_once 'foo.php'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprInstanceOf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprInstanceOf{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo instanceof bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprIsset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `isset($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `list($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo->bar()`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo->{'bar'}()`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall_Arguments(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo->bar($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprNew(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `new foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprNew_Arguments(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `new foo($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPreDec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPreDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `--$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPreInc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPreInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `++$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPostDec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPostDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo--`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPostInc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPostInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo++`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPrint(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPrint{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `print $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPropertyFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Prop: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo->bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPropertyFetch_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Prop: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo->{'bar'}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprRequire(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprRequire{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `require 'foo.php'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprRequireOnce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprRequireOnce{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `require_once 'foo.php'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprShellExec{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := "``"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec_Part(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := "`foo`"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec_Parts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := "`foo $bar baz`"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::bar()`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall_Expr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::{'bar'}()`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall_Arguments(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.Identifier{
			Value: []byte("bar"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::bar($a, $b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticPropertyFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprStaticPropertyFetch{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Prop: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo::$bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprTernary(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		IfTrue: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$baz"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ? $bar : $baz`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprTernary_short(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ?: $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprUnaryMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `-$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprUnaryPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `+$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprVariable{
		Name: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable_Variable(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprVariable{
		Name: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable_Expression(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprVariable{
		Name: &ast.ScalarString{
			Value: []byte("'foo'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `${'foo'}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYield(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprYield{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `yield $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYield_Key(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprYield{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `yield $foo => $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYieldFrom(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `yield from $foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssign(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssign{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo = $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignReference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignReference{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo =& $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignBitwiseAnd{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo &= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignBitwiseOr{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo |= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignBitwiseXor{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ^= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignCoalesce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignCoalesce{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ??= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignConcat(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignConcat{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo .= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignDiv(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignDiv{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo /= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignMinus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo -= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignMod{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo %= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMul(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignMul{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo *= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignPlus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo += $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignPow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignPow{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo **= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignShiftLeft(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignShiftLeft{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo <<= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignShiftRight(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprAssignShiftRight{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo >>= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryBitwiseAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo & $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryBitwiseOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo | $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryBitwiseXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ^ $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBooleanAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryBooleanAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo && $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBooleanOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryBooleanOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo || $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryCoalesce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryCoalesce{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ?? $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryConcat(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryConcat{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo . $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryDiv(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryDiv{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo / $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo == $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryGreater(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryGreater{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo > $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryGreaterOrEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryGreaterOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo >= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryIdentical(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo === $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryLogicalAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo and $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryLogicalOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo or $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryLogicalXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo xor $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryMinus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo - $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryMod{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo % $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMul(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryMul{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo * $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryNotEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryNotEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo != $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryNotIdentical(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryNotIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo !== $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryPlus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo + $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryPow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryPow{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo ** $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryShiftLeft(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryShiftLeft{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo << $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryShiftRight(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinaryShiftRight{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo >> $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySmaller(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinarySmaller{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo < $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySmallerOrEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinarySmallerOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo <= $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySpaceship(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprBinarySpaceship{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `$foo <=> $bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastArray(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastArray{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(array)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastBool(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastBool{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(bool)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastDouble(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastDouble{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(float)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastInt(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastInt{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(int)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastObject(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastObject{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(object)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastString(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastString{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(string)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastUnset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ExprCastUnset{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `(unset)$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarDnumber(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarDnumber{
		Value: []byte("1234"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `1234`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsed{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `""`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed_Part(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `"foo"`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed_Parts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `"foo $bar baz"`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringPart(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsedStringPart{
		Value: []byte("foo"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringVar(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsedStringVar{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `${foo}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringVar_Dim(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsedStringVar{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Dim: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `${foo['bar']}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringBrackets(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarEncapsedStringBrackets{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `{$foo}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarHeredoc{}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `<<<EOT
EOT`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc_Part(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo\n"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `<<<EOT
foo
EOT`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc_Parts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz\n"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `<<<EOT
foo $bar baz
EOT`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarLnumber(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarLnumber{
		Value: []byte("1234"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `1234`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarMagicConstant(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarMagicConstant{
		Value: []byte("__DIR__"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `__DIR__`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarString(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.ScalarString{
		Value: []byte("'foo'"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `'foo'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameName(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo\bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameFullyQualified(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.NameFullyQualified{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `\foo\bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameRelative(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.NameRelative{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `namespace\foo\bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameNamePart(test *testing.T) {
	buffer := bytes.NewBufferString("")

	node := &ast.NamePart{
		Value: []byte("foo"),
	}

	formatterValue := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	node.Accept(formatterValue)

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node.Accept(printerValue)

	expected := `foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
