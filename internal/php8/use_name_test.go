package php8_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// A `use` statement carrying a qualified name used to be parsed as a relative
// name — the `namespace\Foo` form — whose parser assumes a nine-character
// `namespace` keyword sits at the front of the token. With no keyword there it
// cut nine characters out of the name itself, so `use Psr\Log\LoggerInterface;`
// printed back as `use Psr\Log\L o ggerInterface;`. Every PHP 8 file carrying a
// qualified `use` was rewritten corrupted.
//
// Printing a parsed file back must reproduce it byte for byte, so these guard
// the shapes a `use` statement can take.
func TestUseStatementRoundTrip(test *testing.T) {
	// built rather than written literally, so no escape in this file's own
	// source can mask a difference in what the parser actually receives
	backslash := string(rune(92))
	name := func(parts ...string) string {
		return strings.Join(parts, backslash)
	}

	tests := []struct {
		name   string
		source string
	}{
		{
			name:   "qualified name, first segment longer than the namespace keyword",
			source: "<?php\nuse " + name("Symfony", "Bundle", "FrameworkBundle", "Controller", "AbstractController") + ";",
		},
		{
			name:   "qualified name, first segment shorter than the namespace keyword",
			source: "<?php\nuse " + name("A", "B", "C") + ";",
		},
		{
			name:   "qualified name of exactly two segments",
			source: "<?php\nuse " + name("Psr", "Log") + ";",
		},
		{
			name:   "unqualified name",
			source: "<?php\nuse Foo;",
		},
		{
			name:   "fully qualified name",
			source: "<?php\nuse " + backslash + name("Foo", "Bar") + ";",
		},
		{
			name:   "aliased qualified name",
			source: "<?php\nuse " + name("Foo", "Bar") + " as Baz;",
		},
		{
			name:   "group use",
			source: "<?php\nuse " + name("Foo", "Bar") + backslash + "{Baz, Quux};",
		},
		{
			name:   "function use",
			source: "<?php\nuse function " + name("Foo", "Bar", "baz") + ";",
		},
		{
			name:   "const use",
			source: "<?php\nuse const " + name("Foo", "Bar", "BAZ") + ";",
		},
		{
			name:   "namespace declaration with a qualified name",
			source: "<?php\nnamespace " + name("App", "Controller") + ";",
		},
		{
			name:   "relative name keeps working",
			source: "<?php\n$value = namespace" + backslash + name("Foo", "bar") + "();",
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			rootNode, err := parser.Parse([]byte(testCase.source), conf.Config{
				Version: &version.Version{Major: 8, Minor: 2},
			})
			if err != nil {
				test.Fatalf("parse failed: %v", err)
			}

			var output bytes.Buffer
			rootNode.Accept(printer.NewPrinter(&output))

			if output.String() != testCase.source {
				test.Errorf("round trip changed the source\n in: %q\nout: %q", testCase.source, output.String())
			}
		})
	}
}
