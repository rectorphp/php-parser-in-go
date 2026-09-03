package php8_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// parseShebang parses PHP 8 source and returns the root plus every parser error
// it collected.
func parseShebang(source string) (ast.Vertex, []*errors.Error) {
	var parserErrors []*errors.Error
	config := conf.Config{
		Version:          &version.Version{Major: 8, Minor: 2},
		ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
	}
	lexer := php8.NewLexer([]byte(source), config)
	parser := php8.NewParser(lexer, config)
	parser.Parse()
	return parser.GetRootNode(), parserErrors
}

// printSource prints the parsed root back to source.
func printSource(root ast.Vertex) string {
	output := bytes.NewBufferString("")
	root.Accept(printer.NewPrinter(output))
	return output.String()
}

// A CLI entry point opens with a shebang line before the open tag. The scanner
// keeps it as free floating, and the printer must not mistake it for HTML that
// needs an open tag of its own.
func TestShebangPrintsBackUnchanged(test *testing.T) {
	source := `#!/usr/bin/env php
<?php

namespace App\Bin;

use Symfony\Component\Console\Application;

$application = new Application($kernel);

return $application;
`

	root, parserErrors := parseShebang(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if printed := printSource(root); printed != source {
		test.Errorf("printed output differs from source:\n%s", printed)
	}
}

func TestShebangWithoutTrailingCodePrintsBackUnchanged(test *testing.T) {
	source := "#!/usr/bin/env php\n<?php\necho 1;\n"

	root, parserErrors := parseShebang(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if printed := printSource(root); printed != source {
		test.Errorf("printed output differs from source:\n%q", printed)
	}
}

// Leading HTML is not a shebang: it still gets the open tag treatment.
func TestLeadingInlineHtmlStillPrintsBackUnchanged(test *testing.T) {
	source := "<h1>Title</h1>\n<?php\necho 1;\n"

	root, parserErrors := parseShebang(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if printed := printSource(root); printed != source {
		test.Errorf("printed output differs from source:\n%q", printed)
	}
}

// A `#` comment inside PHP code must not be confused with a shebang.
func TestHashCommentInsideCodePrintsBackUnchanged(test *testing.T) {
	source := "<?php\n# a hash comment\necho 1;\n"

	root, parserErrors := parseShebang(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}

	if printed := printSource(root); printed != source {
		test.Errorf("printed output differs from source:\n%q", printed)
	}
}
