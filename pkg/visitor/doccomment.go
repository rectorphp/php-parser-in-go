package visitor

import (
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

// GetDocComment returns the /** */ doc comment immediately preceding a
// declaration node, or nil when there is none. It is supported for the nodes
// that carry a doc block: class, interface, trait and enum declarations,
// functions and methods, and property and class-constant lists. For any other
// node it returns nil.
//
// The comment is read from the free-floating tokens that precede the node's
// first significant token (an attribute group, a modifier, or the declaration
// keyword). When several doc comments precede the node, the closest one is
// returned.
func GetDocComment(node ast.Vertex) *token.Token {
	return lastDocComment(leadingTokens(node))
}

// GetDocCommentText is GetDocComment returning the comment text, or "" when
// there is no doc comment.
func GetDocCommentText(node ast.Vertex) string {
	if doc := GetDocComment(node); doc != nil {
		return string(doc.Value)
	}
	return ""
}

func leadingTokens(node ast.Vertex) []*token.Token {
	switch n := node.(type) {
	case *ast.StmtClass:
		return collectLeading(n.AttrGroups, n.Modifiers, n.ClassTkn)
	case *ast.StmtInterface:
		return collectLeading(n.AttrGroups, nil, n.InterfaceTkn)
	case *ast.StmtTrait:
		return collectLeading(n.AttrGroups, nil, n.TraitTkn)
	case *ast.StmtEnum:
		return collectLeading(n.AttrGroups, nil, n.EnumTkn)
	case *ast.StmtFunction:
		return collectLeading(n.AttrGroups, nil, n.FunctionTkn)
	case *ast.StmtClassMethod:
		return collectLeading(n.AttrGroups, n.Modifiers, n.FunctionTkn)
	case *ast.StmtPropertyList:
		return collectLeading(n.AttrGroups, n.Modifiers, nil)
	case *ast.StmtClassConstList:
		return collectLeading(n.AttrGroups, n.Modifiers, n.ConstTkn)
	default:
		return nil
	}
}

func collectLeading(attrGroups, modifiers []ast.Vertex, keyword *token.Token) []*token.Token {
	var tokens []*token.Token
	for _, group := range attrGroups {
		if attributeGroup, ok := group.(*ast.AttributeGroup); ok {
			tokens = append(tokens, attributeGroup.OpenAttributeTkn)
		}
	}
	for _, modifier := range modifiers {
		if identifier, ok := modifier.(*ast.Identifier); ok {
			tokens = append(tokens, identifier.IdentifierTkn)
		}
	}
	if keyword != nil {
		tokens = append(tokens, keyword)
	}
	return tokens
}

func lastDocComment(tokens []*token.Token) *token.Token {
	var found *token.Token
	for _, tkn := range tokens {
		if tkn == nil {
			continue
		}
		for _, free := range tkn.FreeFloating {
			if free.ID == token.T_DOC_COMMENT {
				found = free
			}
		}
	}
	return found
}
