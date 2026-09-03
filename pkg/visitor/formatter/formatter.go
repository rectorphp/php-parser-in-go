package formatter

import (
	"bytes"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

type formatterState int

const (
	FormatterStateHTML formatterState = iota
	FormatterStatePHP
)

type formatter struct {
	state        formatterState
	indent       int
	freeFloating []*token.Token

	lastSemiColon *token.Token
}

func NewFormatter() *formatter {
	return &formatter{}
}

func (formatter *formatter) WithState(state formatterState) *formatter {
	formatter.state = state
	return formatter
}

func (formatter *formatter) WithIndent(indent int) *formatter {
	formatter.indent = indent
	return formatter
}

func (formatter *formatter) addFreeFloating(tokenID token.ID, value []byte) {
	formatter.freeFloating = append(formatter.freeFloating, &token.Token{
		ID:    tokenID,
		Value: value,
	})
}

func (formatter *formatter) addIndent() {
	if formatter.indent < 1 {
		return
	}

	formatter.freeFloating = append(formatter.freeFloating, &token.Token{
		ID:    token.T_WHITESPACE,
		Value: bytes.Repeat([]byte("    "), formatter.indent),
	})
}

func (formatter *formatter) resetFreeFloating() {
	formatter.freeFloating = nil
}

func (formatter *formatter) getFreeFloating() []*token.Token {
	defer formatter.resetFreeFloating()

	if formatter.state == FormatterStateHTML {
		openTagToken := &token.Token{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php "),
		}
		formatter.freeFloating = append([]*token.Token{openTagToken}, formatter.freeFloating...)

		formatter.state = FormatterStatePHP
	}

	return formatter.freeFloating
}

func (formatter *formatter) newToken(tokenID token.ID, value []byte) *token.Token {
	return &token.Token{
		ID:           tokenID,
		Value:        value,
		FreeFloating: formatter.getFreeFloating(),
	}
}

func (formatter *formatter) formatList(nodes []ast.Vertex, separator byte) []*token.Token {
	separatorTokens := make([]*token.Token, len(nodes)-1)
	for i, node := range nodes {
		node.Accept(formatter)

		if i != len(nodes)-1 {
			separatorTokens[i] = formatter.newToken(token.ID(separator), []byte{separator})
			formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		}
	}

	return separatorTokens
}

func (formatter *formatter) formatStmts(list *[]ast.Vertex) {
	var insertCounter int

	for i, statement := range *list {
		formatter.lastSemiColon = nil

		if _, ok := statement.(*ast.StmtInlineHtml); ok {
			if formatter.lastSemiColon != nil {
				formatter.lastSemiColon.Value = append(formatter.lastSemiColon.Value, '?', '>')
			} else {
				*list = insert(*list, i+insertCounter, &ast.StmtNop{
					SemiColonTkn: &token.Token{
						Value: []byte("?>"),
					},
				})
				insertCounter++
			}
		} else {
			formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			formatter.addIndent()
		}

		statement.Accept(formatter)
	}
}

func (formatter *formatter) newSemicolonTkn() *token.Token {
	formatter.lastSemiColon = formatter.newToken(';', []byte(";"))
	return formatter.lastSemiColon
}

func insert(slice []ast.Vertex, index int, values ...ast.Vertex) []ast.Vertex {
	if newLength := len(slice) + len(values); newLength <= cap(slice) {
		result := slice[:newLength]
		copy(result[index+len(values):], slice[index:])
		copy(result[index:], values)
		return result
	}
	result := make([]ast.Vertex, len(slice)+len(values))
	copy(result, slice[:index])
	copy(result[index:], values)
	copy(result[index+len(values):], slice[index:])
	return result
}

func (formatter *formatter) Root(node *ast.Root) {
	formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
	formatter.addIndent()

	formatter.formatStmts(&node.Stmts)
}

func (formatter *formatter) Nullable(node *ast.Nullable) {
	node.QuestionTkn = formatter.newToken('?', []byte("?"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) Union(node *ast.Union) {
	node.SeparatorTkns = make([]*token.Token, len(node.Types)-1)
	for i, typeNode := range node.Types {
		typeNode.Accept(formatter)

		if i != len(node.Types)-1 {
			node.SeparatorTkns[i] = formatter.newToken('|', []byte("|"))
		}
	}
}

func (formatter *formatter) Intersection(node *ast.Intersection) {
	node.SeparatorTkns = make([]*token.Token, len(node.Types)-1)
	for i, typeNode := range node.Types {
		typeNode.Accept(formatter)

		if i != len(node.Types)-1 {
			node.SeparatorTkns[i] = formatter.newToken('&', []byte("&"))
		}
	}
}

func (formatter *formatter) Parameter(node *ast.Parameter) {
	for _, modifier := range node.Modifiers {
		modifier.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	if node.VariadicTkn != nil {
		node.VariadicTkn = formatter.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	node.Var.Accept(formatter)

	if node.DefaultValue != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.EqualTkn = formatter.newToken('=', []byte("="))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.DefaultValue.Accept(formatter)
	}
}

func (formatter *formatter) Identifier(node *ast.Identifier) {
	if node.IdentifierTkn == nil {
		node.IdentifierTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.IdentifierTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) Argument(node *ast.Argument) {
	if node.Name != nil {
		node.Name.Accept(formatter)
		node.ColonTkn = formatter.newToken(':', []byte(":"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if node.VariadicTkn != nil {
		node.VariadicTkn = formatter.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.Expr.Accept(formatter)
}

func (formatter *formatter) Attribute(node *ast.Attribute) {
	node.Name.Accept(formatter)
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Args) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) AttributeGroup(node *ast.AttributeGroup) {
	node.OpenAttributeTkn = formatter.newToken(token.T_ATTRIBUTE, []byte("#["))
	node.SeparatorTkns = nil
	if len(node.Attrs) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Attrs, ',')
	}
	node.CloseAttributeTkn = formatter.newToken(']', []byte("]"))
}

func (formatter *formatter) StmtBreak(node *ast.StmtBreak) {
	node.BreakTkn = formatter.newToken(token.T_BREAK, []byte("break"))

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Expr.Accept(formatter)
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtCase(node *ast.StmtCase) {
	node.CaseTkn = formatter.newToken(token.T_CASE, []byte("case"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Cond.Accept(formatter)

	node.CaseSeparatorTkn = formatter.newToken(':', []byte(":"))

	formatter.indent++
	formatter.formatStmts(&node.Stmts)
	formatter.indent--
}

func (formatter *formatter) StmtCatch(node *ast.StmtCatch) {
	node.CatchTkn = formatter.newToken(token.T_CATCH, []byte("catch"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.SeparatorTkns = make([]*token.Token, len(node.Types)-1)
	for i, typeNode := range node.Types {
		typeNode.Accept(formatter)

		if i != len(node.Types)-1 {
			formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
			node.SeparatorTkns[i] = formatter.newToken('|', []byte("|"))
			formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		}
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Var.Accept(formatter)

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtClass(node *ast.StmtClass) {
	for _, modifier := range node.Modifiers {
		modifier.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.ClassTkn = formatter.newToken(token.T_CLASS, []byte("class"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Name.Accept(formatter)

	node.OpenParenthesisTkn = nil
	node.CloseParenthesisTkn = nil
	if len(node.Args) > 0 {
		node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

		node.SeparatorTkns = formatter.formatList(node.Args, ',')

		node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	if node.Extends != nil {
		node.ExtendsTkn = formatter.newToken(token.T_EXTENDS, []byte("extends"))
		node.Extends.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if node.Implements != nil {
		node.ImplementsTkn = formatter.newToken(token.T_IMPLEMENTS, []byte("implements"))
		node.ImplementsSeparatorTkns = formatter.formatList(node.Implements, ',')
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtClassConstList(node *ast.StmtClassConstList) {
	for _, modifier := range node.Modifiers {
		modifier.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.ConstTkn = formatter.newToken(token.T_CONST, []byte("const"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.SeparatorTkns = formatter.formatList(node.Consts, ',')

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtClassMethod(node *ast.StmtClassMethod) {
	for _, modifier := range node.Modifiers {
		modifier.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.FunctionTkn = formatter.newToken(token.T_FUNCTION, []byte("function"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.Name.Accept(formatter)

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.SeparatorTkns = nil
	if len(node.Params) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Params, ',')
	}

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	if node.ReturnType != nil {
		node.ColonTkn = formatter.newToken(':', []byte(":"))

		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.ReturnType.Accept(formatter)
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtConstList(node *ast.StmtConstList) {
	node.ConstTkn = formatter.newToken(token.T_CONST, []byte("const"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = formatter.formatList(node.Consts, ',')

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtConstant(node *ast.StmtConstant) {
	node.Name.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken('=', []byte("="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) StmtContinue(node *ast.StmtContinue) {
	node.ContinueTkn = formatter.newToken(token.T_CONTINUE, []byte("continue"))

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Expr.Accept(formatter)
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtDeclare(node *ast.StmtDeclare) {
	node.ColonTkn = nil
	node.EndDeclareTkn = nil
	node.SemiColonTkn = nil

	node.DeclareTkn = formatter.newToken(token.T_DECLARE, []byte("declare"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.SeparatorTkns = nil
	if len(node.Consts) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Consts, ',')
	}

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtDefault(node *ast.StmtDefault) {
	node.DefaultTkn = formatter.newToken(token.T_DEFAULT, []byte("default"))

	node.CaseSeparatorTkn = formatter.newToken(':', []byte(":"))

	formatter.indent++
	formatter.formatStmts(&node.Stmts)
	formatter.indent--
}

func (formatter *formatter) StmtDo(node *ast.StmtDo) {
	node.DoTkn = formatter.newToken(token.T_DO, []byte("do"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.WhileTkn = formatter.newToken(token.T_WHILE, []byte("while"))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Cond.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtEcho(node *ast.StmtEcho) {
	node.EchoTkn = formatter.newToken(token.T_ECHO, []byte("echo"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = nil
	if len(node.Exprs) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Exprs, ',')
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtElse(node *ast.StmtElse) {
	node.ColonTkn = nil

	node.ElseTkn = formatter.newToken(token.T_ELSE, []byte("else"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtElseIf(node *ast.StmtElseIf) {
	node.ColonTkn = nil

	node.ElseIfTkn = formatter.newToken(token.T_ELSEIF, []byte("elseif"))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Cond.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtExpression(node *ast.StmtExpression) {
	node.Expr.Accept(formatter)
	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtFinally(node *ast.StmtFinally) {
	node.FinallyTkn = formatter.newToken(token.T_FINALLY, []byte("finally"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtFor(node *ast.StmtFor) {
	node.ColonTkn = nil
	node.EndForTkn = nil
	node.SemiColonTkn = nil

	node.ForTkn = formatter.newToken(token.T_FOR, []byte("for"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.InitSeparatorTkns = nil
	if len(node.Init) > 0 {
		node.InitSeparatorTkns = formatter.formatList(node.Init, ',')
	}

	node.InitSemiColonTkn = formatter.newSemicolonTkn()
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.CondSeparatorTkns = nil
	if len(node.Cond) > 0 {
		node.CondSeparatorTkns = formatter.formatList(node.Cond, ',')
	}

	node.CondSemiColonTkn = formatter.newSemicolonTkn()
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.LoopSeparatorTkns = nil
	if len(node.Loop) > 0 {
		node.LoopSeparatorTkns = formatter.formatList(node.Loop, ',')
	}

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtForeach(node *ast.StmtForeach) {
	node.ColonTkn = nil
	node.EndForeachTkn = nil
	node.SemiColonTkn = nil

	node.ForeachTkn = formatter.newToken(token.T_FOREACH, []byte("foreach"))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.Expr.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.AsTkn = formatter.newToken(token.T_AS, []byte("as"))

	if node.Key != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Key.Accept(formatter)

		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.DoubleArrowTkn = formatter.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}
	node.Var.Accept(formatter)

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) StmtFunction(node *ast.StmtFunction) {
	node.FunctionTkn = formatter.newToken(token.T_FUNCTION, []byte("function"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.Name.Accept(formatter)

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))

	node.SeparatorTkns = nil
	if len(node.Params) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Params, ',')
	}

	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	if node.ReturnType != nil {
		node.ColonTkn = formatter.newToken(':', []byte(":"))

		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.ReturnType.Accept(formatter)
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtGlobal(node *ast.StmtGlobal) {
	node.GlobalTkn = formatter.newToken(token.T_GLOBAL, []byte("global"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = nil
	if len(node.Vars) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Vars, ',')
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtGoto(node *ast.StmtGoto) {
	node.GotoTkn = formatter.newToken(token.T_GOTO, []byte("goto"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Label.Accept(formatter)

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtHaltCompiler(node *ast.StmtHaltCompiler) {
	node.HaltCompilerTkn = formatter.newToken(token.T_HALT_COMPILER, []byte("__halt_compiler"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtIf(node *ast.StmtIf) {
	node.ColonTkn = nil
	node.EndIfTkn = nil
	node.SemiColonTkn = nil

	node.IfTkn = formatter.newToken(token.T_IF, []byte("if"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Cond.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)

	if len(node.ElseIf) > 0 {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		formatter.formatList(node.ElseIf, ' ')
	}

	if node.Else != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Else.Accept(formatter)
	}
}

func (formatter *formatter) StmtInlineHtml(node *ast.StmtInlineHtml) {
	node.InlineHtmlTkn = formatter.newToken(token.T_STRING, node.Value)
	formatter.state = FormatterStateHTML
}

func (formatter *formatter) StmtEnum(node *ast.StmtEnum) {
	node.EnumTkn = formatter.newToken(token.T_ENUM, []byte("enum"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Name.Accept(formatter)

	if node.Type != nil {
		node.ColonTkn = formatter.newToken(':', []byte(":"))
		node.Type.Accept(formatter)
	}
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Implements != nil {
		node.ImplementsTkn = formatter.newToken(token.T_IMPLEMENTS, []byte("implements"))
		node.ImplementsSeparatorTkns = formatter.formatList(node.Implements, ',')
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtEnumCase(node *ast.StmtEnumCase) {
	node.CaseTkn = formatter.newToken(token.T_CASE, []byte("case"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Name.Accept(formatter)

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.EqualTkn = formatter.newToken('=', []byte("="))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Expr.Accept(formatter)
	}

	node.SemiColonTkn = formatter.newToken(';', []byte(";"))
}

func (formatter *formatter) StmtInterface(node *ast.StmtInterface) {
	node.InterfaceTkn = formatter.newToken(token.T_INTERFACE, []byte("interface"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Name.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Extends != nil {
		node.ExtendsTkn = formatter.newToken(token.T_EXTENDS, []byte("extends"))
		node.ExtendsSeparatorTkns = formatter.formatList(node.Extends, ',')
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtLabel(node *ast.StmtLabel) {
	node.Name.Accept(formatter)
	node.ColonTkn = formatter.newToken(':', []byte(":"))
}

func (formatter *formatter) StmtNamespace(node *ast.StmtNamespace) {
	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	node.SemiColonTkn = nil

	node.NsTkn = formatter.newToken(token.T_NAMESPACE, []byte("namespace"))

	if node.Name != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Name.Accept(formatter)
	}

	if len(node.Stmts) > 0 {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		if len(node.Stmts) > 0 {
			formatter.indent++
			formatter.formatStmts(&node.Stmts)
			formatter.indent--

			formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			formatter.addIndent()
		}
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	} else {
		node.SemiColonTkn = formatter.newSemicolonTkn()
	}

}

func (formatter *formatter) StmtNop(node *ast.StmtNop) {
	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtProperty(node *ast.StmtProperty) {
	node.Var.Accept(formatter)

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.EqualTkn = formatter.newToken('=', []byte("="))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

		node.Expr.Accept(formatter)
	}
}

func (formatter *formatter) StmtPropertyList(node *ast.StmtPropertyList) {
	for _, modifier := range node.Modifiers {
		modifier.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.SeparatorTkns = formatter.formatList(node.Props, ',')

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtReturn(node *ast.StmtReturn) {
	node.ReturnTkn = formatter.newToken(token.T_RETURN, []byte("return"))

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Expr.Accept(formatter)
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtStatic(node *ast.StmtStatic) {
	node.StaticTkn = formatter.newToken(token.T_STATIC, []byte("static"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = nil
	if len(node.Vars) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Vars, ',')
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtStaticVar(node *ast.StmtStaticVar) {
	node.Var.Accept(formatter)

	if node.Expr != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.EqualTkn = formatter.newToken('=', []byte("="))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

		node.Expr.Accept(formatter)
	}
}

func (formatter *formatter) StmtStmtList(node *ast.StmtStmtList) {
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtSwitch(node *ast.StmtSwitch) {
	node.CaseSeparatorTkn = nil
	node.ColonTkn = nil
	node.EndSwitchTkn = nil
	node.SemiColonTkn = nil

	node.SwitchTkn = formatter.newToken(token.T_SWITCH, []byte("switch"))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Cond.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Cases) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Cases)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtThrow(node *ast.StmtThrow) {
	node.ThrowTkn = formatter.newToken(token.T_THROW, []byte("throw"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtTrait(node *ast.StmtTrait) {
	node.TraitTkn = formatter.newToken(token.T_TRAIT, []byte("trait"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Name.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) StmtTraitUse(node *ast.StmtTraitUse) {
	node.UseTkn = formatter.newToken(token.T_USE, []byte("use"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = formatter.formatList(node.Traits, ',')

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	node.SemiColonTkn = nil

	if len(node.Adaptations) > 0 {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

		if len(node.Adaptations) > 0 {
			formatter.indent++
			formatter.formatStmts(&node.Adaptations)
			formatter.indent--

			formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			formatter.addIndent()
		}

		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	} else {
		node.SemiColonTkn = formatter.newToken(';', []byte(";"))
	}
}

func (formatter *formatter) StmtTraitUseAlias(node *ast.StmtTraitUseAlias) {
	if node.Trait != nil {
		node.Trait.Accept(formatter)
		node.DoubleColonTkn = formatter.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	}

	node.Method.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.AsTkn = formatter.newToken(token.T_AS, []byte("as"))

	if node.Modifier != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Modifier.Accept(formatter)
	}

	if node.Alias != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Alias.Accept(formatter)
	}

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtTraitUsePrecedence(node *ast.StmtTraitUsePrecedence) {
	if node.Trait != nil {
		node.Trait.Accept(formatter)
		node.DoubleColonTkn = formatter.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	}

	node.Method.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.InsteadofTkn = formatter.newToken(token.T_INSTEADOF, []byte("insteadof"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.SeparatorTkns = formatter.formatList(node.Insteadof, ',')

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtTry(node *ast.StmtTry) {
	node.TryTkn = formatter.newToken(token.T_TRY, []byte("try"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))

	for _, catch := range node.Catches {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		catch.Accept(formatter)
	}

	if node.Finally != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Finally.Accept(formatter)
	}
}

func (formatter *formatter) StmtUnset(node *ast.StmtUnset) {
	node.UnsetTkn = formatter.newToken(token.T_UNSET, []byte("unset"))

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = formatter.formatList(node.Vars, ',')
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtUse(node *ast.StmtUseList) {
	node.UseTkn = formatter.newToken(token.T_USE, []byte("use"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.SeparatorTkns = formatter.formatList(node.Uses, ',')

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtGroupUse(node *ast.StmtGroupUseList) {
	node.UseTkn = formatter.newToken(token.T_USE, []byte("use"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.LeadingNsSeparatorTkn = nil

	node.Prefix.Accept(formatter)
	node.NsSeparatorTkn = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
	node.SeparatorTkns = formatter.formatList(node.Uses, ',')
	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))

	node.SemiColonTkn = formatter.newSemicolonTkn()
}

func (formatter *formatter) StmtUseDeclaration(node *ast.StmtUse) {
	if node.Type != nil {
		node.Type.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.NsSeparatorTkn = nil

	node.Use.Accept(formatter)

	if node.Alias != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.AsTkn = formatter.newToken(token.T_AS, []byte("as"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.Alias.Accept(formatter)
	}
}

func (formatter *formatter) StmtWhile(node *ast.StmtWhile) {
	node.ColonTkn = nil
	node.EndWhileTkn = nil
	node.SemiColonTkn = nil

	node.WhileTkn = formatter.newToken(token.T_WHILE, []byte("while"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Cond.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Stmt.Accept(formatter)
}

func (formatter *formatter) ExprArray(node *ast.ExprArray) {
	node.ArrayTkn = formatter.newToken(token.T_ARRAY, []byte("array"))
	node.OpenBracketTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = formatter.formatList(node.Items, ',')
	node.CloseBracketTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprArrayDimFetch(node *ast.ExprArrayDimFetch) {
	node.Var.Accept(formatter)
	node.OpenBracketTkn = formatter.newToken('[', []byte("["))
	node.Dim.Accept(formatter)
	node.CloseBracketTkn = formatter.newToken(']', []byte("]"))
}

func (formatter *formatter) ExprArrayItem(node *ast.ExprArrayItem) {
	if node.EllipsisTkn != nil {
		node.EllipsisTkn = formatter.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	if node.Key != nil {
		node.Key.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.DoubleArrowTkn = formatter.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.Val.Accept(formatter)
}

func (formatter *formatter) ExprArrowFunction(node *ast.ExprArrowFunction) {
	if node.StaticTkn != nil {
		node.StaticTkn = formatter.newToken(token.T_STATIC, []byte("static"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.FnTkn = formatter.newToken(token.T_FN, []byte("fn"))

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Params) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Params, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	node.ColonTkn = nil
	if node.ReturnType != nil {
		node.ColonTkn = formatter.newToken(':', []byte(":"))

		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.ReturnType.Accept(formatter)
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.DoubleArrowTkn = formatter.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprBitwiseNot(node *ast.ExprBitwiseNot) {
	node.TildaTkn = formatter.newToken('~', []byte("~"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprBooleanNot(node *ast.ExprBooleanNot) {
	node.ExclamationTkn = formatter.newToken('!', []byte("!"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprBrackets(node *ast.ExprBrackets) {
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Expr.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprClassConstFetch(node *ast.ExprClassConstFetch) {
	node.Class.Accept(formatter)
	node.DoubleColonTkn = formatter.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	node.Const.Accept(formatter)
}

func (formatter *formatter) ExprClone(node *ast.ExprClone) {
	node.CloneTkn = formatter.newToken(token.T_CLONE, []byte("clone"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprThrow(node *ast.ExprThrow) {
	node.ThrowTkn = formatter.newToken(token.T_THROW, []byte("throw"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprClosure(node *ast.ExprClosure) {
	if node.StaticTkn != nil {
		node.StaticTkn = formatter.newToken(token.T_STATIC, []byte("static"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.FunctionTkn = formatter.newToken(token.T_FN, []byte("function"))

	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Params) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Params, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))

	node.UseTkn = nil
	node.UseOpenParenthesisTkn = nil
	node.UseCloseParenthesisTkn = nil
	node.UseSeparatorTkns = nil
	if len(node.Uses) > 0 {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.UseTkn = formatter.newToken(token.T_USE, []byte("use"))
		node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
		node.SeparatorTkns = formatter.formatList(node.Uses, ',')
		node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	}

	node.ColonTkn = nil
	if node.ReturnType != nil {
		node.ColonTkn = formatter.newToken(':', []byte(":"))

		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.ReturnType.Accept(formatter)
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
	if len(node.Stmts) > 0 {
		formatter.indent++
		formatter.formatStmts(&node.Stmts)
		formatter.indent--

		formatter.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		formatter.addIndent()
	}
	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) ExprClosureUse(node *ast.ExprClosureUse) {
	if node.AmpersandTkn != nil {
		node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	}

	node.Var.Accept(formatter)
}

func (formatter *formatter) ExprConstFetch(node *ast.ExprConstFetch) {
	node.Const.Accept(formatter)
}

func (formatter *formatter) ExprEmpty(node *ast.ExprEmpty) {
	node.EmptyTkn = formatter.newToken(token.T_EMPTY, []byte("empty"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Expr.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprErrorSuppress(node *ast.ExprErrorSuppress) {
	node.AtTkn = formatter.newToken('@', []byte("@"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprEval(node *ast.ExprEval) {
	node.EvalTkn = formatter.newToken(token.T_EVAL, []byte("eval"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Expr.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprExit(node *ast.ExprExit) {
	node.ExitTkn = formatter.newToken(token.T_EVAL, []byte("exit"))

	node.OpenParenthesisTkn = nil
	node.CloseParenthesisTkn = nil
	if node.Expr != nil {
		node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
		node.Expr.Accept(formatter)
		node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	}
}

func (formatter *formatter) ExprFunctionCall(node *ast.ExprFunctionCall) {
	node.Function.Accept(formatter)
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Args) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprInclude(node *ast.ExprInclude) {
	node.IncludeTkn = formatter.newToken(token.T_INCLUDE, []byte("include"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprIncludeOnce(node *ast.ExprIncludeOnce) {
	node.IncludeOnceTkn = formatter.newToken(token.T_INCLUDE_ONCE, []byte("include_once"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprInstanceOf(node *ast.ExprInstanceOf) {
	node.Expr.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.InstanceOfTkn = formatter.newToken(token.T_INSTANCEOF, []byte("instanceof"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Class.Accept(formatter)
}

func (formatter *formatter) ExprIsset(node *ast.ExprIsset) {
	node.IssetTkn = formatter.newToken(token.T_ISSET, []byte("isset"))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = formatter.formatList(node.Vars, ',')
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprList(node *ast.ExprList) {
	node.ListTkn = formatter.newToken(token.T_LIST, []byte("list"))
	node.OpenBracketTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = formatter.formatList(node.Items, ',')
	node.CloseBracketTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprMethodCall(node *ast.ExprMethodCall) {
	node.Var.Accept(formatter)
	node.ObjectOperatorTkn = formatter.newToken(token.T_OBJECT_OPERATOR, []byte("->"))

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Method.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Method.Accept(formatter)

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Args) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprNullsafeMethodCall(node *ast.ExprNullsafeMethodCall) {
	node.Var.Accept(formatter)
	node.ObjectOperatorTkn = formatter.newToken(token.T_NULLSAFE_OBJECT_OPERATOR, []byte("?->"))

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Method.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Method.Accept(formatter)

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Args) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprNullsafePropertyFetch(node *ast.ExprNullsafePropertyFetch) {
	node.Var.Accept(formatter)
	node.ObjectOperatorTkn = formatter.newToken(token.T_NULLSAFE_OBJECT_OPERATOR, []byte("?->"))

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Prop.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Prop.Accept(formatter)
}

func (formatter *formatter) ExprNew(node *ast.ExprNew) {
	node.NewTkn = formatter.newToken(token.T_NEW, []byte("new"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Class.Accept(formatter)

	node.SeparatorTkns = nil
	node.OpenParenthesisTkn = nil
	node.CloseParenthesisTkn = nil
	if len(node.Args) > 0 {
		node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
		node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	}
}

func (formatter *formatter) ExprPostDec(node *ast.ExprPostDec) {
	node.Var.Accept(formatter)
	node.DecTkn = formatter.newToken(token.T_DEC, []byte("--"))
}

func (formatter *formatter) ExprPostInc(node *ast.ExprPostInc) {
	node.Var.Accept(formatter)
	node.IncTkn = formatter.newToken(token.T_INC, []byte("++"))
}

func (formatter *formatter) ExprPreDec(node *ast.ExprPreDec) {
	node.DecTkn = formatter.newToken(token.T_DEC, []byte("--"))
	node.Var.Accept(formatter)
}

func (formatter *formatter) ExprPreInc(node *ast.ExprPreInc) {
	node.IncTkn = formatter.newToken(token.T_INC, []byte("++"))
	node.Var.Accept(formatter)
}

func (formatter *formatter) ExprPrint(node *ast.ExprPrint) {
	node.PrintTkn = formatter.newToken(token.T_PRINT, []byte("print"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprPropertyFetch(node *ast.ExprPropertyFetch) {
	node.Var.Accept(formatter)
	node.ObjectOperatorTkn = formatter.newToken(token.T_OBJECT_OPERATOR, []byte("->"))

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Prop.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Prop.Accept(formatter)
}

func (formatter *formatter) ExprRequire(node *ast.ExprRequire) {
	node.RequireTkn = formatter.newToken(token.T_REQUIRE, []byte("require"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprRequireOnce(node *ast.ExprRequireOnce) {
	node.RequireOnceTkn = formatter.newToken(token.T_REQUIRE_ONCE, []byte("require_once"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprShellExec(node *ast.ExprShellExec) {
	node.OpenBacktickTkn = formatter.newToken('`', []byte("`"))
	for _, part := range node.Parts {
		part.Accept(formatter)
	}
	node.CloseBacktickTkn = formatter.newToken('`', []byte("`"))
}

func (formatter *formatter) ExprStaticCall(node *ast.ExprStaticCall) {
	node.Class.Accept(formatter)
	node.DoubleColonTkn = formatter.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Call.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Call.Accept(formatter)

	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.SeparatorTkns = nil
	if len(node.Args) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Args, ',')
	}
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
}

func (formatter *formatter) ExprStaticPropertyFetch(node *ast.ExprStaticPropertyFetch) {
	node.Class.Accept(formatter)
	node.DoubleColonTkn = formatter.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	node.Prop.Accept(formatter)
}

func (formatter *formatter) ExprTernary(node *ast.ExprTernary) {
	node.Cond.Accept(formatter)
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.QuestionTkn = formatter.newToken('?', []byte("?"))
	if node.IfTrue != nil {
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.IfTrue.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}
	node.ColonTkn = formatter.newToken(':', []byte(":"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.IfFalse.Accept(formatter)
}

func (formatter *formatter) ExprMatch(node *ast.ExprMatch) {
	node.MatchTkn = formatter.newToken(token.T_MATCH, []byte("match"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenParenthesisTkn = formatter.newToken('(', []byte("("))
	node.Expr.Accept(formatter)
	node.CloseParenthesisTkn = formatter.newToken(')', []byte(")"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))

	node.SeparatorTkns = nil
	if len(node.Arms) > 0 {
		node.SeparatorTkns = formatter.formatList(node.Arms, ',')
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) MatchArm(node *ast.MatchArm) {
	if node.DefaultTkn != nil {
		node.DefaultTkn = formatter.newToken(token.T_DEFAULT, []byte("default"))
	} else {
		node.SeparatorTkns = nil
		if len(node.Exprs) > 0 {
			node.SeparatorTkns = formatter.formatList(node.Exprs, ',')
		}
	}

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.DoubleArrowTkn = formatter.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.ReturnExpr.Accept(formatter)
}

func (formatter *formatter) ExprUnaryMinus(node *ast.ExprUnaryMinus) {
	node.MinusTkn = formatter.newToken('-', []byte("-"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprUnaryPlus(node *ast.ExprUnaryPlus) {
	node.PlusTkn = formatter.newToken('+', []byte("+"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprVariable(node *ast.ExprVariable) {
	if _, ok := node.Name.(*ast.Identifier); !ok {
		node.DollarTkn = formatter.newToken('$', []byte("$"))
	}

	node.OpenCurlyBracketTkn = nil
	node.CloseCurlyBracketTkn = nil
	switch node.Name.(type) {
	case *ast.Identifier:
	case *ast.ExprVariable:
	default:
		node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
		node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
	}

	node.Name.Accept(formatter)
}

func (formatter *formatter) ExprYield(node *ast.ExprYield) {
	node.YieldTkn = formatter.newToken(token.T_YIELD, []byte("yield"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if node.Key != nil {
		node.Key.Accept(formatter)
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		node.DoubleArrowTkn = formatter.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
		formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	node.Val.Accept(formatter)
}

func (formatter *formatter) ExprYieldFrom(node *ast.ExprYieldFrom) {
	node.YieldFromTkn = formatter.newToken(token.T_YIELD_FROM, []byte("yield from"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssign(node *ast.ExprAssign) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken('=', []byte("="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignReference(node *ast.ExprAssignReference) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken('=', []byte("="))
	node.AmpersandTkn = formatter.newToken('&', []byte("&"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignBitwiseAnd(node *ast.ExprAssignBitwiseAnd) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_AND_EQUAL, []byte("&="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignBitwiseOr(node *ast.ExprAssignBitwiseOr) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_OR_EQUAL, []byte("|="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignBitwiseXor(node *ast.ExprAssignBitwiseXor) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_XOR_EQUAL, []byte("^="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignCoalesce(node *ast.ExprAssignCoalesce) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_COALESCE_EQUAL, []byte("??="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignConcat(node *ast.ExprAssignConcat) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_CONCAT_EQUAL, []byte(".="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignDiv(node *ast.ExprAssignDiv) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_DIV_EQUAL, []byte("/="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignMinus(node *ast.ExprAssignMinus) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_MINUS_EQUAL, []byte("-="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignMod(node *ast.ExprAssignMod) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_MOD_EQUAL, []byte("%="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignMul(node *ast.ExprAssignMul) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_MUL_EQUAL, []byte("*="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignPlus(node *ast.ExprAssignPlus) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_PLUS_EQUAL, []byte("+="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignPow(node *ast.ExprAssignPow) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_POW_EQUAL, []byte("**="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignShiftLeft(node *ast.ExprAssignShiftLeft) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_SL_EQUAL, []byte("<<="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprAssignShiftRight(node *ast.ExprAssignShiftRight) {
	node.Var.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.EqualTkn = formatter.newToken(token.T_SR_EQUAL, []byte(">>="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprBinaryBitwiseAnd(node *ast.ExprBinaryBitwiseAnd) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('&', []byte("&"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryBitwiseOr(node *ast.ExprBinaryBitwiseOr) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('|', []byte("|"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryBitwiseXor(node *ast.ExprBinaryBitwiseXor) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('^', []byte("^"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryBooleanAnd(node *ast.ExprBinaryBooleanAnd) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_BOOLEAN_AND, []byte("&&"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryBooleanOr(node *ast.ExprBinaryBooleanOr) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_BOOLEAN_OR, []byte("||"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryCoalesce(node *ast.ExprBinaryCoalesce) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_COALESCE, []byte("??"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryConcat(node *ast.ExprBinaryConcat) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('.', []byte("."))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryDiv(node *ast.ExprBinaryDiv) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('/', []byte("/"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryEqual(node *ast.ExprBinaryEqual) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_EQUAL, []byte("=="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryGreater(node *ast.ExprBinaryGreater) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('>', []byte(">"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryGreaterOrEqual(node *ast.ExprBinaryGreaterOrEqual) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_GREATER_OR_EQUAL, []byte(">="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryIdentical(node *ast.ExprBinaryIdentical) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_IDENTICAL, []byte("==="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryLogicalAnd(node *ast.ExprBinaryLogicalAnd) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_LOGICAL_AND, []byte("and"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryLogicalOr(node *ast.ExprBinaryLogicalOr) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_LOGICAL_OR, []byte("or"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryLogicalXor(node *ast.ExprBinaryLogicalXor) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_LOGICAL_XOR, []byte("xor"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryMinus(node *ast.ExprBinaryMinus) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('-', []byte("-"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryMod(node *ast.ExprBinaryMod) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('%', []byte("%"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryMul(node *ast.ExprBinaryMul) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('*', []byte("*"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryNotEqual(node *ast.ExprBinaryNotEqual) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_NOT_EQUAL, []byte("!="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryNotIdentical(node *ast.ExprBinaryNotIdentical) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_NOT_IDENTICAL, []byte("!=="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryPlus(node *ast.ExprBinaryPlus) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('+', []byte("+"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryPow(node *ast.ExprBinaryPow) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_POW, []byte("**"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryShiftLeft(node *ast.ExprBinaryShiftLeft) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_SL, []byte("<<"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinaryShiftRight(node *ast.ExprBinaryShiftRight) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_SR, []byte(">>"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinarySmaller(node *ast.ExprBinarySmaller) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken('<', []byte("<"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinarySmallerOrEqual(node *ast.ExprBinarySmallerOrEqual) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_IS_SMALLER_OR_EQUAL, []byte("<="))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprBinarySpaceship(node *ast.ExprBinarySpaceship) {
	node.Left.Accept(formatter)

	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	node.OpTkn = formatter.newToken(token.T_SPACESHIP, []byte("<=>"))
	formatter.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	node.Right.Accept(formatter)
}

func (formatter *formatter) ExprCastArray(node *ast.ExprCastArray) {
	node.CastTkn = formatter.newToken(token.T_ARRAY_CAST, []byte("(array)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastBool(node *ast.ExprCastBool) {
	node.CastTkn = formatter.newToken(token.T_BOOL_CAST, []byte("(bool)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastDouble(node *ast.ExprCastDouble) {
	node.CastTkn = formatter.newToken(token.T_DOUBLE_CAST, []byte("(float)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastInt(node *ast.ExprCastInt) {
	node.CastTkn = formatter.newToken(token.T_INT_CAST, []byte("(int)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastObject(node *ast.ExprCastObject) {
	node.CastTkn = formatter.newToken(token.T_OBJECT_CAST, []byte("(object)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastString(node *ast.ExprCastString) {
	node.CastTkn = formatter.newToken(token.T_STRING_CAST, []byte("(string)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ExprCastUnset(node *ast.ExprCastUnset) {
	node.CastTkn = formatter.newToken(token.T_UNSET_CAST, []byte("(unset)"))
	node.Expr.Accept(formatter)
}

func (formatter *formatter) ScalarDnumber(node *ast.ScalarDnumber) {
	if node.NumberTkn == nil {
		node.NumberTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.NumberTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) ScalarEncapsed(node *ast.ScalarEncapsed) {
	node.OpenQuoteTkn = formatter.newToken('"', []byte("\""))
	for _, part := range node.Parts {
		part.Accept(formatter)
	}
	node.CloseQuoteTkn = formatter.newToken('"', []byte("\""))
}

func (formatter *formatter) ScalarEncapsedStringPart(node *ast.ScalarEncapsedStringPart) {
	if node.EncapsedStrTkn == nil {
		node.EncapsedStrTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.EncapsedStrTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) ScalarEncapsedStringVar(node *ast.ScalarEncapsedStringVar) {
	node.DollarOpenCurlyBracketTkn = formatter.newToken(token.T_DOLLAR_OPEN_CURLY_BRACES, []byte("${"))
	node.Name.Accept(formatter)

	node.OpenSquareBracketTkn = nil
	node.CloseSquareBracketTkn = nil
	if node.Dim != nil {
		node.OpenSquareBracketTkn = formatter.newToken('[', []byte("["))
		node.Dim.Accept(formatter)
		node.CloseSquareBracketTkn = formatter.newToken(']', []byte("]"))
	}

	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) ScalarEncapsedStringBrackets(node *ast.ScalarEncapsedStringBrackets) {
	node.OpenCurlyBracketTkn = formatter.newToken('{', []byte("{"))
	node.Var.Accept(formatter)
	node.CloseCurlyBracketTkn = formatter.newToken('}', []byte("}"))
}

func (formatter *formatter) ScalarHeredoc(node *ast.ScalarHeredoc) {
	node.OpenHeredocTkn = formatter.newToken(token.T_START_HEREDOC, []byte("<<<EOT\n"))
	for _, part := range node.Parts {
		part.Accept(formatter)
	}
	node.CloseHeredocTkn = formatter.newToken(token.T_START_HEREDOC, []byte("EOT"))
}

func (formatter *formatter) ScalarLnumber(node *ast.ScalarLnumber) {
	if node.NumberTkn == nil {
		node.NumberTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.NumberTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) ScalarMagicConstant(node *ast.ScalarMagicConstant) {
	if node.MagicConstTkn == nil {
		node.MagicConstTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.MagicConstTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) ScalarString(node *ast.ScalarString) {
	if node.StringTkn == nil {
		node.StringTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.StringTkn.FreeFloating = formatter.getFreeFloating()
	}
}

func (formatter *formatter) NameName(node *ast.Name) {
	separatorTokens := make([]*token.Token, len(node.Parts)-1)
	for i, part := range node.Parts {
		part.Accept(formatter)

		if i != len(node.Parts)-1 {
			separatorTokens[i] = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (formatter *formatter) NameFullyQualified(node *ast.NameFullyQualified) {
	node.NsSeparatorTkn = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	separatorTokens := make([]*token.Token, len(node.Parts)-1)
	for i, part := range node.Parts {
		part.Accept(formatter)

		if i != len(node.Parts)-1 {
			separatorTokens[i] = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (formatter *formatter) NameRelative(node *ast.NameRelative) {
	node.NsTkn = formatter.newToken(token.T_NAMESPACE, []byte("namespace"))
	node.NsSeparatorTkn = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	separatorTokens := make([]*token.Token, len(node.Parts)-1)
	for i, part := range node.Parts {
		part.Accept(formatter)

		if i != len(node.Parts)-1 {
			separatorTokens[i] = formatter.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (formatter *formatter) NameNamePart(node *ast.NamePart) {
	if node.StringTkn == nil {
		node.StringTkn = formatter.newToken(token.T_STRING, node.Value)
	} else {
		node.StringTkn.FreeFloating = formatter.getFreeFloating()
	}
}
