package printer_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func TestPrinterPrintFile(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer)
	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Foo")},
					},
				},
			},
			&ast.StmtClass{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Bar")},
					},
				},
				Extends: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Baz")},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						Name:      &ast.Identifier{Value: []byte("greet")},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtEcho{
									Exprs: []ast.Vertex{
										&ast.ScalarString{Value: []byte("'Hello world'")},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `<?php namespace Foo;abstract class Bar extends Baz{public function greet(){echo'Hello world';}}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFileInlineHtml(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ScalarString{
						Value: []byte(`"a"`),
					},
				},
			},
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ScalarString{
						Value: []byte(`"b"`),
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `<div>HTML</div><?php echo"a";?><div>HTML</div><?php echo"b";`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// node

func TestPrinterPrintIdentifier(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Identifier{
		Value: []byte("test"),
	}
	node.Accept(printerValue)

	expected := `test`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintParameter(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Parameter{
		Type: &ast.NameFullyQualified{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		VariadicTkn: &token.Token{
			Value: []byte("..."),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		DefaultValue: &ast.ScalarString{
			Value: []byte("'default'"),
		},
	}
	node.Accept(printerValue)

	expected := "\\Foo...$var='default'"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNullable(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Nullable{
		Expr: &ast.Parameter{
			Type: &ast.NameFullyQualified{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("Foo"),
					},
				},
			},
			AmpersandTkn: &token.Token{
				Value: []byte("&"),
			},
			Var: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			DefaultValue: &ast.ScalarString{
				Value: []byte("'default'"),
			},
		},
	}
	node.Accept(printerValue)

	expected := "?\\Foo&$var='default'"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintArgument(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Argument{
		VariadicTkn: &token.Token{
			Value: []byte("..."),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := "...$var"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrinterPrintArgumentByRef(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Argument{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := "&$var"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAttributeGroup(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.AttributeGroup{
		Attrs: []ast.Vertex{
			&ast.Attribute{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("FooAttribute"),
						},
					},
				},
			},
			&ast.Attribute{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("BarAttribute"),
						},
					},
				},
				Args: []ast.Vertex{
					&ast.Argument{
						Expr: &ast.ExprVariable{
							Name: &ast.Identifier{
								Value: []byte("$arg"),
							},
						},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := "#[FooAttribute,BarAttribute($arg)]"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// name

func TestPrinterPrintNameNamePart(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.NamePart{
		Value: []byte("foo"),
	}
	node.Accept(printerValue)

	expected := "foo"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameName(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	node.Accept(printerValue)

	expected := "Foo\\Bar"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameFullyQualified(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.NameFullyQualified{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	node.Accept(printerValue)

	expected := "\\Foo\\Bar"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameRelative(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.NameRelative{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	node.Accept(printerValue)

	expected := "namespace\\Foo\\Bar"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// scalar

func TestPrinterPrintScalarLNumber(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarLnumber{
		Value: []byte("1"),
	}
	node.Accept(printerValue)

	expected := "1"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarDNumber(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarDnumber{
		Value: []byte(".1"),
	}
	node.Accept(printerValue)

	expected := ".1"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarString(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarString{
		Value: []byte("'hello world'"),
	}
	node.Accept(printerValue)

	expected := `'hello world'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsedStringPart(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarEncapsedStringPart{
		Value: []byte("hello world"),
	}
	node.Accept(printerValue)

	expected := `hello world`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsed(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world")},
		},
	}
	node.Accept(printerValue)

	expected := `"hello $var world"`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarHeredoc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world\n")},
		},
	}
	node.Accept(printerValue)

	expected := `<<<EOT
hello $var world
EOT`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarMagicConstant(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ScalarMagicConstant{
		Value: []byte("__DIR__"),
	}
	node.Accept(printerValue)

	if buffer.String() != `__DIR__` {
		test.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrinterPrintAssign(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssign{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignReference{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a=&$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignBitwiseAnd{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a&=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignBitwiseOr{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a|=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignBitwiseXor{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a^=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignCoalesce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignCoalesce{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a??=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignConcat(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignConcat{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a.=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignDiv(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignDiv{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a/=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignMinus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a-=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignMod{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a%=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMul(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignMul{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a*=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignPlus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a+=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignPow{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a**=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftLeft(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignShiftLeft{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a<<=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftRight(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprAssignShiftRight{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a>>=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// binary

func TestPrinterPrintBinaryBitwiseAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryBitwiseAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a&$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryBitwiseOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a|$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryBitwiseXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a^$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryBooleanAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a&&$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryBooleanOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a||$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryCoalesce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryCoalesce{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a??$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryConcat(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryConcat{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a.$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryDiv(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryDiv{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a/$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a==$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreaterOrEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryGreaterOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a>=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreater(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryGreater{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a>$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryIdentical(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a===$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalAnd(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryLogicalAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a and$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalOr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryLogicalOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a or$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalXor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryLogicalXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a xor$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryMinus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a-$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryMod{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a%$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMul(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryMul{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a*$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryNotEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a!=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotIdentical(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryNotIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a!==$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryPlus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a+$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryPow{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a**$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftLeft(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryShiftLeft{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a<<$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftRight(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinaryShiftRight{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a>>$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmallerOrEqual(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinarySmallerOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a<=$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmaller(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinarySmaller{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a<$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySpaceship(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBinarySpaceship{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a<=>$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// cast

func TestPrinterPrintArray(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastArray{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(array)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBool(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastBool{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(bool)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDouble(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastDouble{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(float)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInt(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastInt{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(int)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintObject(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastObject{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(object)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintString(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastString{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(string)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprCastUnset{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `(unset)$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// expr

func TestPrinterPrintExprArrayDimFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Dim: &ast.ScalarLnumber{Value: []byte("1")},
	}
	node.Accept(printerValue)

	expected := `$var[1]`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemWithKey(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArrayItem{
		Key: &ast.ScalarString{Value: []byte("'Hello'")},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	node.Accept(printerValue)

	expected := `'Hello'=>$world`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItem(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArrayItem{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	node.Accept(printerValue)

	expected := `$world`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItem_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArrayItem{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	node.Accept(printerValue)

	expected := `&$world`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemUnpack(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArrayItem{
		EllipsisTkn: &token.Token{
			Value: []byte("..."),
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	node.Accept(printerValue)

	expected := `...$world`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArray(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArray{
		ArrayTkn: &token.Token{
			Value: []byte("array"),
		},
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `array('Hello'=>$world,2=>&$var,$var)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBitwiseNot(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `~$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBooleanNot(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `!$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBracket(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `!$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClassConstFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprClassConstFetch{
		Class: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Const: &ast.Identifier{
			Value: []byte("CONST"),
		},
	}
	node.Accept(printerValue)

	expected := `$var::CONST`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClone(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprClone{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `clone$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprClosureUse{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
	}
	node.Accept(printerValue)

	expected := `$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprClosureUse{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
	}
	node.Accept(printerValue)

	expected := `&$foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosure(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprClosure{
		StaticTkn: &token.Token{
			Value: []byte("static"),
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.ExprClosureUse{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprClosureUse{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `static function&($var)use(&$a,$b):\Foo{$a;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrowFunction(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtExpression{
		Expr: &ast.ExprArrowFunction{
			StaticTkn: &token.Token{
				Value: []byte("static"),
			},
			AmpersandTkn: &token.Token{
				Value: []byte("&"),
			},
			Params: []ast.Vertex{
				&ast.Parameter{
					AmpersandTkn: &token.Token{
						Value: []byte("&"),
					},
					Var: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$var")},
					},
				},
			},
			ReturnType: &ast.NameFullyQualified{
				Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
			},
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `static fn&(&$var):\Foo=>$a;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprConstFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprConstFetch{
		Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}},
	}
	node.Accept(printerValue)

	expected := "null"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprEmpty{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `empty($var)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrettyPrinterrorSuppress(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprErrorSuppress{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `@$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEval(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprEval{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `eval($var)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExit(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprExit{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `exit$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDie(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprExit{
		ExitTkn: &token.Token{
			Value: []byte("die"),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `die$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFunctionCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprFunctionCall{
		Function: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `$var(&$a,...$b,$c)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInclude(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprInclude{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	node.Accept(printerValue)

	expected := `include'path'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIncludeOnce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprIncludeOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	node.Accept(printerValue)

	expected := `include_once'path'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInstanceOf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprInstanceOf{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Class: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
	}
	node.Accept(printerValue)

	expected := `$var instanceof Foo`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIsset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `isset($a,$b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							},
						},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `list($a,list($b,$c))`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintMethodCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
		Method: &ast.Identifier{Value: []byte("bar")},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `$foo->bar($a,$b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNew(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `new Foo($a,$b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostDec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPostDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `$var--`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostInc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPostInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `$var++`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreDec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPreDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `--$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreInc(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPreInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `++$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPrint(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPrint{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `print$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
		Prop: &ast.Identifier{Value: []byte("bar")},
	}
	node.Accept(printerValue)

	expected := `$foo->bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequire(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprRequire{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	node.Accept(printerValue)

	expected := `require'path'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequireOnce(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprRequireOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	node.Accept(printerValue)

	expected := `require_once'path'`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShellExec(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$world")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte("!")},
		},
	}
	node.Accept(printerValue)

	expected := "`hello $world!`"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprShortArray(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `['Hello'=>$world,2=>&$var,$var]`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShortList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprList{
		OpenBracketTkn: &token.Token{
			Value: []byte("["),
		},
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							},
						},
					},
				},
			},
		},
		CloseBracketTkn: &token.Token{
			Value: []byte("]"),
		},
	}
	node.Accept(printerValue)

	expected := `[$a,list($b,$c)]`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticCall(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprStaticCall{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Call:  &ast.Identifier{Value: []byte("bar")},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `Foo::bar($a,$b)`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticPropertyFetch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprStaticPropertyFetch{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Prop: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$bar")},
		},
	}
	node.Accept(printerValue)

	expected := `Foo::$bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernary(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	node.Accept(printerValue)

	expected := `$a?:$b`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernaryFull(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		IfTrue: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$c")},
		},
	}
	node.Accept(printerValue)

	expected := `$a?$b:$c`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryMinus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `-$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryPlus(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `+$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintVariable(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprVariable{
		DollarTkn: &token.Token{
			Value: []byte("$"),
		},
		OpenCurlyBracketTkn: &token.Token{
			Value: []byte("{"),
		},
		Name: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		CloseCurlyBracketTkn: &token.Token{
			Value: []byte("}"),
		},
	}
	node.Accept(printerValue)

	expected := `${$var}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFrom(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `yield from$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYield(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprYield{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `yield$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFull(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.ExprYield{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `yield$k=>$var`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// stmt

func TestPrinterPrintAltElseIf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `elseif($a):$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseIfEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{},
	}
	node.Accept(printerValue)

	expected := `elseif($a):`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElse{
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `else:$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElse{
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{},
	}
	node.Accept(printerValue)

	expected := `else:`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltFor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$c")},
			},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `for($a;$b;$c):$d;endfor;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$key")},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$val")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `foreach($var as$key=>$val):$d;endforeach;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$key")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$val")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `foreach($var as$key=>&$val):$d;endforeach;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltIf(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
				ColonTkn: &token.Token{
					Value: []byte(":"),
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$b")},
						}},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
				ColonTkn: &token.Token{
					Value: []byte(":"),
				},
				Stmt: &ast.StmtStmtList{},
			},
		},
		Else: &ast.StmtElse{
			ColonTkn: &token.Token{
				Value: []byte(":"),
			},
			Stmt: &ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `if($a):$d;elseif($b):$b;elseif($c):else:$b;endif;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAltSwitch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'a'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					}},
				},
			},
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'b'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `switch($var):case'a':$a;case'b':$b;endswitch;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltWhile(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `while($a):$b;endwhile;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtBreak(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtBreak{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
		},
	}
	node.Accept(printerValue)

	expected := "break 1;"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCase(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `case$a:$a;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCaseEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{},
	}
	node.Accept(printerValue)

	expected := "case$a:"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCatch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtCatch{
		Types: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Exception")}}},
			&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("RuntimeException")}}},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$e")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `catch(Exception|\RuntimeException$e){$a;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassMethod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("int")}}}},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `public function&foo(?int&$a=null,...$b):void{$a;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAbstractClassMethod(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("int")}}}},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtNop{},
	}
	node.Accept(printerValue)

	expected := `public static function&foo(?int&$a=null,...$b):void;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClass(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtClass{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
		Name:      &ast.Identifier{Value: []byte("Foo")},
		Extends:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		Implements: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Quuz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassConstList{
				Modifiers: []ast.Vertex{
					&ast.Identifier{Value: []byte("public")},
					&ast.Identifier{Value: []byte("static")},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `abstract class Foo extends Bar implements Baz,Quuz{public static const FOO='bar';}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAnonymousClass(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtClass{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		Extends: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		Implements: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Quuz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassConstList{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `abstract class($a,$b)extends Bar implements Baz,Quuz{public const FOO='bar';}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassConstList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtClassConstList{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("BAR")},
				Expr: &ast.ScalarString{Value: []byte("'b'")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `public const FOO='a',BAR='b';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("BAR")},
				Expr: &ast.ScalarString{Value: []byte("'b'")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `const FOO='a',BAR='b';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstant(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("FOO")},
		Expr: &ast.ScalarString{Value: []byte("'BAR'")},
	}
	node.Accept(printerValue)

	expected := "FOO='BAR'"
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtContinue(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtContinue{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
		},
	}
	node.Accept(printerValue)

	expected := `continue 1;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareStmts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `declare(FOO='bar'){;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareExpr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	node.Accept(printerValue)

	expected := `declare(FOO='bar')'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtNop{},
	}
	node.Accept(printerValue)

	expected := `declare(FOO='bar');`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalut(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `default:$a;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalutEmpty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDefault{
		Stmts: []ast.Vertex{},
	}
	node.Accept(printerValue)

	expected := `default:`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_Expression(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `do$a;while(1);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_StmtList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `do{$a;}while(1);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoHtmlState(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer)
	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					},
					&ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `<?php echo$a,$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoPhpState(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `echo$a,$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfStmts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `elseif($a){;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfExpr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	node.Accept(printerValue)

	expected := `elseif($a)'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
	}
	node.Accept(printerValue)

	expected := `elseif($a);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseStmts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `else{;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseExpr(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElse{
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	node.Accept(printerValue)

	expected := `else'bar';`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtElse{
		Stmt: &ast.StmtNop{},
	}
	node.Accept(printerValue)

	expected := `else;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExpression(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtExpression{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
	}
	node.Accept(printerValue)

	expected := `$a;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFinally(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtFinally{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}
	node.Accept(printerValue)

	expected := `finally{;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFor(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$c")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$d")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$e")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$f")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `for($a,$b;$c,$d;$e,$f){;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$v")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `foreach($a as$k=>$v){;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach_Reference(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$v")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `foreach($a as$k=>&$v){;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFunction(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtFunction{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}
	node.Accept(printerValue)

	expected := `function&foo(&$var):\Foo{;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGlobal(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `global$a,$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGoto(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtGoto{
		Label: &ast.Identifier{Value: []byte("FOO")},
	}
	node.Accept(printerValue)

	expected := `goto FOO;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintHaltCompiler(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtHaltCompiler{}
	node.Accept(printerValue)

	expected := `__halt_compiler();`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfExpression(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{
							Expr: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$d")},
							},
						},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$e")},
				},
				Stmt: &ast.StmtNop{},
			},
		},
		Else: &ast.StmtElse{
			Stmt: &ast.StmtExpression{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$f")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `if($a)$b;elseif($c){$d;}elseif($e);else$f;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfStmtList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{
					Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `if($a){$b;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
	}
	node.Accept(printerValue)

	expected := `if($a);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInlineHtml(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{
				Value: []byte("test"),
			},
		},
	}
	node.Accept(printerValue)

	expected := `test`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInterface(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtInterface{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Extends: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassMethod{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Name:      &ast.Identifier{Value: []byte("foo")},
				Params:    []ast.Vertex{},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `interface Foo extends Bar,Baz{public function foo(){$a;}}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintLabel(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtLabel{
		Name: &ast.Identifier{Value: []byte("FOO")},
	}
	node.Accept(printerValue)

	expected := `FOO:`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespace(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtNamespace{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
	}
	node.Accept(printerValue)

	expected := `namespace Foo;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespaceWithStmts(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtNamespace{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `namespace Foo{$a;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNop(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtNop{}
	node.Accept(printerValue)

	expected := `;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		Type: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `public static Foo$a='a',$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintProperty(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtProperty{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	node.Accept(printerValue)

	expected := `$a=1`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReturn(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtReturn{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	node.Accept(printerValue)

	expected := `return 1;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticVar(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	node.Accept(printerValue)

	expected := `$a=1`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStatic(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `static$a,$b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			}},
		},
	}
	node.Accept(printerValue)

	expected := `{$a;$b;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtListNested(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							}},
						},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `{$a;{$b;{$c;}}}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtSwitch(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'a'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					}},
				},
			},
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'b'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `switch($var){case'a':$a;case'b':$b;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtThrow(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtThrow{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	node.Accept(printerValue)

	expected := `throw$var;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUseAlias(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTraitUseAlias{
		Trait:    &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Method:   &ast.Identifier{Value: []byte("a")},
		Modifier: &ast.Identifier{Value: []byte("public")},
		Alias:    &ast.Identifier{Value: []byte("b")},
	}
	node.Accept(printerValue)

	expected := `Foo::a as public b;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUsePrecedence(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTraitUsePrecedence{
		Trait:  &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Method: &ast.Identifier{Value: []byte("a")},
		Insteadof: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
		},
	}
	node.Accept(printerValue)

	expected := `Foo::a insteadof Bar,Baz;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		},
	}
	node.Accept(printerValue)

	expected := `use Foo,Bar;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitAdaptations(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		},
		Adaptations: []ast.Vertex{
			&ast.StmtTraitUseAlias{
				Trait:  &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Method: &ast.Identifier{Value: []byte("a")},
				Alias:  &ast.Identifier{Value: []byte("b")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `use Foo,Bar{Foo::a as b;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTrait(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTrait{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtClassMethod{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Name:      &ast.Identifier{Value: []byte("foo")},
				Params:    []ast.Vertex{},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
			},
		},
	}
	node.Accept(printerValue)

	expected := `trait Foo{public function foo(){$a;}}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTry(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
		Catches: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Exception")}}},
					&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("RuntimeException")}}},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$e")},
				},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
		Finally: &ast.StmtFinally{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	node.Accept(printerValue)

	expected := `try{$a;}catch(Exception|\RuntimeException$e){$b;}finally{;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUnset(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	node.Accept(printerValue)

	expected := `unset($a,$b);`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtUseList{
		Type: &ast.Identifier{Value: []byte("function")},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUse{
				Use: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `use function Foo as Bar,Baz;`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGroupUse(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtGroupUseList{
		Type:   &ast.Identifier{Value: []byte("function")},
		Prefix: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUse{
				Use: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `use function Foo\{Foo as Bar,Baz};`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUseDeclaration(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtUse{
		Type:  &ast.Identifier{Value: []byte("function")},
		Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Alias: &ast.Identifier{Value: []byte("Bar")},
	}
	node.Accept(printerValue)

	expected := `function Foo as Bar`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintWhileStmtList(test *testing.T) {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer).WithState(printer.PrinterStatePHP)
	node := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	node.Accept(printerValue)

	expected := `while($a){$a;}`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
