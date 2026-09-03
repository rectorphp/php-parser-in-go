package traverser

import (
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

type Traverser struct {
	v ast.Visitor
}

func NewTraverser(visitor ast.Visitor) *Traverser {
	return &Traverser{
		v: visitor,
	}
}

func (traverser *Traverser) Traverse(node ast.Vertex) {
	if node != nil {
		node.Accept(traverser)
	}
}

func (traverser *Traverser) Root(node *ast.Root) {
	node.Accept(traverser.v)

	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) Nullable(node *ast.Nullable) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) Union(node *ast.Union) {
	node.Accept(traverser.v)

	for _, childNode := range node.Types {
		traverser.Traverse(childNode)
	}
}

func (traverser *Traverser) Intersection(node *ast.Intersection) {
	node.Accept(traverser.v)

	for _, childNode := range node.Types {
		traverser.Traverse(childNode)
	}
}

func (traverser *Traverser) Parameter(node *ast.Parameter) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		traverser.Traverse(childNode)
	}
	for _, childNode := range node.Modifiers {
		traverser.Traverse(childNode)
	}
	traverser.Traverse(node.Type)
	traverser.Traverse(node.Var)
	traverser.Traverse(node.DefaultValue)
}

func (traverser *Traverser) Identifier(node *ast.Identifier) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) Argument(node *ast.Argument) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) Attribute(node *ast.Attribute) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) AttributeGroup(node *ast.AttributeGroup) {
	node.Accept(traverser.v)

	for _, childNode := range node.Attrs {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtBreak(node *ast.StmtBreak) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtCase(node *ast.StmtCase) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtCatch(node *ast.StmtCatch) {
	node.Accept(traverser.v)

	for _, childNode := range node.Types {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Var)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtClass(node *ast.StmtClass) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Modifiers {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Extends)
	for _, childNode := range node.Implements {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtClassConstList(node *ast.StmtClassConstList) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Modifiers {
		childNode.Accept(traverser)
	}
	if node.Type != nil {
		node.Type.Accept(traverser)
	}
	for _, childNode := range node.Consts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtClassMethod(node *ast.StmtClassMethod) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Modifiers {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	for _, childNode := range node.Params {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.ReturnType)
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtConstList(node *ast.StmtConstList) {
	node.Accept(traverser.v)

	for _, childNode := range node.Consts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtConstant(node *ast.StmtConstant) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtContinue(node *ast.StmtContinue) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtDeclare(node *ast.StmtDeclare) {
	node.Accept(traverser.v)

	for _, childNode := range node.Consts {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtDefault(node *ast.StmtDefault) {
	node.Accept(traverser.v)

	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtDo(node *ast.StmtDo) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Stmt)
	traverser.Traverse(node.Cond)
}

func (traverser *Traverser) StmtEcho(node *ast.StmtEcho) {
	node.Accept(traverser.v)

	for _, childNode := range node.Exprs {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtElse(node *ast.StmtElse) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtElseIf(node *ast.StmtElseIf) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtExpression(node *ast.StmtExpression) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtFinally(node *ast.StmtFinally) {
	node.Accept(traverser.v)

	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtFor(node *ast.StmtFor) {
	node.Accept(traverser.v)

	for _, childNode := range node.Init {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Cond {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Loop {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtForeach(node *ast.StmtForeach) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
	traverser.Traverse(node.Key)
	traverser.Traverse(node.Var)
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) StmtFunction(node *ast.StmtFunction) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	for _, childNode := range node.Params {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.ReturnType)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtGlobal(node *ast.StmtGlobal) {
	node.Accept(traverser.v)

	for _, childNode := range node.Vars {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtGoto(node *ast.StmtGoto) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Label)
}

func (traverser *Traverser) StmtHaltCompiler(node *ast.StmtHaltCompiler) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) StmtIf(node *ast.StmtIf) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	traverser.Traverse(node.Stmt)
	for _, childNode := range node.ElseIf {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Else)
}

func (traverser *Traverser) StmtInlineHtml(node *ast.StmtInlineHtml) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) StmtEnum(node *ast.StmtEnum) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	traverser.Traverse(node.Type)
	for _, childNode := range node.Implements {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtEnumCase(node *ast.StmtEnumCase) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtInterface(node *ast.StmtInterface) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	for _, childNode := range node.Extends {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtLabel(node *ast.StmtLabel) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
}

func (traverser *Traverser) StmtNamespace(node *ast.StmtNamespace) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtNop(node *ast.StmtNop) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) StmtProperty(node *ast.StmtProperty) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtPropertyList(node *ast.StmtPropertyList) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Modifiers {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Type)
	for _, childNode := range node.Props {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtReturn(node *ast.StmtReturn) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtStatic(node *ast.StmtStatic) {
	node.Accept(traverser.v)

	for _, childNode := range node.Vars {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtStaticVar(node *ast.StmtStaticVar) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtStmtList(node *ast.StmtStmtList) {
	node.Accept(traverser.v)

	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtSwitch(node *ast.StmtSwitch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	for _, childNode := range node.Cases {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtThrow(node *ast.StmtThrow) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) StmtTrait(node *ast.StmtTrait) {
	node.Accept(traverser.v)

	for _, childNode := range node.AttrGroups {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Name)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtTraitUse(node *ast.StmtTraitUse) {
	node.Accept(traverser.v)

	for _, childNode := range node.Traits {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Adaptations {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtTraitUseAlias(node *ast.StmtTraitUseAlias) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Trait)
	traverser.Traverse(node.Method)
	traverser.Traverse(node.Modifier)
	traverser.Traverse(node.Alias)
}

func (traverser *Traverser) StmtTraitUsePrecedence(node *ast.StmtTraitUsePrecedence) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Trait)
	traverser.Traverse(node.Method)
	for _, childNode := range node.Insteadof {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtTry(node *ast.StmtTry) {
	node.Accept(traverser.v)

	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Catches {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.Finally)
}

func (traverser *Traverser) StmtUnset(node *ast.StmtUnset) {
	node.Accept(traverser.v)

	for _, childNode := range node.Vars {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtUse(node *ast.StmtUseList) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Type)
	for _, childNode := range node.Uses {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtGroupUse(node *ast.StmtGroupUseList) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Type)
	traverser.Traverse(node.Prefix)
	for _, childNode := range node.Uses {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) StmtUseDeclaration(node *ast.StmtUse) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Type)
	traverser.Traverse(node.Use)
	traverser.Traverse(node.Alias)
}

func (traverser *Traverser) StmtWhile(node *ast.StmtWhile) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	traverser.Traverse(node.Stmt)
}

func (traverser *Traverser) ExprArray(node *ast.ExprArray) {
	node.Accept(traverser.v)

	for _, childNode := range node.Items {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprArrayDimFetch(node *ast.ExprArrayDimFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Dim)
}

func (traverser *Traverser) ExprArrayItem(node *ast.ExprArrayItem) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Key)
	traverser.Traverse(node.Val)
}

func (traverser *Traverser) ExprArrowFunction(node *ast.ExprArrowFunction) {
	node.Accept(traverser.v)

	for _, childNode := range node.Params {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.ReturnType)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprBitwiseNot(node *ast.ExprBitwiseNot) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprBooleanNot(node *ast.ExprBooleanNot) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprBrackets(node *ast.ExprBrackets) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprClassConstFetch(node *ast.ExprClassConstFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Class)
	traverser.Traverse(node.Const)
}

func (traverser *Traverser) ExprClone(node *ast.ExprClone) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprThrow(node *ast.ExprThrow) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprClosure(node *ast.ExprClosure) {
	node.Accept(traverser.v)

	for _, childNode := range node.Params {
		childNode.Accept(traverser)
	}
	for _, childNode := range node.Uses {
		childNode.Accept(traverser)
	}
	traverser.Traverse(node.ReturnType)
	for _, childNode := range node.Stmts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprClosureUse(node *ast.ExprClosureUse) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ExprConstFetch(node *ast.ExprConstFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Const)
}

func (traverser *Traverser) ExprEmpty(node *ast.ExprEmpty) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprErrorSuppress(node *ast.ExprErrorSuppress) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprEval(node *ast.ExprEval) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprExit(node *ast.ExprExit) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprFunctionCall(node *ast.ExprFunctionCall) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Function)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprInclude(node *ast.ExprInclude) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprIncludeOnce(node *ast.ExprIncludeOnce) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprInstanceOf(node *ast.ExprInstanceOf) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
	traverser.Traverse(node.Class)
}

func (traverser *Traverser) ExprIsset(node *ast.ExprIsset) {
	node.Accept(traverser.v)

	for _, childNode := range node.Vars {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprList(node *ast.ExprList) {
	node.Accept(traverser.v)

	for _, childNode := range node.Items {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprMethodCall(node *ast.ExprMethodCall) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Method)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprNullsafeMethodCall(node *ast.ExprNullsafeMethodCall) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Method)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprNullsafePropertyFetch(node *ast.ExprNullsafePropertyFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Prop)
}

func (traverser *Traverser) ExprNew(node *ast.ExprNew) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Class)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprPostDec(node *ast.ExprPostDec) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ExprPostInc(node *ast.ExprPostInc) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ExprPreDec(node *ast.ExprPreDec) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ExprPreInc(node *ast.ExprPreInc) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ExprPrint(node *ast.ExprPrint) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprPropertyFetch(node *ast.ExprPropertyFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Prop)
}

func (traverser *Traverser) ExprRequire(node *ast.ExprRequire) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprRequireOnce(node *ast.ExprRequireOnce) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprShellExec(node *ast.ExprShellExec) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprStaticCall(node *ast.ExprStaticCall) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Class)
	traverser.Traverse(node.Call)
	for _, childNode := range node.Args {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ExprStaticPropertyFetch(node *ast.ExprStaticPropertyFetch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Class)
	traverser.Traverse(node.Prop)
}

func (traverser *Traverser) ExprTernary(node *ast.ExprTernary) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Cond)
	traverser.Traverse(node.IfTrue)
	traverser.Traverse(node.IfFalse)
}

func (traverser *Traverser) ExprMatch(node *ast.ExprMatch) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
	for _, childNode := range node.Arms {
		traverser.Traverse(childNode)
	}
}

func (traverser *Traverser) MatchArm(node *ast.MatchArm) {
	node.Accept(traverser.v)

	for _, childNode := range node.Exprs {
		traverser.Traverse(childNode)
	}
	traverser.Traverse(node.ReturnExpr)
}

func (traverser *Traverser) ExprUnaryMinus(node *ast.ExprUnaryMinus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprUnaryPlus(node *ast.ExprUnaryPlus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprVariable(node *ast.ExprVariable) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
}

func (traverser *Traverser) ExprYield(node *ast.ExprYield) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Key)
	traverser.Traverse(node.Val)
}

func (traverser *Traverser) ExprYieldFrom(node *ast.ExprYieldFrom) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssign(node *ast.ExprAssign) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignReference(node *ast.ExprAssignReference) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignBitwiseAnd(node *ast.ExprAssignBitwiseAnd) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignBitwiseOr(node *ast.ExprAssignBitwiseOr) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignBitwiseXor(node *ast.ExprAssignBitwiseXor) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignCoalesce(node *ast.ExprAssignCoalesce) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignConcat(node *ast.ExprAssignConcat) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignDiv(node *ast.ExprAssignDiv) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignMinus(node *ast.ExprAssignMinus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignMod(node *ast.ExprAssignMod) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignMul(node *ast.ExprAssignMul) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignPlus(node *ast.ExprAssignPlus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignPow(node *ast.ExprAssignPow) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignShiftLeft(node *ast.ExprAssignShiftLeft) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprAssignShiftRight(node *ast.ExprAssignShiftRight) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprBinaryBitwiseAnd(node *ast.ExprBinaryBitwiseAnd) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryBitwiseOr(node *ast.ExprBinaryBitwiseOr) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryBitwiseXor(node *ast.ExprBinaryBitwiseXor) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryBooleanAnd(node *ast.ExprBinaryBooleanAnd) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryBooleanOr(node *ast.ExprBinaryBooleanOr) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryCoalesce(node *ast.ExprBinaryCoalesce) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryConcat(node *ast.ExprBinaryConcat) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryDiv(node *ast.ExprBinaryDiv) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryEqual(node *ast.ExprBinaryEqual) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryGreater(node *ast.ExprBinaryGreater) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryGreaterOrEqual(node *ast.ExprBinaryGreaterOrEqual) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryIdentical(node *ast.ExprBinaryIdentical) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryLogicalAnd(node *ast.ExprBinaryLogicalAnd) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryLogicalOr(node *ast.ExprBinaryLogicalOr) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryLogicalXor(node *ast.ExprBinaryLogicalXor) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryMinus(node *ast.ExprBinaryMinus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryMod(node *ast.ExprBinaryMod) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryMul(node *ast.ExprBinaryMul) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryNotEqual(node *ast.ExprBinaryNotEqual) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryNotIdentical(node *ast.ExprBinaryNotIdentical) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryPlus(node *ast.ExprBinaryPlus) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryPow(node *ast.ExprBinaryPow) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryShiftLeft(node *ast.ExprBinaryShiftLeft) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinaryShiftRight(node *ast.ExprBinaryShiftRight) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinarySmaller(node *ast.ExprBinarySmaller) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinarySmallerOrEqual(node *ast.ExprBinarySmallerOrEqual) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprBinarySpaceship(node *ast.ExprBinarySpaceship) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Left)
	traverser.Traverse(node.Right)
}

func (traverser *Traverser) ExprCastArray(node *ast.ExprCastArray) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastBool(node *ast.ExprCastBool) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastDouble(node *ast.ExprCastDouble) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastInt(node *ast.ExprCastInt) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastObject(node *ast.ExprCastObject) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastString(node *ast.ExprCastString) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ExprCastUnset(node *ast.ExprCastUnset) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Expr)
}

func (traverser *Traverser) ScalarDnumber(node *ast.ScalarDnumber) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) ScalarEncapsed(node *ast.ScalarEncapsed) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ScalarEncapsedStringPart(node *ast.ScalarEncapsedStringPart) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) ScalarEncapsedStringVar(node *ast.ScalarEncapsedStringVar) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Name)
	traverser.Traverse(node.Dim)
}

func (traverser *Traverser) ScalarEncapsedStringBrackets(node *ast.ScalarEncapsedStringBrackets) {
	node.Accept(traverser.v)

	traverser.Traverse(node.Var)
}

func (traverser *Traverser) ScalarHeredoc(node *ast.ScalarHeredoc) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) ScalarLnumber(node *ast.ScalarLnumber) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) ScalarMagicConstant(node *ast.ScalarMagicConstant) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) ScalarString(node *ast.ScalarString) {
	node.Accept(traverser.v)
}

func (traverser *Traverser) NameName(node *ast.Name) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) NameFullyQualified(node *ast.NameFullyQualified) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) NameRelative(node *ast.NameRelative) {
	node.Accept(traverser.v)

	for _, childNode := range node.Parts {
		childNode.Accept(traverser)
	}
}

func (traverser *Traverser) NameNamePart(node *ast.NamePart) {
	node.Accept(traverser.v)
}
