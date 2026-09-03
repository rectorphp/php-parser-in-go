package php8

import (
	"bytes"
	"strings"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

type Lexer struct {
	data           []byte
	phpVersion     *version.Version
	errHandlerFunc func(*errors.Error)

	p, pe, cs   int
	ts, te, act int
	stack       []int
	top         int

	heredocLabel []byte
	tokenPool    *token.Pool
	positionPool *position.Pool
	newLines     NewLines
}

func NewLexer(data []byte, config conf.Config) *Lexer {
	lex := &Lexer{
		data:           data,
		phpVersion:     config.Version,
		errHandlerFunc: config.ErrorHandlerFunc,

		pe:    len(data),
		stack: make([]int, 0),

		tokenPool:    token.NewPool(position.DefaultBlockSize),
		positionPool: position.NewPool(token.DefaultBlockSize),
		newLines:     NewLines{make([]int, 0, 128)},
	}

	initLexer(lex)

	return lex
}

func (lex *Lexer) setTokenPosition(token *token.Token) {
	pos := lex.positionPool.Get()

	pos.StartLine = lex.newLines.GetLine(lex.ts)
	pos.EndLine = lex.newLines.GetLine(lex.te - 1)
	pos.StartPos = lex.ts
	pos.EndPos = lex.te

	token.Position = pos
}

func (lex *Lexer) addFreeFloatingToken(parentToken *token.Token, tokenID token.ID, start, end int) {
	skippedTkn := lex.tokenPool.Get()
	skippedTkn.ID = tokenID
	skippedTkn.Value = lex.data[start:end]

	lex.setTokenPosition(skippedTkn)

	if parentToken.FreeFloating == nil {
		parentToken.FreeFloating = make([]*token.Token, 0, 2)
	}

	parentToken.FreeFloating = append(parentToken.FreeFloating, skippedTkn)
}

func (lex *Lexer) isNotStringVar() bool {
	position := lex.p
	if lex.data[position-1] == '\\' && lex.data[position-2] != '\\' {
		return true
	}

	if len(lex.data) < position+1 {
		return true
	}

	if lex.data[position] == '$' && (lex.data[position+1] == '{' || isValidVarNameStart(lex.data[position+1])) {
		return false
	}

	if lex.data[position] == '{' && lex.data[position+1] == '$' {
		return false
	}

	return true
}

func (lex *Lexer) isNotStringEnd(endChar byte) bool {
	position := lex.p
	if lex.data[position-1] == '\\' && lex.data[position-2] != '\\' {
		return true
	}

	return !(lex.data[position] == endChar)
}

// heredocFlexibleSince is PHP 7.3, when flexible heredoc/nowdoc syntax landed.
var heredocFlexibleSince = &version.Version{Major: 7, Minor: 3}

func (lex *Lexer) isHeredocEnd(position int) bool {
	if lex.phpVersion.GreaterOrEqual(heredocFlexibleSince) {
		return lex.isHeredocEndSince73(position)
	}

	return lex.isHeredocEndBefore73(position)
}

func (lex *Lexer) isHeredocEndBefore73(position int) bool {
	if lex.data[position-1] != '\r' && lex.data[position-1] != '\n' {
		return false
	}

	labelLen := len(lex.heredocLabel)
	if len(lex.data) < position+labelLen {
		return false
	}

	if len(lex.data) > position+labelLen && lex.data[position+labelLen] != ';' && lex.data[position+labelLen] != '\r' && lex.data[position+labelLen] != '\n' {
		return false
	}

	if len(lex.data) > position+labelLen+1 && lex.data[position+labelLen] == ';' && lex.data[position+labelLen+1] != '\r' && lex.data[position+labelLen+1] != '\n' {
		return false
	}

	return bytes.Equal(lex.heredocLabel, lex.data[position:position+labelLen])
}

func (lex *Lexer) isHeredocEndSince73(position int) bool {
	if lex.data[position-1] != '\r' && lex.data[position-1] != '\n' {
		return false
	}

	if position == len(lex.data) {
		return false
	}

	for lex.data[position] == ' ' || lex.data[position] == '\t' {
		position++
	}

	labelLen := len(lex.heredocLabel)
	if len(lex.data) < position+labelLen {
		return false
	}

	if len(lex.data) > position+labelLen && isValidVarName(lex.data[position+labelLen]) {
		return false
	}

	label := string(lex.heredocLabel)
	dataLabel := string(lex.data[position : position+labelLen])

	_, _ = label, dataLabel

	if bytes.Equal(lex.heredocLabel, lex.data[position:position+labelLen]) {
		lex.p = position
		return true
	}

	return false
}

func (lex *Lexer) isNotHeredocEnd(position int) bool {
	return !lex.isHeredocEnd(position)
}

func (lex *Lexer) growCallStack() {
	if lex.top == len(lex.stack) {
		lex.stack = append(lex.stack, 0)
	}
}

func (lex *Lexer) isNotPhpCloseToken() bool {
	if lex.p+1 == len(lex.data) {
		return true
	}

	return lex.data[lex.p] != '?' || lex.data[lex.p+1] != '>'
}

func (lex *Lexer) isNotNewLine() bool {
	if lex.data[lex.p] == '\n' && lex.data[lex.p-1] == '\r' {
		return true
	}

	return lex.data[lex.p-1] != '\n' && lex.data[lex.p-1] != '\r'
}

func (lex *Lexer) call(state int, fnext int) {
	lex.growCallStack()

	lex.stack[lex.top] = state
	lex.top++

	lex.p++
	lex.cs = fnext
}

func (lex *Lexer) ret(count int) {
	lex.top = max(lex.top-count, 0)
	lex.cs = lex.stack[lex.top]
	lex.p++
}

func (lex *Lexer) ungetStr(suffix string) {
	tokenStr := string(lex.data[lex.ts:lex.te])
	if strings.HasSuffix(tokenStr, suffix) {
		lex.ungetCnt(len(suffix))
	}
}

func (lex *Lexer) ungetCnt(count int) {
	lex.p = lex.p - count
	lex.te = lex.te - count
}

func (lex *Lexer) error(message string) {
	if lex.errHandlerFunc == nil {
		return
	}

	pos := position.NewPosition(
		lex.newLines.GetLine(lex.ts),
		lex.newLines.GetLine(lex.te-1),
		lex.ts,
		lex.te,
	)

	lex.errHandlerFunc(errors.NewError(message, pos))
}

func isValidVarNameStart(nameByte byte) bool {
	return (nameByte >= 'A' && nameByte <= 'Z') || (nameByte >= 'a' && nameByte <= 'z') || nameByte == '_' || nameByte >= 0x80
}

func isValidVarName(nameByte byte) bool {
	return (nameByte >= 'A' && nameByte <= 'Z') || (nameByte >= 'a' && nameByte <= 'z') || (nameByte >= '0' && nameByte <= '9') || nameByte == '_' || nameByte >= 0x80
}
