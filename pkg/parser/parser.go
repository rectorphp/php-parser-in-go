package parser

import (
	"errors"
	"github.com/rectorphp/php-parser-in-go/internal/php8"

	"github.com/rectorphp/php-parser-in-go/internal/php5"
	"github.com/rectorphp/php-parser-in-go/internal/php7"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

var (
	// ErrVersionOutOfRange is returned if the version is not supported
	ErrVersionOutOfRange = errors.New("the version is out of supported range")

	php5RangeStart = &version.Version{Major: 5}
	php5RangeEnd   = &version.Version{Major: 5, Minor: 6}

	php7RangeStart = &version.Version{Major: 7}
	php7RangeEnd   = &version.Version{Major: 7, Minor: 4}

	php8RangeStart = &version.Version{Major: 8}
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() ast.Vertex
}

func Parse(source []byte, config conf.Config) (ast.Vertex, error) {
	var parser Parser

	if config.Version == nil {
		config.Version = php7RangeEnd
	}

	if config.Version.InRange(php5RangeStart, php5RangeEnd) {
		lexer := scanner.NewLexer(source, config)
		parser = php5.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	if config.Version.InRange(php7RangeStart, php7RangeEnd) {
		lexer := scanner.NewLexer(source, config)
		parser = php7.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	// PHP 8 and higher: the php8 grammar is the newest supported, so any 8.x or
	// later version is parsed with it.
	if config.Version.GreaterOrEqual(php8RangeStart) {
		lexer := php8.NewLexer(source, config)
		parser = php8.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	return nil, ErrVersionOutOfRange
}
