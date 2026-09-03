package php8

import (
	"bytes"

	"github.com/rectorphp/php-parser-in-go/internal/position"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

// enumKeyword is the contextual keyword that opens an enum declaration.
var enumKeyword = []byte("enum")

// Parser structure
type Parser struct {
	Lexer          *Lexer
	currentToken   *token.Token
	peekedToken    *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *position.Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *Lexer, config conf.Config) *Parser {
	return &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
		builder:        position.NewBuilder(),
	}
}

func (parser *Parser) Lex(lval *yySymType) int {
	lexedToken := parser.nextToken()

	parser.currentToken = lexedToken
	lval.token = lexedToken

	// `&` is ambiguous: it starts a by-reference parameter (`Foo &$bar`) or an
	// intersection type (`Foo&Bar $baz`), and one token of lookahead is needed
	// to tell them apart. PHP splits the token in its lexer for the same
	// reason, so the grammar can accept only the right flavour in each place.
	if lexedToken.ID == '&' && !parser.peekIsVarOrVariadic() {
		return int(token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG)
	}

	// `enum` is contextual, not reserved: it only opens a declaration when a
	// name follows it, so `function enum()` or `$foo->enum` keep working.
	if lexedToken.ID == token.T_STRING && bytes.EqualFold(lexedToken.Value, enumKeyword) && parser.peekIsName() {
		return int(token.T_ENUM)
	}

	return int(lexedToken.ID)
}

// nextToken returns the buffered token from a previous peek, or lexes a new one.
func (parser *Parser) nextToken() *token.Token {
	if parser.peekedToken != nil {
		peekedToken := parser.peekedToken
		parser.peekedToken = nil

		return peekedToken
	}

	return parser.Lexer.Lex()
}

// peekIsName reports whether the token after the current one is a plain name,
// buffering it for the next Lex call.
func (parser *Parser) peekIsName() bool {
	if parser.peekedToken == nil {
		parser.peekedToken = parser.Lexer.Lex()
	}

	return parser.peekedToken.ID == token.T_STRING
}

// peekIsVarOrVariadic reports whether the token after the current one starts a
// variable or a variadic parameter, buffering it for the next Lex call.
func (parser *Parser) peekIsVarOrVariadic() bool {
	if parser.peekedToken == nil {
		parser.peekedToken = parser.Lexer.Lex()
	}

	return parser.peekedToken.ID == token.T_VARIABLE || parser.peekedToken.ID == token.T_ELLIPSIS
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

func (parser *Parser) parseNameRelativeToken(nameToken *token.Token) *ast.NameRelative {
	node := &ast.NameRelative{
		Position: parser.builder.NewTokenPosition(nameToken),
	}
	startPos := nameToken.Position.StartPos
	chars := nameToken.Value

	// namespace token

	position := parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + 9

	// the keyword opens the name, so it carries the trivia that preceded the
	// whole token — leave it on a later part and it prints after the separator,
	// turning `namespace\Foo` into `namespace\ Foo`
	node.NsTkn = &token.Token{
		ID:           token.T_NAMESPACE,
		Value:        chars[:9],
		Position:     position,
		FreeFloating: nameToken.FreeFloating,
	}
	nameToken.FreeFloating = nil

	startPos = startPos + 9
	chars = chars[9:]

	// ns separator token

	position = parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + 1

	node.NsSeparatorTkn = &token.Token{
		ID:       token.T_NS_SEPARATOR,
		Value:    chars[:1],
		Position: position,
	}

	startPos = startPos + 1
	chars = chars[1:]

	// parts

	for {
		i := bytes.Index(chars, []byte("\\"))
		if i < 0 {
			break
		}

		position = parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + i

		stringTokenPosition := parser.Lexer.positionPool.Get()
		*stringTokenPosition = *position

		node.Parts = append(node.Parts, &ast.NamePart{
			Position: position,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        chars[:i],
				Position:     stringTokenPosition,
				FreeFloating: nameToken.FreeFloating,
			},
			Value: chars[:i],
		})
		nameToken.FreeFloating = nil
		startPos = startPos + i
		chars = chars[i:]

		position = parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + 1

		node.SeparatorTkns = append(node.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    chars[:1],
			Position: position,
		})
		startPos = startPos + 1
		chars = chars[1:]
	}

	// last part

	position = parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + len(chars)
	stringTokenPosition := parser.Lexer.positionPool.Get()
	*stringTokenPosition = *position

	node.Parts = append(node.Parts, &ast.NamePart{
		Position: position,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        chars,
			Position:     stringTokenPosition,
			FreeFloating: nameToken.FreeFloating,
		},
		Value: chars,
	})

	return node
}

func (parser *Parser) parseNameFullyQualifiedToken(nameToken *token.Token) *ast.NameFullyQualified {
	node := &ast.NameFullyQualified{
		Position: parser.builder.NewTokenPosition(nameToken),
	}
	startPos := nameToken.Position.StartPos
	chars := nameToken.Value

	// ns separator token

	position := parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + 1

	// the separator opens the name, so it carries the trivia that preceded the
	// whole token — leave it on a later part and it prints after the separator,
	// turning `\Foo` into `\ Foo`
	node.NsSeparatorTkn = &token.Token{
		ID:           token.T_NS_SEPARATOR,
		Value:        chars[:1],
		Position:     position,
		FreeFloating: nameToken.FreeFloating,
	}
	nameToken.FreeFloating = nil

	startPos = startPos + 1
	chars = chars[1:]

	// parts

	for {
		i := bytes.Index(chars, []byte("\\"))
		if i < 0 {
			break
		}

		position = parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + i

		stringTokenPosition := parser.Lexer.positionPool.Get()
		*stringTokenPosition = *position

		node.Parts = append(node.Parts, &ast.NamePart{
			Position: position,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        chars[:i],
				Position:     stringTokenPosition,
				FreeFloating: nameToken.FreeFloating,
			},
			Value: chars[:i],
		})
		nameToken.FreeFloating = nil
		startPos = startPos + i
		chars = chars[i:]

		position = parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + 1

		node.SeparatorTkns = append(node.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    chars[:1],
			Position: position,
		})
		startPos = startPos + 1
		chars = chars[1:]
	}

	// last part

	position = parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + len(chars)
	stringTokenPosition := parser.Lexer.positionPool.Get()
	*stringTokenPosition = *position

	node.Parts = append(node.Parts, &ast.NamePart{
		Position: position,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        chars,
			Position:     stringTokenPosition,
			FreeFloating: nameToken.FreeFloating,
		},
		Value: chars,
	})

	return node
}

func (parser *Parser) parseNameToken(nameToken *token.Token) *ast.Name {
	node := &ast.Name{
		Position: parser.builder.NewTokenPosition(nameToken),
	}
	startPos := nameToken.Position.StartPos
	chars := nameToken.Value

	for {
		i := bytes.Index(chars, []byte("\\"))
		if i < 0 {
			break
		}

		position := parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + i

		stringTokenPosition := parser.Lexer.positionPool.Get()
		*stringTokenPosition = *position

		node.Parts = append(node.Parts, &ast.NamePart{
			Position: position,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        chars[:i],
				Position:     stringTokenPosition,
				FreeFloating: nameToken.FreeFloating,
			},
			Value: chars[:i],
		})
		nameToken.FreeFloating = nil
		startPos = startPos + i
		chars = chars[i:]

		position = parser.Lexer.positionPool.Get()
		position.StartLine = nameToken.Position.StartLine
		position.EndLine = nameToken.Position.EndLine
		position.StartPos = startPos
		position.EndPos = startPos + 1

		node.SeparatorTkns = append(node.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    chars[:1],
			Position: position,
		})
		startPos = startPos + 1
		chars = chars[1:]
	}

	position := parser.Lexer.positionPool.Get()
	position.StartLine = nameToken.Position.StartLine
	position.EndLine = nameToken.Position.EndLine
	position.StartPos = startPos
	position.EndPos = startPos + len(chars)
	stringTokenPosition := parser.Lexer.positionPool.Get()
	*stringTokenPosition = *position

	node.Parts = append(node.Parts, &ast.NamePart{
		Position: position,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        chars,
			Position:     stringTokenPosition,
			FreeFloating: nameToken.FreeFloating,
		},
		Value: chars,
	})

	return node
}
