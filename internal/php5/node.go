package php5

import (
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

type ParserBrackets struct {
	Position        *position.Position
	OpenBracketTkn  *token.Token
	Child           ast.Vertex
	CloseBracketTkn *token.Token
}

func (parserBrackets *ParserBrackets) Accept(visitor ast.Visitor) {
	// do nothing
}

func (parserBrackets *ParserBrackets) GetPosition() *position.Position {
	return parserBrackets.Position
}

type ParserSeparatedList struct {
	Position      *position.Position
	Items         []ast.Vertex
	SeparatorTkns []*token.Token
}

func (parserSeparatedList *ParserSeparatedList) Accept(visitor ast.Visitor) {
	// do nothing
}

func (parserSeparatedList *ParserSeparatedList) GetPosition() *position.Position {
	return parserSeparatedList.Position
}

// TraitAdaptationList node
type TraitAdaptationList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []ast.Vertex
	CloseCurlyBracketTkn *token.Token
}

func (traitAdaptationList *TraitAdaptationList) Accept(visitor ast.Visitor) {
	// do nothing
}

func (traitAdaptationList *TraitAdaptationList) GetPosition() *position.Position {
	return traitAdaptationList.Position
}

// ArgumentList node
type ArgumentList struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Arguments           []ast.Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (argumentList *ArgumentList) Accept(visitor ast.Visitor) {
	// do nothing
}

func (argumentList *ArgumentList) GetPosition() *position.Position {
	return argumentList.Position
}

// TraitMethodRef node
type TraitMethodRef struct {
	Position       *position.Position
	Trait          ast.Vertex
	DoubleColonTkn *token.Token
	Method         ast.Vertex
}

func (traitMethodRef *TraitMethodRef) Accept(visitor ast.Visitor) {
	// do nothing
}

func (traitMethodRef *TraitMethodRef) GetPosition() *position.Position {
	return traitMethodRef.Position
}
