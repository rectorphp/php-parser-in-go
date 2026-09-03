package position_test

import (
	"gotest.tools/assert"
	"testing"

	builder "github.com/rectorphp/php-parser-in-go/internal/position"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

func TestNewTokenPosition(test *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewTokenPosition(tkn)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokensPosition(test *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewBuilder().NewTokensPosition(token1, token2)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewNodePosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodePosition(node)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokenNodePosition(test *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    12,
		},
	}

	pos := builder.NewBuilder().NewTokenNodePosition(tkn, node)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeTokenPosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    12,
		},
	}

	pos := builder.NewBuilder().NewNodeTokenPosition(node, tkn)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeListPosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewBuilder().NewNodeListPosition([]ast.Vertex{node, nodeValue})

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodesPosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewBuilder().NewNodesPosition(node, nodeValue)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodeListTokenPosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  20,
			EndPos:    22,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPosition([]ast.Vertex{node, nodeValue}, tkn)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 3, EndPos: 22}, pos)
}

func TestNewTokenNodeListPosition(test *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    2,
		},
	}

	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  3,
			EndPos:    10,
		},
	}

	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  11,
			EndPos:    20,
		},
	}

	pos := builder.NewBuilder().NewTokenNodeListPosition(tkn, []ast.Vertex{node, nodeValue})

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 3, EndPos: 20}, pos)
}

func TestNewNodeNodeListPosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}

	node2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPosition(node, []ast.Vertex{nodeValue, node2})

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewNodeListNodePosition(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}
	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	node2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewBuilder().NewNodeListNodePosition([]ast.Vertex{node, nodeValue}, node2)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewOptionalListTokensPosition(test *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewBuilder().NewOptionalListTokensPosition(nil, token1, token2)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewOptionalListTokensPosition2(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	nodeValue := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 4,
			EndLine:   4,
			StartPos:  27,
			EndPos:    29,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 5,
			EndLine:   5,
			StartPos:  30,
			EndPos:    32,
		},
	}

	pos := builder.NewBuilder().NewOptionalListTokensPosition([]ast.Vertex{node, nodeValue}, token1, token2)

	assert.DeepEqual(test, &position.Position{StartLine: 2, EndLine: 5, StartPos: 9, EndPos: 32}, pos)
}

func TestNilNodePos(test *testing.T) {
	pos := builder.NewBuilder().NewNodesPosition(nil, nil)

	assert.DeepEqual(test, &position.Position{StartLine: -1, EndLine: -1, StartPos: -1, EndPos: -1}, pos)
}

func TestNilNodeListPos(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPosition(node, nil)

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestNilNodeListTokenPos(test *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPosition(nil, tkn)

	assert.DeepEqual(test, &position.Position{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}

func TestEmptyNodeListPos(test *testing.T) {
	node := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPosition(node, []ast.Vertex{})

	assert.DeepEqual(test, &position.Position{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestEmptyNodeListTokenPos(test *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPosition([]ast.Vertex{}, tkn)

	assert.DeepEqual(test, &position.Position{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}
