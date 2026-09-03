package dumper_test

import (
	"bytes"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/dumper"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

func TestDumper_root(test *testing.T) {
	buffer := bytes.NewBufferString("")

	dumperValue := dumper.NewDumper(buffer).WithTokens().WithPositions()
	node := &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   2,
			StartPos:  3,
			EndPos:    4,
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		EndTkn: &token.Token{
			FreeFloating: []*token.Token{
				{
					ID:    token.T_WHITESPACE,
					Value: []byte(" "),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   2,
						StartPos:  3,
						EndPos:    4,
					},
				},
			},
		},
	}
	node.Accept(dumperValue)

	expected := `&ast.Root{
	Position: &position.Position{
		StartLine: 1,
		EndLine:   2,
		StartPos:  3,
		EndPos:    4,
	},
	Stmts: []ast.Vertex{
		&ast.StmtNop{
		},
	},
	EndTkn: &token.Token{
		FreeFloating: []*token.Token{
			{
				ID: token.T_WHITESPACE,
				Value: []byte(" "),
				Position: &position.Position{
					StartLine: 1,
					EndLine:   2,
					StartPos:  3,
					EndPos:    4,
				},
			},
		},
	},
},
`
	actual := buffer.String()

	if expected != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
