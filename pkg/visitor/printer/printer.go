package printer

import (
	"bytes"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"io"
)

type printerState int

const (
	PrinterStateHTML printerState = iota
	PrinterStatePHP
)

type printer struct {
	output io.Writer
	state  printerState
	last   []byte
}

func NewPrinter(output io.Writer) *printer {
	return &printer{
		output: output,
	}
}

func (printer *printer) WithState(state printerState) *printer {
	printer.state = state
	return printer
}

func isValidVarName(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '_' || char >= 0x80
}

func (printer *printer) write(content []byte) {
	if len(content) == 0 {
		return
	}

	if printer.state == PrinterStateHTML {
		// A shebang line can only open the very first line of a file, where it
		// precedes the open tag. It is not PHP code, so it is written as is and
		// the printer stays in HTML state until the open tag arrives.
		if printer.last == nil && bytes.HasPrefix(content, []byte("#!")) {
			printer.last = content
			printer.output.Write(content)

			return
		}

		if !bytes.HasPrefix(content, []byte("<?")) {
			printer.output.Write([]byte("<?php "))
		}
		printer.state = PrinterStatePHP
	}

	if printer.last != nil && isValidVarName(printer.last[len(printer.last)-1]) && isValidVarName(content[0]) {
		printer.output.Write([]byte(" "))
	}

	printer.last = content
	printer.output.Write(content)
}

func (printer *printer) printNode(node ast.Vertex) {
	if node != nil {
		node.Accept(printer)
	}
}

func (printer *printer) printList(list []ast.Vertex) {
	for _, childNode := range list {
		printer.printNode(childNode)
	}
}

func (printer *printer) printSeparatedList(list []ast.Vertex, separators []*token.Token, defaultSeparator []byte) {
	for k, childNode := range list {
		printer.printNode(childNode)
		if k < len(separators) {
			printer.printToken(separators[k], defaultSeparator)
		} else if k < len(list)-1 {
			printer.write(defaultSeparator)
		}
	}
}

func (printer *printer) printToken(tokenItem *token.Token, defaultValue []byte) {
	if tokenItem == nil && defaultValue == nil {
		return
	}

	if tokenItem == nil {
		printer.write(defaultValue)
		return
	}

	for _, freeFloating := range tokenItem.FreeFloating {
		printer.write(freeFloating.Value)
	}
	printer.write(tokenItem.Value)
}

func (printer *printer) ifNode(node ast.Vertex, value []byte) []byte {
	if node == nil {
		return nil
	}

	return value
}

func (printer *printer) ifNodeList(nodes []ast.Vertex, value []byte) []byte {
	if nodes == nil {
		return nil
	}

	return value
}

func (printer *printer) ifNotNodeList(nodes []ast.Vertex, value []byte) []byte {
	if nodes != nil {
		return nil
	}

	return value
}

func (printer *printer) ifToken(tokenItem *token.Token, trueValue []byte, falseValue []byte) []byte {
	if tokenItem == nil {
		return falseValue
	}

	return trueValue
}

func (printer *printer) ifNotToken(tokenItem *token.Token, value []byte) []byte {
	if tokenItem != nil {
		return nil
	}

	return value
}

func (printer *printer) Root(node *ast.Root) {
	printer.printList(node.Stmts)
	printer.printToken(node.EndTkn, nil)
}

func (printer *printer) Nullable(node *ast.Nullable) {
	printer.printToken(node.QuestionTkn, []byte("?"))
	printer.printNode(node.Expr)
}

func (printer *printer) Union(node *ast.Union) {
	printer.printSeparatedList(node.Types, node.SeparatorTkns, []byte("|"))
}

func (printer *printer) Intersection(node *ast.Intersection) {
	printer.printToken(node.OpenParenthesisTkn, nil)
	printer.printSeparatedList(node.Types, node.SeparatorTkns, []byte("&"))
	printer.printToken(node.CloseParenthesisTkn, nil)
}

func (printer *printer) Parameter(node *ast.Parameter) {
	printer.printList(node.AttrGroups)
	printer.printList(node.Modifiers)
	printer.printNode(node.Type)
	printer.printToken(node.AmpersandTkn, nil)
	printer.printToken(node.VariadicTkn, nil)
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, printer.ifNode(node.DefaultValue, []byte("=")))
	printer.printNode(node.DefaultValue)
}

func (printer *printer) Identifier(node *ast.Identifier) {
	printer.printToken(node.IdentifierTkn, node.Value)
}

func (printer *printer) Argument(node *ast.Argument) {
	printer.printNode(node.Name)
	printer.printToken(node.ColonTkn, printer.ifNode(node.Name, []byte(":")))
	printer.printToken(node.VariadicTkn, nil)
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Expr)
}

func (printer *printer) Attribute(node *ast.Attribute) {
	printer.printNode(node.Name)
	printer.printToken(node.OpenParenthesisTkn, printer.ifNodeList(node.Args, []byte("(")))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, printer.ifNodeList(node.Args, []byte(")")))
}

func (printer *printer) AttributeGroup(node *ast.AttributeGroup) {
	printer.printToken(node.OpenAttributeTkn, []byte("#["))
	printer.printSeparatedList(node.Attrs, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseAttributeTkn, []byte("]"))
}

func (printer *printer) StmtBreak(node *ast.StmtBreak) {
	printer.printToken(node.BreakTkn, []byte("break"))
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtCase(node *ast.StmtCase) {
	printer.printToken(node.CaseTkn, []byte("case"))
	printer.printNode(node.Cond)
	printer.printToken(node.CaseSeparatorTkn, []byte(":"))
	printer.printList(node.Stmts)
}

func (printer *printer) StmtCatch(node *ast.StmtCatch) {
	printer.printToken(node.CatchTkn, []byte("catch"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Types, node.SeparatorTkns, []byte("|"))
	printer.printNode(node.Var)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtClass(node *ast.StmtClass) {
	printer.printList(node.AttrGroups)
	printer.printList(node.Modifiers)
	printer.printToken(node.ClassTkn, []byte("class"))
	printer.printNode(node.Name)
	printer.printToken(node.OpenParenthesisTkn, printer.ifNodeList(node.Args, []byte("(")))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, printer.ifNodeList(node.Args, []byte(")")))
	printer.printToken(node.ExtendsTkn, printer.ifNode(node.Extends, []byte("extends")))
	printer.printNode(node.Extends)
	printer.printToken(node.ImplementsTkn, printer.ifNodeList(node.Implements, []byte("implements")))
	printer.printSeparatedList(node.Implements, node.ImplementsSeparatorTkns, []byte(","))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtClassConstList(node *ast.StmtClassConstList) {
	printer.printList(node.AttrGroups)
	printer.printList(node.Modifiers)
	printer.printToken(node.ConstTkn, []byte("const"))
	printer.printNode(node.Type)
	printer.printSeparatedList(node.Consts, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtClassMethod(node *ast.StmtClassMethod) {
	printer.printList(node.AttrGroups)
	printer.printList(node.Modifiers)
	printer.printToken(node.FunctionTkn, []byte("function"))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Name)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Params, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, printer.ifNode(node.ReturnType, []byte(":")))
	printer.printNode(node.ReturnType)
	printer.printNode(node.Stmt)
}

func (printer *printer) StmtConstList(node *ast.StmtConstList) {
	printer.printToken(node.ConstTkn, []byte("const"))
	printer.printSeparatedList(node.Consts, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtConstant(node *ast.StmtConstant) {
	printer.printNode(node.Name)
	printer.printToken(node.EqualTkn, []byte("="))
	printer.printNode(node.Expr)
}

func (printer *printer) StmtContinue(node *ast.StmtContinue) {
	printer.printToken(node.ContinueTkn, []byte("continue"))
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtDeclare(node *ast.StmtDeclare) {
	printer.printToken(node.DeclareTkn, []byte("declare"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Consts, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
	printer.printToken(node.EndDeclareTkn, printer.ifToken(node.ColonTkn, []byte("enddeclare"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) StmtDefault(node *ast.StmtDefault) {
	printer.printToken(node.DefaultTkn, []byte("default"))
	printer.printToken(node.CaseSeparatorTkn, []byte(":"))
	printer.printList(node.Stmts)
}

func (printer *printer) StmtDo(node *ast.StmtDo) {
	printer.printToken(node.DoTkn, []byte("do"))
	printer.printNode(node.Stmt)
	printer.printToken(node.WhileTkn, []byte("while"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Cond)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.SemiColonTkn, []byte(";"))

}

func (printer *printer) StmtEcho(node *ast.StmtEcho) {
	printer.printToken(node.EchoTkn, []byte("echo"))
	printer.printSeparatedList(node.Exprs, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtElse(node *ast.StmtElse) {
	printer.printToken(node.ElseTkn, []byte("else"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
}

func (printer *printer) StmtElseIf(node *ast.StmtElseIf) {
	printer.printToken(node.ElseIfTkn, []byte("elseif"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Cond)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
}

func (printer *printer) StmtExpression(node *ast.StmtExpression) {
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtFinally(node *ast.StmtFinally) {
	printer.printToken(node.FinallyTkn, []byte("finally"))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtFor(node *ast.StmtFor) {
	printer.printToken(node.ForTkn, []byte("for"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Init, node.InitSeparatorTkns, []byte(","))
	printer.printToken(node.InitSemiColonTkn, []byte(";"))
	printer.printSeparatedList(node.Cond, node.CondSeparatorTkns, []byte(","))
	printer.printToken(node.CondSemiColonTkn, []byte(";"))
	printer.printSeparatedList(node.Loop, node.LoopSeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
	printer.printToken(node.EndForTkn, printer.ifToken(node.ColonTkn, []byte("endfor"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) StmtForeach(node *ast.StmtForeach) {
	printer.printToken(node.ForeachTkn, []byte("foreach"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Expr)
	printer.printToken(node.AsTkn, []byte("as"))
	printer.printNode(node.Key)
	printer.printToken(node.DoubleArrowTkn, printer.ifNode(node.Key, []byte("=>")))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Var)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
	printer.printToken(node.EndForeachTkn, printer.ifToken(node.ColonTkn, []byte("endforeach"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) StmtFunction(node *ast.StmtFunction) {
	printer.printList(node.AttrGroups)
	printer.printToken(node.FunctionTkn, []byte("function"))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Name)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Params, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, printer.ifNode(node.ReturnType, []byte(":")))
	printer.printNode(node.ReturnType)
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtGlobal(node *ast.StmtGlobal) {
	printer.printToken(node.GlobalTkn, []byte("global"))
	printer.printSeparatedList(node.Vars, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtGoto(node *ast.StmtGoto) {
	printer.printToken(node.GotoTkn, []byte("goto"))
	printer.printNode(node.Label)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtHaltCompiler(node *ast.StmtHaltCompiler) {
	printer.printToken(node.HaltCompilerTkn, []byte("__halt_compiler"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtIf(node *ast.StmtIf) {
	printer.printToken(node.IfTkn, []byte("if"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Cond)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
	printer.printList(node.ElseIf)
	printer.printNode(node.Else)
	printer.printToken(node.EndIfTkn, printer.ifToken(node.ColonTkn, []byte("endif"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) StmtInlineHtml(node *ast.StmtInlineHtml) {
	printer.state = PrinterStatePHP
	if printer.last != nil && !bytes.HasSuffix(printer.last, []byte("?>")) && !bytes.HasSuffix(printer.last, []byte("?>\n")) {
		printer.write([]byte("?>"))
	}

	printer.printToken(node.InlineHtmlTkn, node.Value)
	printer.state = PrinterStateHTML
}

func (printer *printer) StmtEnum(node *ast.StmtEnum) {
	printer.printList(node.AttrGroups)
	printer.printToken(node.EnumTkn, []byte("enum"))
	printer.printNode(node.Name)
	printer.printToken(node.ColonTkn, printer.ifNode(node.Type, []byte(":")))
	printer.printNode(node.Type)
	printer.printToken(node.ImplementsTkn, printer.ifNodeList(node.Implements, []byte("implements")))
	printer.printSeparatedList(node.Implements, node.ImplementsSeparatorTkns, []byte(","))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtEnumCase(node *ast.StmtEnumCase) {
	printer.printList(node.AttrGroups)
	printer.printToken(node.CaseTkn, []byte("case"))
	printer.printNode(node.Name)
	printer.printToken(node.EqualTkn, printer.ifNode(node.Expr, []byte("=")))
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtInterface(node *ast.StmtInterface) {
	printer.printList(node.AttrGroups)
	printer.printToken(node.InterfaceTkn, []byte("interface"))
	printer.printNode(node.Name)
	printer.printToken(node.ExtendsTkn, printer.ifNodeList(node.Extends, []byte("extends")))
	printer.printSeparatedList(node.Extends, node.ExtendsSeparatorTkns, []byte(","))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtLabel(node *ast.StmtLabel) {
	printer.printNode(node.Name)
	printer.printToken(node.ColonTkn, []byte(":"))
}

func (printer *printer) StmtNamespace(node *ast.StmtNamespace) {
	printer.printToken(node.NsTkn, []byte("namespace"))
	printer.printNode(node.Name)
	printer.printToken(node.OpenCurlyBracketTkn, printer.ifNodeList(node.Stmts, []byte("{")))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, printer.ifNodeList(node.Stmts, []byte("}")))
	printer.printToken(node.SemiColonTkn, printer.ifNotNodeList(node.Stmts, []byte(";")))
}

func (printer *printer) StmtNop(node *ast.StmtNop) {
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtProperty(node *ast.StmtProperty) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, printer.ifNode(node.Expr, []byte("=")))
	printer.printNode(node.Expr)
}

func (printer *printer) StmtPropertyList(node *ast.StmtPropertyList) {
	printer.printList(node.AttrGroups)
	printer.printList(node.Modifiers)
	printer.printNode(node.Type)
	printer.printSeparatedList(node.Props, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtReturn(node *ast.StmtReturn) {
	printer.printToken(node.ReturnTkn, []byte("return"))
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtStatic(node *ast.StmtStatic) {
	printer.printToken(node.StaticTkn, []byte("static"))
	printer.printSeparatedList(node.Vars, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtStaticVar(node *ast.StmtStaticVar) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, printer.ifNode(node.Expr, []byte("=")))
	printer.printNode(node.Expr)
}

func (printer *printer) StmtStmtList(node *ast.StmtStmtList) {
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtSwitch(node *ast.StmtSwitch) {
	printer.printToken(node.SwitchTkn, []byte("switch"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Cond)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	printer.printToken(node.OpenCurlyBracketTkn, printer.ifNotToken(node.ColonTkn, []byte("{")))
	printer.printToken(node.CaseSeparatorTkn, nil)
	printer.printList(node.Cases)
	printer.printToken(node.CloseCurlyBracketTkn, printer.ifNotToken(node.ColonTkn, []byte("}")))
	printer.printToken(node.EndSwitchTkn, printer.ifToken(node.ColonTkn, []byte("endswitch"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) StmtThrow(node *ast.StmtThrow) {
	printer.printToken(node.ThrowTkn, []byte("throw"))
	printer.printNode(node.Expr)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtTrait(node *ast.StmtTrait) {
	printer.printList(node.AttrGroups)
	printer.printToken(node.TraitTkn, []byte("trait"))
	printer.printNode(node.Name)
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) StmtTraitUse(node *ast.StmtTraitUse) {
	printer.printToken(node.UseTkn, []byte("use"))
	printer.printSeparatedList(node.Traits, node.SeparatorTkns, []byte(","))
	printer.printToken(node.OpenCurlyBracketTkn, printer.ifNodeList(node.Adaptations, []byte("{")))
	printer.printList(node.Adaptations)
	printer.printToken(node.CloseCurlyBracketTkn, printer.ifNodeList(node.Adaptations, []byte("}")))
	printer.printToken(node.SemiColonTkn, printer.ifNotToken(node.OpenCurlyBracketTkn, printer.ifNotNodeList(node.Adaptations, []byte(";"))))
}

func (printer *printer) StmtTraitUseAlias(node *ast.StmtTraitUseAlias) {
	printer.printNode(node.Trait)
	printer.printToken(node.DoubleColonTkn, printer.ifNode(node.Trait, []byte("::")))
	printer.printNode(node.Method)
	printer.printToken(node.AsTkn, []byte("as"))
	printer.printNode(node.Modifier)
	printer.printNode(node.Alias)
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtTraitUsePrecedence(node *ast.StmtTraitUsePrecedence) {
	printer.printNode(node.Trait)
	printer.printToken(node.DoubleColonTkn, printer.ifNode(node.Trait, []byte("::")))
	printer.printNode(node.Method)
	printer.printToken(node.InsteadofTkn, []byte("insteadof"))
	printer.printSeparatedList(node.Insteadof, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtTry(node *ast.StmtTry) {
	printer.printToken(node.TryTkn, []byte("try"))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
	printer.printList(node.Catches)
	printer.printNode(node.Finally)
}

func (printer *printer) StmtUnset(node *ast.StmtUnset) {
	printer.printToken(node.UnsetTkn, []byte("unset"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Vars, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtUse(node *ast.StmtUseList) {
	printer.printToken(node.UseTkn, []byte("use"))
	printer.printNode(node.Type)
	printer.printSeparatedList(node.Uses, node.SeparatorTkns, []byte(","))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtGroupUse(node *ast.StmtGroupUseList) {
	printer.printToken(node.UseTkn, []byte("use"))
	printer.printNode(node.Type)
	printer.printToken(node.LeadingNsSeparatorTkn, nil)
	printer.printNode(node.Prefix)
	printer.printToken(node.NsSeparatorTkn, []byte("\\"))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printSeparatedList(node.Uses, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
	printer.printToken(node.SemiColonTkn, []byte(";"))
}

func (printer *printer) StmtUseDeclaration(node *ast.StmtUse) {
	printer.printNode(node.Type)
	printer.printToken(node.NsSeparatorTkn, nil)
	printer.printNode(node.Use)
	printer.printToken(node.AsTkn, printer.ifNode(node.Alias, []byte("as")))
	printer.printNode(node.Alias)
}

func (printer *printer) StmtWhile(node *ast.StmtWhile) {
	printer.printToken(node.WhileTkn, []byte("while"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Cond)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, nil)
	if statementList, ok := node.Stmt.(*ast.StmtStmtList); ok && node.ColonTkn != nil {
		printer.printToken(statementList.OpenCurlyBracketTkn, nil)
		printer.printList(statementList.Stmts)
		printer.printToken(statementList.CloseCurlyBracketTkn, nil)
	} else {
		printer.printNode(node.Stmt)
	}
	printer.printToken(node.EndWhileTkn, printer.ifToken(node.ColonTkn, []byte("endwhile"), nil))
	printer.printToken(node.SemiColonTkn, printer.ifToken(node.ColonTkn, []byte(";"), nil))
}

func (printer *printer) ExprArray(node *ast.ExprArray) {
	printer.printToken(node.ArrayTkn, nil)
	printer.printToken(node.OpenBracketTkn, printer.ifToken(node.ArrayTkn, []byte("("), []byte("[")))
	printer.printSeparatedList(node.Items, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseBracketTkn, printer.ifToken(node.ArrayTkn, []byte(")"), []byte("]")))
}

func (printer *printer) ExprArrayDimFetch(node *ast.ExprArrayDimFetch) {
	printer.printNode(node.Var)
	printer.printToken(node.OpenBracketTkn, []byte("["))
	printer.printNode(node.Dim)
	printer.printToken(node.CloseBracketTkn, []byte("]"))
}

func (printer *printer) ExprArrayItem(node *ast.ExprArrayItem) {
	printer.printToken(node.EllipsisTkn, nil)
	printer.printNode(node.Key)
	printer.printToken(node.DoubleArrowTkn, printer.ifNode(node.Key, []byte("=>")))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Val)
}

func (printer *printer) ExprArrowFunction(node *ast.ExprArrowFunction) {
	printer.printToken(node.StaticTkn, nil)
	printer.printToken(node.FnTkn, []byte("fn"))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Params, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.ColonTkn, printer.ifNode(node.ReturnType, []byte(":")))
	printer.printNode(node.ReturnType)
	printer.printToken(node.DoubleArrowTkn, []byte("=>"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprBitwiseNot(node *ast.ExprBitwiseNot) {
	printer.printToken(node.TildaTkn, []byte("~"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprBooleanNot(node *ast.ExprBooleanNot) {
	printer.printToken(node.ExclamationTkn, []byte("!"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprBrackets(node *ast.ExprBrackets) {
	printer.printToken(node.OpenParenthesisTkn, nil)
	printer.printNode(node.Expr)
	printer.printToken(node.CloseParenthesisTkn, nil)
}

func (printer *printer) ExprClassConstFetch(node *ast.ExprClassConstFetch) {
	printer.printNode(node.Class)
	printer.printToken(node.DoubleColonTkn, []byte("::"))
	printer.printNode(node.Const)
}

func (printer *printer) ExprClone(node *ast.ExprClone) {
	printer.printToken(node.CloneTkn, []byte("clone"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprThrow(node *ast.ExprThrow) {
	printer.printToken(node.ThrowTkn, []byte("throw"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprClosure(node *ast.ExprClosure) {
	printer.printToken(node.StaticTkn, nil)
	printer.printToken(node.FunctionTkn, []byte("function"))
	printer.printToken(node.AmpersandTkn, nil)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Params, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.UseTkn, printer.ifNodeList(node.Uses, []byte("use")))
	printer.printToken(node.UseOpenParenthesisTkn, printer.ifNodeList(node.Uses, []byte("(")))
	printer.printSeparatedList(node.Uses, node.UseSeparatorTkns, []byte(","))
	printer.printToken(node.UseCloseParenthesisTkn, printer.ifNodeList(node.Uses, []byte(")")))
	printer.printToken(node.ColonTkn, printer.ifNode(node.ReturnType, []byte(":")))
	printer.printNode(node.ReturnType)
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printList(node.Stmts)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) ExprClosureUse(node *ast.ExprClosureUse) {
	printer.printToken(node.AmpersandTkn, nil)
	printer.printNode(node.Var)
}

func (printer *printer) ExprConstFetch(node *ast.ExprConstFetch) {
	printer.printNode(node.Const)
}

func (printer *printer) ExprEmpty(node *ast.ExprEmpty) {
	printer.printToken(node.EmptyTkn, []byte("empty"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Expr)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprErrorSuppress(node *ast.ExprErrorSuppress) {
	printer.printToken(node.AtTkn, []byte("@"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprEval(node *ast.ExprEval) {
	printer.printToken(node.EvalTkn, []byte("eval"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Expr)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprExit(node *ast.ExprExit) {
	printer.printToken(node.ExitTkn, []byte("exit"))
	printer.printToken(node.OpenParenthesisTkn, nil)
	printer.printNode(node.Expr)
	printer.printToken(node.CloseParenthesisTkn, printer.ifToken(node.OpenParenthesisTkn, []byte(")"), nil))
}

func (printer *printer) ExprFunctionCall(node *ast.ExprFunctionCall) {
	printer.printNode(node.Function)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprInclude(node *ast.ExprInclude) {
	printer.printToken(node.IncludeTkn, []byte("include"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprIncludeOnce(node *ast.ExprIncludeOnce) {
	printer.printToken(node.IncludeOnceTkn, []byte("include_once"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprInstanceOf(node *ast.ExprInstanceOf) {
	printer.printNode(node.Expr)
	printer.printToken(node.InstanceOfTkn, []byte("instanceof"))
	printer.printNode(node.Class)
}

func (printer *printer) ExprIsset(node *ast.ExprIsset) {
	printer.printToken(node.IssetTkn, []byte("isset"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Vars, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprList(node *ast.ExprList) {
	printer.printToken(node.ListTkn, printer.ifToken(node.OpenBracketTkn, nil, []byte("list")))
	printer.printToken(node.OpenBracketTkn, []byte("("))
	printer.printSeparatedList(node.Items, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseBracketTkn, []byte(")"))
}

func (printer *printer) ExprMethodCall(node *ast.ExprMethodCall) {
	printer.printNode(node.Var)
	printer.printToken(node.ObjectOperatorTkn, []byte("->"))
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Method)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprNullsafeMethodCall(node *ast.ExprNullsafeMethodCall) {
	printer.printNode(node.Var)
	printer.printToken(node.ObjectOperatorTkn, []byte("?->"))
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Method)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
}

func (printer *printer) ExprNullsafePropertyFetch(node *ast.ExprNullsafePropertyFetch) {
	printer.printNode(node.Var)
	printer.printToken(node.ObjectOperatorTkn, []byte("?->"))
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Prop)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
}

func (printer *printer) ExprNew(node *ast.ExprNew) {
	printer.printToken(node.NewTkn, []byte("new"))
	printer.printNode(node.Class)
	printer.printToken(node.OpenParenthesisTkn, printer.ifNodeList(node.Args, []byte("(")))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, printer.ifNodeList(node.Args, []byte(")")))
}

func (printer *printer) ExprPostDec(node *ast.ExprPostDec) {
	printer.printNode(node.Var)
	printer.printToken(node.DecTkn, []byte("--"))
}

func (printer *printer) ExprPostInc(node *ast.ExprPostInc) {
	printer.printNode(node.Var)
	printer.printToken(node.IncTkn, []byte("++"))
}

func (printer *printer) ExprPreDec(node *ast.ExprPreDec) {
	printer.printToken(node.DecTkn, []byte("--"))
	printer.printNode(node.Var)
}

func (printer *printer) ExprPreInc(node *ast.ExprPreInc) {
	printer.printToken(node.IncTkn, []byte("++"))
	printer.printNode(node.Var)
}

func (printer *printer) ExprPrint(node *ast.ExprPrint) {
	printer.printToken(node.PrintTkn, []byte("print"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprPropertyFetch(node *ast.ExprPropertyFetch) {
	printer.printNode(node.Var)
	printer.printToken(node.ObjectOperatorTkn, []byte("->"))
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Prop)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
}

func (printer *printer) ExprRequire(node *ast.ExprRequire) {
	printer.printToken(node.RequireTkn, []byte("require"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprRequireOnce(node *ast.ExprRequireOnce) {
	printer.printToken(node.RequireOnceTkn, []byte("require_once"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprShellExec(node *ast.ExprShellExec) {
	printer.printToken(node.OpenBacktickTkn, []byte("`"))
	printer.printList(node.Parts)
	printer.printToken(node.CloseBacktickTkn, []byte("`"))
}

func (printer *printer) ExprStaticCall(node *ast.ExprStaticCall) {
	printer.printNode(node.Class)
	printer.printToken(node.DoubleColonTkn, []byte("::"))
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Call)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
	printer.printToken(node.OpenParenthesisTkn, printer.ifNodeList(node.Args, []byte("(")))
	printer.printSeparatedList(node.Args, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseParenthesisTkn, printer.ifNodeList(node.Args, []byte(")")))
}

func (printer *printer) ExprStaticPropertyFetch(node *ast.ExprStaticPropertyFetch) {
	printer.printNode(node.Class)
	printer.printToken(node.DoubleColonTkn, []byte("::"))
	printer.printNode(node.Prop)
}

func (printer *printer) ExprTernary(node *ast.ExprTernary) {
	printer.printNode(node.Cond)
	printer.printToken(node.QuestionTkn, []byte("?"))
	printer.printNode(node.IfTrue)
	printer.printToken(node.ColonTkn, []byte(":"))
	printer.printNode(node.IfFalse)
}

func (printer *printer) ExprMatch(node *ast.ExprMatch) {
	printer.printToken(node.MatchTkn, []byte("match"))
	printer.printToken(node.OpenParenthesisTkn, []byte("("))
	printer.printNode(node.Expr)
	printer.printToken(node.CloseParenthesisTkn, []byte(")"))
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printSeparatedList(node.Arms, node.SeparatorTkns, []byte(","))
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) MatchArm(node *ast.MatchArm) {
	printer.printToken(node.DefaultTkn, printer.ifNotNodeList(node.Exprs, []byte("default")))
	printer.printSeparatedList(node.Exprs, node.SeparatorTkns, []byte(","))
	printer.printToken(node.DoubleArrowTkn, []byte("=>"))
	printer.printNode(node.ReturnExpr)
}

func (printer *printer) ExprUnaryMinus(node *ast.ExprUnaryMinus) {
	printer.printToken(node.MinusTkn, []byte("-"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprUnaryPlus(node *ast.ExprUnaryPlus) {
	printer.printToken(node.PlusTkn, []byte("+"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprVariable(node *ast.ExprVariable) {
	printer.printToken(node.DollarTkn, nil)
	printer.printToken(node.OpenCurlyBracketTkn, nil)
	printer.printNode(node.Name)
	printer.printToken(node.CloseCurlyBracketTkn, nil)
}

func (printer *printer) ExprYield(node *ast.ExprYield) {
	printer.printToken(node.YieldTkn, []byte("yield"))
	printer.printNode(node.Key)
	printer.printToken(node.DoubleArrowTkn, printer.ifNode(node.Key, []byte("=>")))
	printer.printNode(node.Val)
}

func (printer *printer) ExprYieldFrom(node *ast.ExprYieldFrom) {
	printer.printToken(node.YieldFromTkn, []byte("yield from"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssign(node *ast.ExprAssign) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignReference(node *ast.ExprAssignReference) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("="))
	printer.printToken(node.AmpersandTkn, []byte("&"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignBitwiseAnd(node *ast.ExprAssignBitwiseAnd) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("&="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignBitwiseOr(node *ast.ExprAssignBitwiseOr) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("|="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignBitwiseXor(node *ast.ExprAssignBitwiseXor) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("^="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignCoalesce(node *ast.ExprAssignCoalesce) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("??="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignConcat(node *ast.ExprAssignConcat) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte(".="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignDiv(node *ast.ExprAssignDiv) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("/="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignMinus(node *ast.ExprAssignMinus) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("-="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignMod(node *ast.ExprAssignMod) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("%="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignMul(node *ast.ExprAssignMul) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("*="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignPlus(node *ast.ExprAssignPlus) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("+="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignPow(node *ast.ExprAssignPow) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("**="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignShiftLeft(node *ast.ExprAssignShiftLeft) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte("<<="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprAssignShiftRight(node *ast.ExprAssignShiftRight) {
	printer.printNode(node.Var)
	printer.printToken(node.EqualTkn, []byte(">>="))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprBinaryBitwiseAnd(node *ast.ExprBinaryBitwiseAnd) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("&"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryBitwiseOr(node *ast.ExprBinaryBitwiseOr) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("|"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryBitwiseXor(node *ast.ExprBinaryBitwiseXor) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("^"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryBooleanAnd(node *ast.ExprBinaryBooleanAnd) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("&&"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryBooleanOr(node *ast.ExprBinaryBooleanOr) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("||"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryCoalesce(node *ast.ExprBinaryCoalesce) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("??"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryConcat(node *ast.ExprBinaryConcat) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("."))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryDiv(node *ast.ExprBinaryDiv) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("/"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryEqual(node *ast.ExprBinaryEqual) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("=="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryGreater(node *ast.ExprBinaryGreater) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte(">"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryGreaterOrEqual(node *ast.ExprBinaryGreaterOrEqual) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte(">="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryIdentical(node *ast.ExprBinaryIdentical) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("==="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryLogicalAnd(node *ast.ExprBinaryLogicalAnd) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("and"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryLogicalOr(node *ast.ExprBinaryLogicalOr) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("or"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryLogicalXor(node *ast.ExprBinaryLogicalXor) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("xor"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryMinus(node *ast.ExprBinaryMinus) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("-"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryMod(node *ast.ExprBinaryMod) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("%"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryMul(node *ast.ExprBinaryMul) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("*"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryNotEqual(node *ast.ExprBinaryNotEqual) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("!="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryNotIdentical(node *ast.ExprBinaryNotIdentical) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("!=="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryPlus(node *ast.ExprBinaryPlus) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("+"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryPow(node *ast.ExprBinaryPow) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("**"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryShiftLeft(node *ast.ExprBinaryShiftLeft) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("<<"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinaryShiftRight(node *ast.ExprBinaryShiftRight) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte(">>"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinarySmaller(node *ast.ExprBinarySmaller) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("<"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinarySmallerOrEqual(node *ast.ExprBinarySmallerOrEqual) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("<="))
	printer.printNode(node.Right)
}

func (printer *printer) ExprBinarySpaceship(node *ast.ExprBinarySpaceship) {
	printer.printNode(node.Left)
	printer.printToken(node.OpTkn, []byte("<=>"))
	printer.printNode(node.Right)
}

func (printer *printer) ExprCastArray(node *ast.ExprCastArray) {
	printer.printToken(node.CastTkn, []byte("(array)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastBool(node *ast.ExprCastBool) {
	printer.printToken(node.CastTkn, []byte("(bool)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastDouble(node *ast.ExprCastDouble) {
	printer.printToken(node.CastTkn, []byte("(float)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastInt(node *ast.ExprCastInt) {
	printer.printToken(node.CastTkn, []byte("(int)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastObject(node *ast.ExprCastObject) {
	printer.printToken(node.CastTkn, []byte("(object)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastString(node *ast.ExprCastString) {
	printer.printToken(node.CastTkn, []byte("(string)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ExprCastUnset(node *ast.ExprCastUnset) {
	printer.printToken(node.CastTkn, []byte("(unset)"))
	printer.printNode(node.Expr)
}

func (printer *printer) ScalarDnumber(node *ast.ScalarDnumber) {
	printer.printToken(node.NumberTkn, node.Value)
}

func (printer *printer) ScalarEncapsed(node *ast.ScalarEncapsed) {
	printer.printToken(node.OpenQuoteTkn, []byte("\""))
	printer.printList(node.Parts)
	printer.printToken(node.CloseQuoteTkn, []byte("\""))
}

func (printer *printer) ScalarEncapsedStringPart(node *ast.ScalarEncapsedStringPart) {
	printer.printToken(node.EncapsedStrTkn, node.Value)
}

func (printer *printer) ScalarEncapsedStringVar(node *ast.ScalarEncapsedStringVar) {
	printer.printToken(node.DollarOpenCurlyBracketTkn, []byte("${"))
	printer.printNode(node.Name)
	printer.printToken(node.OpenSquareBracketTkn, printer.ifNode(node.Dim, []byte("[")))
	printer.printNode(node.Dim)
	printer.printToken(node.CloseSquareBracketTkn, printer.ifNode(node.Dim, []byte("]")))
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) ScalarEncapsedStringBrackets(node *ast.ScalarEncapsedStringBrackets) {
	printer.printToken(node.OpenCurlyBracketTkn, []byte("{"))
	printer.printNode(node.Var)
	printer.printToken(node.CloseCurlyBracketTkn, []byte("}"))
}

func (printer *printer) ScalarHeredoc(node *ast.ScalarHeredoc) {
	printer.printToken(node.OpenHeredocTkn, []byte("<<<EOT\n"))
	printer.printList(node.Parts)
	printer.printToken(node.CloseHeredocTkn, []byte("EOT"))
}

func (printer *printer) ScalarLnumber(node *ast.ScalarLnumber) {
	printer.printToken(node.NumberTkn, node.Value)
}

func (printer *printer) ScalarMagicConstant(node *ast.ScalarMagicConstant) {
	printer.printToken(node.MagicConstTkn, node.Value)
}

func (printer *printer) ScalarString(node *ast.ScalarString) {
	printer.printToken(node.MinusTkn, nil)
	printer.printToken(node.StringTkn, node.Value)
}

func (printer *printer) NameName(node *ast.Name) {
	printer.printSeparatedList(node.Parts, node.SeparatorTkns, []byte("\\"))
}

func (printer *printer) NameFullyQualified(node *ast.NameFullyQualified) {
	printer.printToken(node.NsSeparatorTkn, []byte("\\"))
	printer.printSeparatedList(node.Parts, node.SeparatorTkns, []byte("\\"))
}

func (printer *printer) NameRelative(node *ast.NameRelative) {
	printer.printToken(node.NsTkn, []byte("namespace"))
	printer.printToken(node.NsSeparatorTkn, []byte("\\"))
	printer.printSeparatedList(node.Parts, node.SeparatorTkns, []byte("\\"))
}

func (printer *printer) NameNamePart(node *ast.NamePart) {
	printer.printToken(node.StringTkn, node.Value)
}
