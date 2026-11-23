package formatter_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/formatter"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// formatHeredoc parses src, lays it out from scratch and prints it,
// without the open tag the formatter writes.
func formatHeredoc(test *testing.T, src string) string {
	test.Helper()
	ver, err := version.New("8.4")
	if err != nil {
		test.Fatal(err)
	}
	root, err := parser.Parse([]byte(src), conf.Config{Version: ver})
	if err != nil {
		test.Fatal(err)
	}
	root.Accept(formatter.NewFormatter())
	var buffer bytes.Buffer
	root.(*ast.Root).Accept(printer.NewPrinter(&buffer))
	return strings.TrimLeft(strings.TrimPrefix(buffer.String(), "<?php"), " \n")
}

// A nowdoc stays one: as a heredoc, PHP would interpolate its body.
func TestFormatter_KeepsNowdoc(test *testing.T) {
	actual := formatHeredoc(test, "<?php $a = <<<'TXT'\nval $x \\t end\nTXT;\n")
	expected := "$a = <<<'TXT'\nval $x \\t end\nTXT;\n"
	if actual != expected {
		test.Errorf("\nexpected: %q\nactual: %q\n", expected, actual)
	}
}

// A heredoc keeps its label, and takes another where its body holds a line
// that would end it: here EOT.
func TestFormatter_KeepsHeredocLabel(test *testing.T) {
	actual := formatHeredoc(test, "<?php $a = <<<TXT\nEOT; x\nTXT;\n$b = <<<TXT\nTXT1\nTXT;\n")
	expected := "$a = <<<TXT\nEOT; x\nTXT;\n$b = <<<TXT\nTXT1\nTXT;\n"
	if actual != expected {
		test.Errorf("\nexpected: %q\nactual: %q\n", expected, actual)
	}
}

// Before PHP 7.3 nothing but a semicolon may follow a heredoc's closing
// label on its line.
func TestFormatter_HeredocArgumentEndsLine(test *testing.T) {
	actual := formatHeredoc(test, "<?php echo strtoupper(<<<EOT\nabc\nEOT\n), 1;")
	expected := "echo strtoupper(<<<EOT\nabc\nEOT\n), 1;"
	if actual != expected {
		test.Errorf("\nexpected: %q\nactual: %q\n", expected, actual)
	}
}
