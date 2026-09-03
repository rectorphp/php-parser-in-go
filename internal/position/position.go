package position

import (
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

type Builder struct {
	pool *position.Pool
}

func NewBuilder() *Builder {
	return &Builder{
		pool: position.NewPool(position.DefaultBlockSize),
	}
}

func getListStartPos(nodes []ast.Vertex) startPos {
	if nodes == nil {
		return startPos{-1, -1}
	}

	if len(nodes) == 0 {
		return startPos{-1, -1}
	}

	return getNodeStartPos(nodes[0])
}

func getNodeStartPos(node ast.Vertex) startPos {
	startLine := -1
	startPosition := -1

	if node == nil {
		return startPos{-1, -1}
	}

	position := node.GetPosition()
	if position != nil {
		startLine = position.StartLine
		startPosition = position.StartPos
	}

	return startPos{startLine, startPosition}
}

func getListEndPos(nodes []ast.Vertex) endPos {
	if nodes == nil {
		return endPos{-1, -1}
	}

	if len(nodes) == 0 {
		return endPos{-1, -1}
	}

	return getNodeEndPos(nodes[len(nodes)-1])
}

func getNodeEndPos(node ast.Vertex) endPos {
	endLine := -1
	endPosition := -1

	if node == nil {
		return endPos{-1, -1}
	}

	position := node.GetPosition()
	if position != nil {
		endLine = position.EndLine
		endPosition = position.EndPos
	}

	return endPos{endLine, endPosition}
}

// NewNodeListPosition returns new Position
func (builder *Builder) NewNodeListPosition(list []ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodePosition returns new Position
func (builder *Builder) NewNodePosition(node ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getNodeStartPos(node).startLine
	pos.EndLine = getNodeEndPos(node).endLine
	pos.StartPos = getNodeStartPos(node).startPos
	pos.EndPos = getNodeEndPos(node).endPos

	return pos
}

// NewTokenPosition returns new Position
func (builder *Builder) NewTokenPosition(token *token.Token) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = token.Position.StartLine
	pos.EndLine = token.Position.EndLine
	pos.StartPos = token.Position.StartPos
	pos.EndPos = token.Position.EndPos

	return pos
}

// NewTokensPosition returns new Position
func (builder *Builder) NewTokensPosition(startToken *token.Token, endToken *token.Token) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = startToken.Position.StartLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = startToken.Position.StartPos
	pos.EndPos = endToken.Position.EndPos

	return pos
}

// NewTokenNodePosition returns new Position
func (builder *Builder) NewTokenNodePosition(token *token.Token, node ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = token.Position.StartLine
	pos.EndLine = getNodeEndPos(node).endLine
	pos.StartPos = token.Position.StartPos
	pos.EndPos = getNodeEndPos(node).endPos

	return pos
}

// NewNodeTokenPosition returns new Position
func (builder *Builder) NewNodeTokenPosition(node ast.Vertex, token *token.Token) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getNodeStartPos(node).startLine
	pos.EndLine = token.Position.EndLine
	pos.StartPos = getNodeStartPos(node).startPos
	pos.EndPos = token.Position.EndPos

	return pos
}

// NewNodesPosition returns new Position
func (builder *Builder) NewNodesPosition(startNode ast.Vertex, endNode ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getNodeStartPos(startNode).startLine
	pos.EndLine = getNodeEndPos(endNode).endLine
	pos.StartPos = getNodeStartPos(startNode).startPos
	pos.EndPos = getNodeEndPos(endNode).endPos

	return pos
}

// NewNodeListTokenPosition returns new Position
func (builder *Builder) NewNodeListTokenPosition(list []ast.Vertex, token *token.Token) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = token.Position.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = token.Position.EndPos

	return pos
}

// NewTokenNodeListPosition returns new Position
func (builder *Builder) NewTokenNodeListPosition(token *token.Token, list []ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = token.Position.StartLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = token.Position.StartPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodeNodeListPosition returns new Position
func (builder *Builder) NewNodeNodeListPosition(node ast.Vertex, list []ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getNodeStartPos(node).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getNodeStartPos(node).startPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodeListNodePosition returns new Position
func (builder *Builder) NewNodeListNodePosition(list []ast.Vertex, node ast.Vertex) *position.Position {
	pos := builder.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getNodeEndPos(node).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getNodeEndPos(node).endPos

	return pos
}

// NewOptionalListTokensPosition returns new Position
func (builder *Builder) NewOptionalListTokensPosition(list []ast.Vertex, token *token.Token, endToken *token.Token) *position.Position {
	pos := builder.pool.Get()

	if list == nil {
		pos.StartLine = token.Position.StartLine
		pos.EndLine = endToken.Position.EndLine
		pos.StartPos = token.Position.StartPos
		pos.EndPos = endToken.Position.EndPos

		return pos
	}
	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = endToken.Position.EndPos

	return pos
}
