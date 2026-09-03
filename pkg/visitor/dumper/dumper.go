package dumper

import (
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"io"
	"strconv"
	"strings"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

type Dumper struct {
	writer        io.Writer
	indent        int
	withTokens    bool
	withPositions bool
}

func NewDumper(writer io.Writer) *Dumper {
	return &Dumper{writer: writer}
}

func (dumper *Dumper) WithTokens() *Dumper {
	dumper.withTokens = true
	return dumper
}

func (dumper *Dumper) WithPositions() *Dumper {
	dumper.withPositions = true
	return dumper
}

func (dumper *Dumper) Dump(node ast.Vertex) {
	node.Accept(dumper)
}

func (dumper *Dumper) print(indent int, text string) {
	_, err := io.WriteString(dumper.writer, strings.Repeat("\t", indent))
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(dumper.writer, text)
	if err != nil {
		panic(err)
	}
}

func (dumper *Dumper) dumpVertex(key string, node ast.Vertex) {
	if node == nil {
		return
	}

	dumper.print(dumper.indent, key+": ")
	node.Accept(dumper)
}

func (dumper *Dumper) dumpVertexList(key string, list []ast.Vertex) {
	if list == nil {
		return
	}

	if len(list) == 0 {
		dumper.print(dumper.indent, key+": []ast.Vertex{},\n")
		return
	}

	dumper.print(dumper.indent, key+": []ast.Vertex{\n")
	dumper.indent++

	for _, node := range list {
		dumper.print(dumper.indent, "")
		node.Accept(dumper)
	}

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) dumpToken(key string, tokenItem *token.Token) {
	if !dumper.withTokens {
		return
	}

	if tokenItem == nil {
		return
	}

	if key == "" {
		dumper.print(dumper.indent, "{\n")
	} else {
		dumper.print(dumper.indent, key+": &token.Token{\n")
	}

	dumper.indent++

	if tokenItem.ID > 0 {
		dumper.print(dumper.indent, "ID: token."+tokenItem.ID.String()+",\n")
	}
	if tokenItem.Value != nil {
		dumper.print(dumper.indent, "Value: []byte("+strconv.Quote(string(tokenItem.Value))+"),\n")
	}
	dumper.dumpPosition(tokenItem.Position)
	dumper.dumpTokenList("FreeFloating", tokenItem.FreeFloating)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) dumpTokenList(key string, list []*token.Token) {
	if !dumper.withTokens {
		return
	}

	if list == nil {
		return
	}

	if len(list) == 0 {
		dumper.print(dumper.indent, key+": []*token.Token{},\n")
		return
	}

	dumper.print(dumper.indent, key+": []*token.Token{\n")
	dumper.indent++

	for _, tokenItem := range list {
		dumper.dumpToken("", tokenItem)
	}

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) dumpPosition(positionItem *position.Position) {
	if !dumper.withPositions {
		return
	}

	if positionItem == nil {
		return
	}

	dumper.print(dumper.indent, "Position: &position.Position{\n")
	dumper.indent++

	dumper.print(dumper.indent, "StartLine: "+strconv.Itoa(positionItem.StartLine)+",\n")
	dumper.print(dumper.indent, "EndLine:   "+strconv.Itoa(positionItem.EndLine)+",\n")
	dumper.print(dumper.indent, "StartPos:  "+strconv.Itoa(positionItem.StartPos)+",\n")
	dumper.print(dumper.indent, "EndPos:    "+strconv.Itoa(positionItem.EndPos)+",\n")

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) dumpValue(key string, value []byte) {
	if value == nil {
		return
	}

	dumper.print(dumper.indent, key+": []byte("+strconv.Quote(string(value))+"),\n")

}

func (dumper *Dumper) Root(node *ast.Root) {
	dumper.print(0, "&ast.Root{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("EndTkn", node.EndTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Nullable(node *ast.Nullable) {
	dumper.print(0, "&ast.Nullable{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("QuestionTkn", node.QuestionTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Union(node *ast.Union) {
	dumper.print(0, "&ast.Union{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("Types", node.Types)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Intersection(node *ast.Intersection) {
	dumper.print(0, "&ast.Intersection{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("Types", node.Types)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Parameter(node *ast.Parameter) {
	dumper.print(0, "&ast.Parameter{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpVertexList("Modifiers", node.Modifiers)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpToken("VariadicTkn", node.VariadicTkn)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("DefaultValue", node.DefaultValue)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Identifier(node *ast.Identifier) {
	dumper.print(0, "&ast.Identifier{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IdentifierTkn", node.IdentifierTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Argument(node *ast.Argument) {
	dumper.print(0, "&ast.Argument{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpToken("VariadicTkn", node.VariadicTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) Attribute(node *ast.Attribute) {
	dumper.print(0, "&ast.Attribute{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) AttributeGroup(node *ast.AttributeGroup) {
	dumper.print(0, "&ast.AttributeGroup{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenAttributeTkn", node.OpenAttributeTkn)
	dumper.dumpVertexList("Attrs", node.Attrs)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseAttributeTkn", node.CloseAttributeTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtBreak(node *ast.StmtBreak) {
	dumper.print(0, "&ast.StmtBreak{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("BreakTkn", node.BreakTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtCase(node *ast.StmtCase) {
	dumper.print(0, "&ast.StmtCase{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CaseTkn", node.CaseTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CaseSeparatorTkn", node.CaseSeparatorTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtCatch(node *ast.StmtCatch) {
	dumper.print(0, "&ast.StmtCatch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CatchTkn", node.CatchTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Types", node.Types)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtClass(node *ast.StmtClass) {
	dumper.print(0, "&ast.StmtClass{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpVertexList("Modifiers", node.Modifiers)
	dumper.dumpToken("ClassTkn", node.ClassTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ExtendsTkn", node.ExtendsTkn)
	dumper.dumpVertex("Extends", node.Extends)
	dumper.dumpToken("ImplementsTkn", node.ImplementsTkn)
	dumper.dumpVertexList("Implements", node.Implements)
	dumper.dumpTokenList("ImplementsSeparatorTkns", node.ImplementsSeparatorTkns)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtClassConstList(node *ast.StmtClassConstList) {
	dumper.print(0, "&ast.StmtClassConstList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpVertexList("Modifiers", node.Modifiers)
	dumper.dumpToken("ConstTkn", node.ConstTkn)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpVertexList("Consts", node.Consts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtClassMethod(node *ast.StmtClassMethod) {
	dumper.print(0, "&ast.StmtClassMethod{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpVertexList("Modifiers", node.Modifiers)
	dumper.dumpToken("FunctionTkn", node.FunctionTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Params", node.Params)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("ReturnType", node.ReturnType)
	dumper.dumpVertex("Stmt", node.Stmt)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtConstList(node *ast.StmtConstList) {
	dumper.print(0, "&ast.StmtConstList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ConstTkn", node.ConstTkn)
	dumper.dumpVertexList("Consts", node.Consts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtConstant(node *ast.StmtConstant) {
	dumper.print(0, "&ast.StmtConstant{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtContinue(node *ast.StmtContinue) {
	dumper.print(0, "&ast.StmtContinue{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ContinueTkn", node.ContinueTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtDeclare(node *ast.StmtDeclare) {
	dumper.print(0, "&ast.StmtDeclare{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DeclareTkn", node.DeclareTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Consts", node.Consts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpToken("EndDeclareTkn", node.EndDeclareTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtDefault(node *ast.StmtDefault) {
	dumper.print(0, "&ast.StmtDefault{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DefaultTkn", node.DefaultTkn)
	dumper.dumpToken("CaseSeparatorTkn", node.CaseSeparatorTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtDo(node *ast.StmtDo) {
	dumper.print(0, "&ast.StmtDo{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DoTkn", node.DoTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpToken("WhileTkn", node.WhileTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")

}

func (dumper *Dumper) StmtEcho(node *ast.StmtEcho) {
	dumper.print(0, "&ast.StmtEcho{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("EchoTkn", node.EchoTkn)
	dumper.dumpVertexList("Exprs", node.Exprs)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtElse(node *ast.StmtElse) {
	dumper.print(0, "&ast.StmtElse{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ElseTkn", node.ElseTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtElseIf(node *ast.StmtElseIf) {
	dumper.print(0, "&ast.StmtElseIf{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ElseIfTkn", node.ElseIfTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtExpression(node *ast.StmtExpression) {
	dumper.print(0, "&ast.StmtExpression{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtFinally(node *ast.StmtFinally) {
	dumper.print(0, "&ast.StmtFinally{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("FinallyTkn", node.FinallyTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtFor(node *ast.StmtFor) {
	dumper.print(0, "&ast.StmtFor{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ForTkn", node.ForTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Init", node.Init)
	dumper.dumpTokenList("InitSeparatorTkns", node.InitSeparatorTkns)
	dumper.dumpToken("InitSemiColonTkn", node.InitSemiColonTkn)
	dumper.dumpVertexList("Cond", node.Cond)
	dumper.dumpTokenList("CondSeparatorTkns", node.CondSeparatorTkns)
	dumper.dumpToken("CondSemiColonTkn", node.CondSemiColonTkn)
	dumper.dumpVertexList("Loop", node.Loop)
	dumper.dumpTokenList("LoopSeparatorTkns", node.LoopSeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpToken("EndForTkn", node.EndForTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtForeach(node *ast.StmtForeach) {
	dumper.print(0, "&ast.StmtForeach{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ForeachTkn", node.ForeachTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("AsTkn", node.AsTkn)
	dumper.dumpVertex("Key", node.Key)
	dumper.dumpToken("DoubleArrowTkn", node.DoubleArrowTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpToken("EndForeachTkn", node.EndForeachTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtFunction(node *ast.StmtFunction) {
	dumper.print(0, "&ast.StmtFunction{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpToken("FunctionTkn", node.FunctionTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Params", node.Params)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("ReturnType", node.ReturnType)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtGlobal(node *ast.StmtGlobal) {
	dumper.print(0, "&ast.StmtGlobal{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("GlobalTkn", node.GlobalTkn)
	dumper.dumpVertexList("Vars", node.Vars)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtGoto(node *ast.StmtGoto) {
	dumper.print(0, "&ast.StmtGoto{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("GotoTkn", node.GotoTkn)
	dumper.dumpVertex("Label", node.Label)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtHaltCompiler(node *ast.StmtHaltCompiler) {
	dumper.print(0, "&ast.StmtHaltCompiler{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("HaltCompilerTkn", node.HaltCompilerTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtIf(node *ast.StmtIf) {
	dumper.print(0, "&ast.StmtIf{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IfTkn", node.IfTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpVertexList("ElseIf", node.ElseIf)
	dumper.dumpVertex("Else", node.Else)
	dumper.dumpToken("EndIfTkn", node.EndIfTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtInlineHtml(node *ast.StmtInlineHtml) {
	dumper.print(0, "&ast.StmtInlineHtml{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("InlineHtmlTkn", node.InlineHtmlTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtEnum(node *ast.StmtEnum) {
	dumper.print(0, "&ast.StmtEnum{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpToken("EnumTkn", node.EnumTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpToken("ImplementsTkn", node.ImplementsTkn)
	dumper.dumpVertexList("Implements", node.Implements)
	dumper.dumpTokenList("ImplementsSeparatorTkns", node.ImplementsSeparatorTkns)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtEnumCase(node *ast.StmtEnumCase) {
	dumper.print(0, "&ast.StmtEnumCase{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpToken("CaseTkn", node.CaseTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtInterface(node *ast.StmtInterface) {
	dumper.print(0, "&ast.StmtInterface{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpToken("InterfaceTkn", node.InterfaceTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("ExtendsTkn", node.ExtendsTkn)
	dumper.dumpVertexList("Extends", node.Extends)
	dumper.dumpTokenList("ExtendsSeparatorTkns", node.ExtendsSeparatorTkns)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtLabel(node *ast.StmtLabel) {
	dumper.print(0, "&ast.StmtLabel{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("ColonTkn", node.ColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtNamespace(node *ast.StmtNamespace) {
	dumper.print(0, "&ast.StmtNamespace{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NsTkn", node.NsTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtNop(node *ast.StmtNop) {
	dumper.print(0, "&ast.StmtNop{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtProperty(node *ast.StmtProperty) {
	dumper.print(0, "&ast.StmtProperty{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtPropertyList(node *ast.StmtPropertyList) {
	dumper.print(0, "&ast.StmtPropertyList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpVertexList("Modifiers", node.Modifiers)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpVertexList("Props", node.Props)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtReturn(node *ast.StmtReturn) {
	dumper.print(0, "&ast.StmtReturn{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ReturnTkn", node.ReturnTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtStatic(node *ast.StmtStatic) {
	dumper.print(0, "&ast.StmtStatic{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("StaticTkn", node.StaticTkn)
	dumper.dumpVertexList("Vars", node.Vars)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtStaticVar(node *ast.StmtStaticVar) {
	dumper.print(0, "&ast.StmtStaticVar{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtStmtList(node *ast.StmtStmtList) {
	dumper.print(0, "&ast.StmtStmtList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtSwitch(node *ast.StmtSwitch) {
	dumper.print(0, "&ast.StmtSwitch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("SwitchTkn", node.SwitchTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpToken("CaseSeparatorTkn", node.CaseSeparatorTkn)
	dumper.dumpVertexList("Cases", node.Cases)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("EndSwitchTkn", node.EndSwitchTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtThrow(node *ast.StmtThrow) {
	dumper.print(0, "&ast.StmtThrow{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ThrowTkn", node.ThrowTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtTrait(node *ast.StmtTrait) {
	dumper.print(0, "&ast.StmtTrait{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("AttrGroups", node.AttrGroups)
	dumper.dumpToken("TraitTkn", node.TraitTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtTraitUse(node *ast.StmtTraitUse) {
	dumper.print(0, "&ast.StmtTraitUse{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("UseTkn", node.UseTkn)
	dumper.dumpVertexList("Traits", node.Traits)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Adaptations", node.Adaptations)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtTraitUseAlias(node *ast.StmtTraitUseAlias) {
	dumper.print(0, "&ast.StmtTraitUseAlias{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Trait", node.Trait)
	dumper.dumpToken("DoubleColonTkn", node.DoubleColonTkn)
	dumper.dumpVertex("Method", node.Method)
	dumper.dumpToken("AsTkn", node.AsTkn)
	dumper.dumpVertex("Modifier", node.Modifier)
	dumper.dumpVertex("Alias", node.Alias)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtTraitUsePrecedence(node *ast.StmtTraitUsePrecedence) {
	dumper.print(0, "&ast.StmtTraitUsePrecedence{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Trait", node.Trait)
	dumper.dumpToken("DoubleColonTkn", node.DoubleColonTkn)
	dumper.dumpVertex("Method", node.Method)
	dumper.dumpToken("InsteadofTkn", node.InsteadofTkn)
	dumper.dumpVertexList("Insteadof", node.Insteadof)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtTry(node *ast.StmtTry) {
	dumper.print(0, "&ast.StmtTry{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("TryTkn", node.TryTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpVertexList("Catches", node.Catches)
	dumper.dumpVertex("Finally", node.Finally)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtUnset(node *ast.StmtUnset) {
	dumper.print(0, "&ast.StmtUnset{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("UnsetTkn", node.UnsetTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Vars", node.Vars)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtUse(node *ast.StmtUseList) {
	dumper.print(0, "&ast.StmtUseList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("UseTkn", node.UseTkn)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpVertexList("Uses", node.Uses)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtGroupUse(node *ast.StmtGroupUseList) {
	dumper.print(0, "&ast.StmtGroupUseList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("UseTkn", node.UseTkn)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpToken("LeadingNsSeparatorTkn", node.LeadingNsSeparatorTkn)
	dumper.dumpVertex("Prefix", node.Prefix)
	dumper.dumpToken("NsSeparatorTkn", node.NsSeparatorTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Uses", node.Uses)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtUseDeclaration(node *ast.StmtUse) {
	dumper.print(0, "&ast.StmtUse{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Type", node.Type)
	dumper.dumpToken("NsSeparatorTkn", node.NsSeparatorTkn)
	dumper.dumpVertex("Uses", node.Use)
	dumper.dumpToken("AsTkn", node.AsTkn)
	dumper.dumpVertex("Alias", node.Alias)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) StmtWhile(node *ast.StmtWhile) {
	dumper.print(0, "&ast.StmtWhile{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("WhileTkn", node.WhileTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("Stmt", node.Stmt)
	dumper.dumpToken("EndWhileTkn", node.EndWhileTkn)
	dumper.dumpToken("SemiColonTkn", node.SemiColonTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprArray(node *ast.ExprArray) {
	dumper.print(0, "&ast.ExprArray{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ArrayTkn", node.ArrayTkn)
	dumper.dumpToken("OpenBracketTkn", node.OpenBracketTkn)
	dumper.dumpVertexList("Items", node.Items)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseBracketTkn", node.CloseBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprArrayDimFetch(node *ast.ExprArrayDimFetch) {
	dumper.print(0, "&ast.ExprArrayDimFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("OpenBracketTkn", node.OpenBracketTkn)
	dumper.dumpVertex("Dim", node.Dim)
	dumper.dumpToken("CloseBracketTkn", node.CloseBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprArrayItem(node *ast.ExprArrayItem) {
	dumper.print(0, "&ast.ExprArrayItem{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("EllipsisTkn", node.EllipsisTkn)
	dumper.dumpVertex("Key", node.Key)
	dumper.dumpToken("DoubleArrowTkn", node.DoubleArrowTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Val", node.Val)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprArrowFunction(node *ast.ExprArrowFunction) {
	dumper.print(0, "&ast.ExprArrowFunction{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("StaticTkn", node.StaticTkn)
	dumper.dumpToken("FnTkn", node.FnTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Params", node.Params)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("ReturnType", node.ReturnType)
	dumper.dumpToken("DoubleArrowTkn", node.DoubleArrowTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBitwiseNot(node *ast.ExprBitwiseNot) {
	dumper.print(0, "&ast.ExprBitwiseNot{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("TildaTkn", node.TildaTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBooleanNot(node *ast.ExprBooleanNot) {
	dumper.print(0, "&ast.ExprBooleanNot{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ExclamationTkn", node.ExclamationTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBrackets(node *ast.ExprBrackets) {
	dumper.print(0, "&ast.ExprBrackets{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprClassConstFetch(node *ast.ExprClassConstFetch) {
	dumper.print(0, "&ast.ExprClassConstFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Class", node.Class)
	dumper.dumpToken("DoubleColonTkn", node.DoubleColonTkn)
	dumper.dumpVertex("Const", node.Const)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprClone(node *ast.ExprClone) {
	dumper.print(0, "&ast.ExprClone{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CloneTkn", node.CloneTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprThrow(node *ast.ExprThrow) {
	dumper.print(0, "&ast.ExprThrow{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ThrowTkn", node.ThrowTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprClosure(node *ast.ExprClosure) {
	dumper.print(0, "&ast.ExprClosure{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("StaticTkn", node.StaticTkn)
	dumper.dumpToken("FunctionTkn", node.FunctionTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Params", node.Params)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("UseTkn", node.UseTkn)
	dumper.dumpToken("UseOpenParenthesisTkn", node.UseOpenParenthesisTkn)
	dumper.dumpVertexList("Uses", node.Uses)
	dumper.dumpTokenList("UseSeparatorTkns", node.UseSeparatorTkns)
	dumper.dumpToken("UseCloseParenthesisTkn", node.UseCloseParenthesisTkn)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("ReturnType", node.ReturnType)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Stmts", node.Stmts)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprClosureUse(node *ast.ExprClosureUse) {
	dumper.print(0, "&ast.ExprClosureUse{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Var", node.Var)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprConstFetch(node *ast.ExprConstFetch) {
	dumper.print(0, "&ast.ExprConstFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Const", node.Const)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprEmpty(node *ast.ExprEmpty) {
	dumper.print(0, "&ast.ExprEmpty{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("EmptyTkn", node.EmptyTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprErrorSuppress(node *ast.ExprErrorSuppress) {
	dumper.print(0, "&ast.ExprErrorSuppress{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("AtTkn", node.AtTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprEval(node *ast.ExprEval) {
	dumper.print(0, "&ast.ExprEval{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("EvalTkn", node.EvalTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprExit(node *ast.ExprExit) {
	dumper.print(0, "&ast.ExprExit{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ExitTkn", node.ExitTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprFunctionCall(node *ast.ExprFunctionCall) {
	dumper.print(0, "&ast.ExprFunctionCall{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Function", node.Function)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprInclude(node *ast.ExprInclude) {
	dumper.print(0, "&ast.ExprInclude{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IncludeOnceTkn", node.IncludeTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprIncludeOnce(node *ast.ExprIncludeOnce) {
	dumper.print(0, "&ast.ExprIncludeOnce{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IncludeOnceTkn", node.IncludeOnceTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprInstanceOf(node *ast.ExprInstanceOf) {
	dumper.print(0, "&ast.ExprInstanceOf{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("InstanceOfTkn", node.InstanceOfTkn)
	dumper.dumpVertex("Class", node.Class)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprIsset(node *ast.ExprIsset) {
	dumper.print(0, "&ast.ExprIsset{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IssetTkn", node.IssetTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Vars", node.Vars)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprList(node *ast.ExprList) {
	dumper.print(0, "&ast.ExprList{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("ListTkn", node.ListTkn)
	dumper.dumpToken("OpenBracketTkn", node.OpenBracketTkn)
	dumper.dumpVertexList("Items", node.Items)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseBracketTkn", node.CloseBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprMethodCall(node *ast.ExprMethodCall) {
	dumper.print(0, "&ast.ExprMethodCall{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("ObjectOperatorTkn", node.ObjectOperatorTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Method", node.Method)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprNullsafeMethodCall(node *ast.ExprNullsafeMethodCall) {
	dumper.print(0, "&ast.ExprNullsafeMethodCall{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("ObjectOperatorTkn", node.ObjectOperatorTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Method", node.Method)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprNullsafePropertyFetch(node *ast.ExprNullsafePropertyFetch) {
	dumper.print(0, "&ast.ExprNullsafePropertyFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("ObjectOperatorTkn", node.ObjectOperatorTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Prop", node.Prop)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprNew(node *ast.ExprNew) {
	dumper.print(0, "&ast.ExprNew{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NewTkn", node.NewTkn)
	dumper.dumpVertex("Class", node.Class)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPostDec(node *ast.ExprPostDec) {
	dumper.print(0, "&ast.ExprPostDec{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("DecTkn", node.DecTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPostInc(node *ast.ExprPostInc) {
	dumper.print(0, "&ast.ExprPostInc{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("IncTkn", node.IncTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPreDec(node *ast.ExprPreDec) {
	dumper.print(0, "&ast.ExprPreDec{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DecTkn", node.DecTkn)
	dumper.dumpVertex("Var", node.Var)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPreInc(node *ast.ExprPreInc) {
	dumper.print(0, "&ast.ExprPreInc{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("IncTkn", node.IncTkn)
	dumper.dumpVertex("Var", node.Var)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPrint(node *ast.ExprPrint) {
	dumper.print(0, "&ast.ExprPrint{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("PrintTkn", node.PrintTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprPropertyFetch(node *ast.ExprPropertyFetch) {
	dumper.print(0, "&ast.ExprPropertyFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("ObjectOperatorTkn", node.ObjectOperatorTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Prop", node.Prop)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprRequire(node *ast.ExprRequire) {
	dumper.print(0, "&ast.ExprRequire{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("RequireTkn", node.RequireTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprRequireOnce(node *ast.ExprRequireOnce) {
	dumper.print(0, "&ast.ExprRequireOnce{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("RequireOnceTkn", node.RequireOnceTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprShellExec(node *ast.ExprShellExec) {
	dumper.print(0, "&ast.ExprShellExec{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenBacktickTkn", node.OpenBacktickTkn)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpToken("CloseBacktickTkn", node.CloseBacktickTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprStaticCall(node *ast.ExprStaticCall) {
	dumper.print(0, "&ast.ExprStaticCall{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Class", node.Class)
	dumper.dumpToken("DoubleColonTkn", node.DoubleColonTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Call", node.Call)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertexList("Args", node.Args)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprStaticPropertyFetch(node *ast.ExprStaticPropertyFetch) {
	dumper.print(0, "&ast.ExprStaticPropertyFetch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Class", node.Class)
	dumper.dumpToken("DoubleColonTkn", node.DoubleColonTkn)
	dumper.dumpVertex("Prop", node.Prop)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprTernary(node *ast.ExprTernary) {
	dumper.print(0, "&ast.ExprTernary{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Cond", node.Cond)
	dumper.dumpToken("QuestionTkn", node.QuestionTkn)
	dumper.dumpVertex("IfTrue", node.IfTrue)
	dumper.dumpToken("ColonTkn", node.ColonTkn)
	dumper.dumpVertex("IfFalse", node.IfFalse)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprMatch(node *ast.ExprMatch) {
	dumper.print(0, "&ast.ExprMatch{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("MatchTkn", node.MatchTkn)
	dumper.dumpToken("OpenParenthesisTkn", node.OpenParenthesisTkn)
	dumper.dumpVertex("Expr", node.Expr)
	dumper.dumpToken("CloseParenthesisTkn", node.CloseParenthesisTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertexList("Arms", node.Arms)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) MatchArm(node *ast.MatchArm) {
	dumper.print(0, "&ast.MatchArm{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DefaultTkn", node.DefaultTkn)
	dumper.dumpVertexList("Exprs", node.Exprs)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)
	dumper.dumpToken("DoubleArrowTkn", node.DoubleArrowTkn)
	dumper.dumpVertex("ReturnExpr", node.ReturnExpr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprUnaryMinus(node *ast.ExprUnaryMinus) {
	dumper.print(0, "&ast.ExprUnaryMinus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("MinusTkn", node.MinusTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprUnaryPlus(node *ast.ExprUnaryPlus) {
	dumper.print(0, "&ast.ExprUnaryPlus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("PlusTkn", node.PlusTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprVariable(node *ast.ExprVariable) {
	dumper.print(0, "&ast.ExprVariable{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DollarTkn", node.DollarTkn)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprYield(node *ast.ExprYield) {
	dumper.print(0, "&ast.ExprYield{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("YieldTkn", node.YieldTkn)
	dumper.dumpVertex("Key", node.Key)
	dumper.dumpToken("DoubleArrowTkn", node.DoubleArrowTkn)
	dumper.dumpVertex("Val", node.Val)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprYieldFrom(node *ast.ExprYieldFrom) {
	dumper.print(0, "&ast.ExprYieldFrom{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("YieldFromTkn", node.YieldFromTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssign(node *ast.ExprAssign) {
	dumper.print(0, "&ast.ExprAssign{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignReference(node *ast.ExprAssignReference) {
	dumper.print(0, "&ast.ExprAssignReference{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpToken("AmpersandTkn", node.AmpersandTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignBitwiseAnd(node *ast.ExprAssignBitwiseAnd) {
	dumper.print(0, "&ast.ExprAssignBitwiseAnd{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignBitwiseOr(node *ast.ExprAssignBitwiseOr) {
	dumper.print(0, "&ast.ExprAssignBitwiseOr{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignBitwiseXor(node *ast.ExprAssignBitwiseXor) {
	dumper.print(0, "&ast.ExprAssignBitwiseXor{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignCoalesce(node *ast.ExprAssignCoalesce) {
	dumper.print(0, "&ast.ExprAssignCoalesce{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignConcat(node *ast.ExprAssignConcat) {
	dumper.print(0, "&ast.ExprAssignConcat{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignDiv(node *ast.ExprAssignDiv) {
	dumper.print(0, "&ast.ExprAssignDiv{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignMinus(node *ast.ExprAssignMinus) {
	dumper.print(0, "&ast.ExprAssignMinus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignMod(node *ast.ExprAssignMod) {
	dumper.print(0, "&ast.ExprAssignMod{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignMul(node *ast.ExprAssignMul) {
	dumper.print(0, "&ast.ExprAssignMul{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignPlus(node *ast.ExprAssignPlus) {
	dumper.print(0, "&ast.ExprAssignPlus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignPow(node *ast.ExprAssignPow) {
	dumper.print(0, "&ast.ExprAssignPow{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignShiftLeft(node *ast.ExprAssignShiftLeft) {
	dumper.print(0, "&ast.ExprAssignShiftLeft{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprAssignShiftRight(node *ast.ExprAssignShiftRight) {
	dumper.print(0, "&ast.ExprAssignShiftRight{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("EqualTkn", node.EqualTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryBitwiseAnd(node *ast.ExprBinaryBitwiseAnd) {
	dumper.print(0, "&ast.ExprBinaryBitwiseAnd{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryBitwiseOr(node *ast.ExprBinaryBitwiseOr) {
	dumper.print(0, "&ast.ExprBinaryBitwiseOr{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryBitwiseXor(node *ast.ExprBinaryBitwiseXor) {
	dumper.print(0, "&ast.ExprBinaryBitwiseXor{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryBooleanAnd(node *ast.ExprBinaryBooleanAnd) {
	dumper.print(0, "&ast.ExprBinaryBooleanAnd{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryBooleanOr(node *ast.ExprBinaryBooleanOr) {
	dumper.print(0, "&ast.ExprBinaryBooleanOr{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryCoalesce(node *ast.ExprBinaryCoalesce) {
	dumper.print(0, "&ast.ExprBinaryCoalesce{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryConcat(node *ast.ExprBinaryConcat) {
	dumper.print(0, "&ast.ExprBinaryConcat{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryDiv(node *ast.ExprBinaryDiv) {
	dumper.print(0, "&ast.ExprBinaryDiv{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryEqual(node *ast.ExprBinaryEqual) {
	dumper.print(0, "&ast.ExprBinaryEqual{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryGreater(node *ast.ExprBinaryGreater) {
	dumper.print(0, "&ast.ExprBinaryGreater{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryGreaterOrEqual(node *ast.ExprBinaryGreaterOrEqual) {
	dumper.print(0, "&ast.ExprBinaryGreaterOrEqual{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryIdentical(node *ast.ExprBinaryIdentical) {
	dumper.print(0, "&ast.ExprBinaryIdentical{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryLogicalAnd(node *ast.ExprBinaryLogicalAnd) {
	dumper.print(0, "&ast.ExprBinaryLogicalAnd{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryLogicalOr(node *ast.ExprBinaryLogicalOr) {
	dumper.print(0, "&ast.ExprBinaryLogicalOr{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryLogicalXor(node *ast.ExprBinaryLogicalXor) {
	dumper.print(0, "&ast.ExprBinaryLogicalXor{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryMinus(node *ast.ExprBinaryMinus) {
	dumper.print(0, "&ast.ExprBinaryMinus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryMod(node *ast.ExprBinaryMod) {
	dumper.print(0, "&ast.ExprBinaryMod{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryMul(node *ast.ExprBinaryMul) {
	dumper.print(0, "&ast.ExprBinaryMul{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryNotEqual(node *ast.ExprBinaryNotEqual) {
	dumper.print(0, "&ast.ExprBinaryNotEqual{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryNotIdentical(node *ast.ExprBinaryNotIdentical) {
	dumper.print(0, "&ast.ExprBinaryNotIdentical{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryPlus(node *ast.ExprBinaryPlus) {
	dumper.print(0, "&ast.ExprBinaryPlus{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryPow(node *ast.ExprBinaryPow) {
	dumper.print(0, "&ast.ExprBinaryPow{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryShiftLeft(node *ast.ExprBinaryShiftLeft) {
	dumper.print(0, "&ast.ExprBinaryShiftLeft{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinaryShiftRight(node *ast.ExprBinaryShiftRight) {
	dumper.print(0, "&ast.ExprBinaryShiftRight{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinarySmaller(node *ast.ExprBinarySmaller) {
	dumper.print(0, "&ast.ExprBinarySmaller{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinarySmallerOrEqual(node *ast.ExprBinarySmallerOrEqual) {
	dumper.print(0, "&ast.ExprBinarySmallerOrEqual{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprBinarySpaceship(node *ast.ExprBinarySpaceship) {
	dumper.print(0, "&ast.ExprBinarySpaceship{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertex("Left", node.Left)
	dumper.dumpToken("OpTkn", node.OpTkn)
	dumper.dumpVertex("Right", node.Right)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastArray(node *ast.ExprCastArray) {
	dumper.print(0, "&ast.ExprCastArray{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastBool(node *ast.ExprCastBool) {
	dumper.print(0, "&ast.ExprCastBool{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastDouble(node *ast.ExprCastDouble) {
	dumper.print(0, "&ast.ExprCastDouble{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastInt(node *ast.ExprCastInt) {
	dumper.print(0, "&ast.ExprCastInt{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastObject(node *ast.ExprCastObject) {
	dumper.print(0, "&ast.ExprCastObject{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastString(node *ast.ExprCastString) {
	dumper.print(0, "&ast.ExprCastString{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ExprCastUnset(node *ast.ExprCastUnset) {
	dumper.print(0, "&ast.ExprCastUnset{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("CastTkn", node.CastTkn)
	dumper.dumpVertex("Expr", node.Expr)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarDnumber(node *ast.ScalarDnumber) {
	dumper.print(0, "&ast.ScalarDnumber{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NumberTkn", node.NumberTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarEncapsed(node *ast.ScalarEncapsed) {
	dumper.print(0, "&ast.ScalarEncapsed{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenQuoteTkn", node.OpenQuoteTkn)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpToken("CloseQuoteTkn", node.CloseQuoteTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarEncapsedStringPart(node *ast.ScalarEncapsedStringPart) {
	dumper.print(0, "&ast.ScalarEncapsedStringPart{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("EncapsedStrTkn", node.EncapsedStrTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarEncapsedStringVar(node *ast.ScalarEncapsedStringVar) {
	dumper.print(0, "&ast.ScalarEncapsedStringVar{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("DollarOpenCurlyBracketTkn", node.DollarOpenCurlyBracketTkn)
	dumper.dumpVertex("Name", node.Name)
	dumper.dumpToken("OpenSquareBracketTkn", node.OpenSquareBracketTkn)
	dumper.dumpVertex("Dim", node.Dim)
	dumper.dumpToken("CloseSquareBracketTkn", node.CloseSquareBracketTkn)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarEncapsedStringBrackets(node *ast.ScalarEncapsedStringBrackets) {
	dumper.print(0, "&ast.ScalarEncapsedStringBrackets{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenCurlyBracketTkn", node.OpenCurlyBracketTkn)
	dumper.dumpVertex("Var", node.Var)
	dumper.dumpToken("CloseCurlyBracketTkn", node.CloseCurlyBracketTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarHeredoc(node *ast.ScalarHeredoc) {
	dumper.print(0, "&ast.ScalarHeredoc{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("OpenHeredocTkn", node.OpenHeredocTkn)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpToken("CloseHeredocTkn", node.CloseHeredocTkn)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarLnumber(node *ast.ScalarLnumber) {
	dumper.print(0, "&ast.ScalarLnumber{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NumberTkn", node.NumberTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarMagicConstant(node *ast.ScalarMagicConstant) {
	dumper.print(0, "&ast.ScalarMagicConstant{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("MagicConstTkn", node.MagicConstTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) ScalarString(node *ast.ScalarString) {
	dumper.print(0, "&ast.ScalarString{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("MinusTkn", node.MinusTkn)
	dumper.dumpToken("StringTkn", node.StringTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) NameName(node *ast.Name) {
	dumper.print(0, "&ast.Name{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) NameFullyQualified(node *ast.NameFullyQualified) {
	dumper.print(0, "&ast.NameFullyQualified{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NsSeparatorTkn", node.NsSeparatorTkn)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) NameRelative(node *ast.NameRelative) {
	dumper.print(0, "&ast.NameRelative{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("NsTkn", node.NsTkn)
	dumper.dumpToken("NsSeparatorTkn", node.NsSeparatorTkn)
	dumper.dumpVertexList("Parts", node.Parts)
	dumper.dumpTokenList("SeparatorTkns", node.SeparatorTkns)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}

func (dumper *Dumper) NameNamePart(node *ast.NamePart) {
	dumper.print(0, "&ast.NamePart{\n")
	dumper.indent++

	dumper.dumpPosition(node.Position)
	dumper.dumpToken("StringTkn", node.StringTkn)
	dumper.dumpValue("Value", node.Value)

	dumper.indent--
	dumper.print(dumper.indent, "},\n")
}
