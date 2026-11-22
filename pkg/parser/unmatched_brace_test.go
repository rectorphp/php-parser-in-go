package parser_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

// A closing brace with nothing open is a syntax error, not a lexer panic.
func TestParse_UnmatchedClosingBrace(test *testing.T) {
	for _, grammar := range []string{"8.4", "7.4", "5.6"} {
		for _, source := range []string{"<?php }", "<?php } }", "<?php { } }"} {
			ver, err := version.New(grammar)
			if err != nil {
				test.Fatal(err)
			}
			var parserErrors []*errors.Error
			config := conf.Config{
				Version:          ver,
				ErrorHandlerFunc: func(parserError *errors.Error) { parserErrors = append(parserErrors, parserError) },
			}
			_, _ = parser.Parse([]byte(source), config)
			if len(parserErrors) == 0 {
				test.Errorf("%s %q: expected a syntax error", grammar, source)
			}
		}
	}
}
