package php7

import (
	"github.com/rectorphp/php-parser-in-go/internal/position"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

// Parser structure
type Parser struct {
	Lexer          *scanner.Lexer
	currentToken   *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *position.Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *scanner.Lexer, config conf.Config) *Parser {
	return &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
		builder:        position.NewBuilder(),
	}
}

func (parser *Parser) Lex(lval *yySymType) int {
	token := parser.Lexer.Lex()

	parser.currentToken = token
	lval.token = token

	return int(token.ID)
}

func (parser *Parser) Error(message string) {
	if parser.errHandlerFunc == nil {
		return
	}

	parser.errHandlerFunc(errors.NewError(message, parser.currentToken.Position))
}

// Parse the php7 Parser entrypoint
func (parser *Parser) Parse() int {
	parser.rootNode = nil

	return yyParse(parser)
}

// GetRootNode returns root node
func (parser *Parser) GetRootNode() ast.Vertex {
	return parser.rootNode
}
