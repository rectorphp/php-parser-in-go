package visitor

import (
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

type Null struct {
}

func (visitor *Null) Enter(_ string, _ bool) {
	// do nothing
}
func (visitor *Null) Leave(_ string, _ bool) {
	// do nothing
}

func (visitor *Null) EnterNode(_ ast.Vertex) bool {
	return true
}

func (visitor *Null) LeaveNode(_ ast.Vertex) {
	// do nothing
}

func (visitor *Null) Root(_ *ast.Root) {
	// do nothing
}

func (visitor *Null) Nullable(_ *ast.Nullable) {
	// do nothing
}

func (visitor *Null) Intersection(_ *ast.Intersection) {
}

func (visitor *Null) Union(_ *ast.Union) {
	// do nothing
}

func (visitor *Null) Parameter(_ *ast.Parameter) {
	// do nothing
}

func (visitor *Null) Identifier(_ *ast.Identifier) {
	// do nothing
}

func (visitor *Null) Argument(_ *ast.Argument) {
	// do nothing
}

func (visitor *Null) Attribute(_ *ast.Attribute) {
	// do nothing
}

func (visitor *Null) AttributeGroup(_ *ast.AttributeGroup) {
	// do nothing
}

func (visitor *Null) StmtBreak(_ *ast.StmtBreak) {
	// do nothing
}

func (visitor *Null) StmtCase(_ *ast.StmtCase) {
	// do nothing
}

func (visitor *Null) StmtCatch(_ *ast.StmtCatch) {
	// do nothing
}

func (visitor *Null) StmtClass(_ *ast.StmtClass) {
	// do nothing
}

func (visitor *Null) StmtClassConstList(_ *ast.StmtClassConstList) {
	// do nothing
}

func (visitor *Null) StmtClassMethod(_ *ast.StmtClassMethod) {
	// do nothing
}

func (visitor *Null) StmtConstList(_ *ast.StmtConstList) {
	// do nothing
}

func (visitor *Null) StmtConstant(_ *ast.StmtConstant) {
	// do nothing
}

func (visitor *Null) StmtContinue(_ *ast.StmtContinue) {
	// do nothing
}

func (visitor *Null) StmtDeclare(_ *ast.StmtDeclare) {
	// do nothing
}

func (visitor *Null) StmtDefault(_ *ast.StmtDefault) {
	// do nothing
}

func (visitor *Null) StmtDo(_ *ast.StmtDo) {
	// do nothing
}

func (visitor *Null) StmtEcho(_ *ast.StmtEcho) {
	// do nothing
}

func (visitor *Null) StmtElse(_ *ast.StmtElse) {
	// do nothing
}

func (visitor *Null) StmtElseIf(_ *ast.StmtElseIf) {
	// do nothing
}

func (visitor *Null) StmtExpression(_ *ast.StmtExpression) {
	// do nothing
}

func (visitor *Null) StmtFinally(_ *ast.StmtFinally) {
	// do nothing
}

func (visitor *Null) StmtFor(_ *ast.StmtFor) {
	// do nothing
}

func (visitor *Null) StmtForeach(_ *ast.StmtForeach) {
	// do nothing
}

func (visitor *Null) StmtFunction(_ *ast.StmtFunction) {
	// do nothing
}

func (visitor *Null) StmtGlobal(_ *ast.StmtGlobal) {
	// do nothing
}

func (visitor *Null) StmtGoto(_ *ast.StmtGoto) {
	// do nothing
}

func (visitor *Null) StmtHaltCompiler(_ *ast.StmtHaltCompiler) {
	// do nothing
}

func (visitor *Null) StmtIf(_ *ast.StmtIf) {
	// do nothing
}

func (visitor *Null) StmtInlineHtml(_ *ast.StmtInlineHtml) {
	// do nothing
}

func (visitor *Null) StmtEnum(_ *ast.StmtEnum) {
}

func (visitor *Null) StmtEnumCase(_ *ast.StmtEnumCase) {
}

func (visitor *Null) StmtInterface(_ *ast.StmtInterface) {
	// do nothing
}

func (visitor *Null) StmtLabel(_ *ast.StmtLabel) {
	// do nothing
}

func (visitor *Null) StmtNamespace(_ *ast.StmtNamespace) {
	// do nothing
}

func (visitor *Null) StmtNop(_ *ast.StmtNop) {
	// do nothing
}

func (visitor *Null) StmtProperty(_ *ast.StmtProperty) {
	// do nothing
}

func (visitor *Null) StmtPropertyList(_ *ast.StmtPropertyList) {
	// do nothing
}

func (visitor *Null) StmtReturn(_ *ast.StmtReturn) {
	// do nothing
}

func (visitor *Null) StmtStatic(_ *ast.StmtStatic) {
	// do nothing
}

func (visitor *Null) StmtStaticVar(_ *ast.StmtStaticVar) {
	// do nothing
}

func (visitor *Null) StmtStmtList(_ *ast.StmtStmtList) {
	// do nothing
}

func (visitor *Null) StmtSwitch(_ *ast.StmtSwitch) {
	// do nothing
}

func (visitor *Null) StmtThrow(_ *ast.StmtThrow) {
	// do nothing
}

func (visitor *Null) StmtTrait(_ *ast.StmtTrait) {
	// do nothing
}

func (visitor *Null) StmtTraitUse(_ *ast.StmtTraitUse) {
	// do nothing
}

func (visitor *Null) StmtTraitUseAlias(_ *ast.StmtTraitUseAlias) {
	// do nothing
}

func (visitor *Null) StmtTraitUsePrecedence(_ *ast.StmtTraitUsePrecedence) {
	// do nothing
}

func (visitor *Null) StmtTry(_ *ast.StmtTry) {
	// do nothing
}

func (visitor *Null) StmtUnset(_ *ast.StmtUnset) {
	// do nothing
}

func (visitor *Null) StmtUse(_ *ast.StmtUseList) {
	// do nothing
}

func (visitor *Null) StmtGroupUse(_ *ast.StmtGroupUseList) {
	// do nothing
}

func (visitor *Null) StmtUseDeclaration(_ *ast.StmtUse) {
	// do nothing
}

func (visitor *Null) StmtWhile(_ *ast.StmtWhile) {
	// do nothing
}

func (visitor *Null) ExprArray(_ *ast.ExprArray) {
	// do nothing
}

func (visitor *Null) ExprArrayDimFetch(_ *ast.ExprArrayDimFetch) {
	// do nothing
}

func (visitor *Null) ExprArrayItem(_ *ast.ExprArrayItem) {
	// do nothing
}

func (visitor *Null) ExprArrowFunction(_ *ast.ExprArrowFunction) {
	// do nothing
}

func (visitor *Null) ExprBitwiseNot(_ *ast.ExprBitwiseNot) {
	// do nothing
}

func (visitor *Null) ExprBooleanNot(_ *ast.ExprBooleanNot) {
	// do nothing
}

func (visitor *Null) ExprBrackets(_ *ast.ExprBrackets) {
	// do nothing
}

func (visitor *Null) ExprClassConstFetch(_ *ast.ExprClassConstFetch) {
	// do nothing
}

func (visitor *Null) ExprClone(_ *ast.ExprClone) {
	// do nothing
}

func (visitor *Null) ExprThrow(_ *ast.ExprThrow) {
	// do nothing
}

func (visitor *Null) ExprClosure(_ *ast.ExprClosure) {
	// do nothing
}

func (visitor *Null) ExprClosureUse(_ *ast.ExprClosureUse) {
	// do nothing
}

func (visitor *Null) ExprConstFetch(_ *ast.ExprConstFetch) {
	// do nothing
}

func (visitor *Null) ExprEmpty(_ *ast.ExprEmpty) {
	// do nothing
}

func (visitor *Null) ExprErrorSuppress(_ *ast.ExprErrorSuppress) {
	// do nothing
}

func (visitor *Null) ExprEval(_ *ast.ExprEval) {
	// do nothing
}

func (visitor *Null) ExprExit(_ *ast.ExprExit) {
	// do nothing
}

func (visitor *Null) ExprFunctionCall(_ *ast.ExprFunctionCall) {
	// do nothing
}

func (visitor *Null) ExprInclude(_ *ast.ExprInclude) {
	// do nothing
}

func (visitor *Null) ExprIncludeOnce(_ *ast.ExprIncludeOnce) {
	// do nothing
}

func (visitor *Null) ExprInstanceOf(_ *ast.ExprInstanceOf) {
	// do nothing
}

func (visitor *Null) ExprIsset(_ *ast.ExprIsset) {
	// do nothing
}

func (visitor *Null) ExprList(_ *ast.ExprList) {
	// do nothing
}

func (visitor *Null) ExprMethodCall(_ *ast.ExprMethodCall) {
	// do nothing
}

func (visitor *Null) ExprNullsafeMethodCall(_ *ast.ExprNullsafeMethodCall) {
	// do nothing
}

func (visitor *Null) ExprNullsafePropertyFetch(_ *ast.ExprNullsafePropertyFetch) {
	// do nothing
}

func (visitor *Null) ExprNew(_ *ast.ExprNew) {
	// do nothing
}

func (visitor *Null) ExprPostDec(_ *ast.ExprPostDec) {
	// do nothing
}

func (visitor *Null) ExprPostInc(_ *ast.ExprPostInc) {
	// do nothing
}

func (visitor *Null) ExprPreDec(_ *ast.ExprPreDec) {
	// do nothing
}

func (visitor *Null) ExprPreInc(_ *ast.ExprPreInc) {
	// do nothing
}

func (visitor *Null) ExprPrint(_ *ast.ExprPrint) {
	// do nothing
}

func (visitor *Null) ExprPropertyFetch(_ *ast.ExprPropertyFetch) {
	// do nothing
}

func (visitor *Null) ExprRequire(_ *ast.ExprRequire) {
	// do nothing
}

func (visitor *Null) ExprRequireOnce(_ *ast.ExprRequireOnce) {
	// do nothing
}

func (visitor *Null) ExprShellExec(_ *ast.ExprShellExec) {
	// do nothing
}

func (visitor *Null) ExprStaticCall(_ *ast.ExprStaticCall) {
	// do nothing
}

func (visitor *Null) ExprStaticPropertyFetch(_ *ast.ExprStaticPropertyFetch) {
	// do nothing
}

func (visitor *Null) ExprTernary(_ *ast.ExprTernary) {
	// do nothing
}

func (visitor *Null) ExprMatch(_ *ast.ExprMatch) {
	// do nothing
}

func (visitor *Null) MatchArm(_ *ast.MatchArm) {
	// do nothing
}

func (visitor *Null) ExprUnaryMinus(_ *ast.ExprUnaryMinus) {
	// do nothing
}

func (visitor *Null) ExprUnaryPlus(_ *ast.ExprUnaryPlus) {
	// do nothing
}

func (visitor *Null) ExprVariable(_ *ast.ExprVariable) {
	// do nothing
}

func (visitor *Null) ExprYield(_ *ast.ExprYield) {
	// do nothing
}

func (visitor *Null) ExprYieldFrom(_ *ast.ExprYieldFrom) {
	// do nothing
}

func (visitor *Null) ExprAssign(_ *ast.ExprAssign) {
	// do nothing
}

func (visitor *Null) ExprAssignReference(_ *ast.ExprAssignReference) {
	// do nothing
}

func (visitor *Null) ExprAssignBitwiseAnd(_ *ast.ExprAssignBitwiseAnd) {
	// do nothing
}

func (visitor *Null) ExprAssignBitwiseOr(_ *ast.ExprAssignBitwiseOr) {
	// do nothing
}

func (visitor *Null) ExprAssignBitwiseXor(_ *ast.ExprAssignBitwiseXor) {
	// do nothing
}

func (visitor *Null) ExprAssignCoalesce(_ *ast.ExprAssignCoalesce) {
	// do nothing
}

func (visitor *Null) ExprAssignConcat(_ *ast.ExprAssignConcat) {
	// do nothing
}

func (visitor *Null) ExprAssignDiv(_ *ast.ExprAssignDiv) {
	// do nothing
}

func (visitor *Null) ExprAssignMinus(_ *ast.ExprAssignMinus) {
	// do nothing
}

func (visitor *Null) ExprAssignMod(_ *ast.ExprAssignMod) {
	// do nothing
}

func (visitor *Null) ExprAssignMul(_ *ast.ExprAssignMul) {
	// do nothing
}

func (visitor *Null) ExprAssignPlus(_ *ast.ExprAssignPlus) {
	// do nothing
}

func (visitor *Null) ExprAssignPow(_ *ast.ExprAssignPow) {
	// do nothing
}

func (visitor *Null) ExprAssignShiftLeft(_ *ast.ExprAssignShiftLeft) {
	// do nothing
}

func (visitor *Null) ExprAssignShiftRight(_ *ast.ExprAssignShiftRight) {
	// do nothing
}

func (visitor *Null) ExprBinaryBitwiseAnd(_ *ast.ExprBinaryBitwiseAnd) {
	// do nothing
}

func (visitor *Null) ExprBinaryBitwiseOr(_ *ast.ExprBinaryBitwiseOr) {
	// do nothing
}

func (visitor *Null) ExprBinaryBitwiseXor(_ *ast.ExprBinaryBitwiseXor) {
	// do nothing
}

func (visitor *Null) ExprBinaryBooleanAnd(_ *ast.ExprBinaryBooleanAnd) {
	// do nothing
}

func (visitor *Null) ExprBinaryBooleanOr(_ *ast.ExprBinaryBooleanOr) {
	// do nothing
}

func (visitor *Null) ExprBinaryCoalesce(_ *ast.ExprBinaryCoalesce) {
	// do nothing
}

func (visitor *Null) ExprBinaryConcat(_ *ast.ExprBinaryConcat) {
	// do nothing
}

func (visitor *Null) ExprBinaryDiv(_ *ast.ExprBinaryDiv) {
	// do nothing
}

func (visitor *Null) ExprBinaryEqual(_ *ast.ExprBinaryEqual) {
	// do nothing
}

func (visitor *Null) ExprBinaryGreater(_ *ast.ExprBinaryGreater) {
	// do nothing
}

func (visitor *Null) ExprBinaryGreaterOrEqual(_ *ast.ExprBinaryGreaterOrEqual) {
	// do nothing
}

func (visitor *Null) ExprBinaryIdentical(_ *ast.ExprBinaryIdentical) {
	// do nothing
}

func (visitor *Null) ExprBinaryLogicalAnd(_ *ast.ExprBinaryLogicalAnd) {
	// do nothing
}

func (visitor *Null) ExprBinaryLogicalOr(_ *ast.ExprBinaryLogicalOr) {
	// do nothing
}

func (visitor *Null) ExprBinaryLogicalXor(_ *ast.ExprBinaryLogicalXor) {
	// do nothing
}

func (visitor *Null) ExprBinaryMinus(_ *ast.ExprBinaryMinus) {
	// do nothing
}

func (visitor *Null) ExprBinaryMod(_ *ast.ExprBinaryMod) {
	// do nothing
}

func (visitor *Null) ExprBinaryMul(_ *ast.ExprBinaryMul) {
	// do nothing
}

func (visitor *Null) ExprBinaryNotEqual(_ *ast.ExprBinaryNotEqual) {
	// do nothing
}

func (visitor *Null) ExprBinaryNotIdentical(_ *ast.ExprBinaryNotIdentical) {
	// do nothing
}

func (visitor *Null) ExprBinaryPlus(_ *ast.ExprBinaryPlus) {
	// do nothing
}

func (visitor *Null) ExprBinaryPow(_ *ast.ExprBinaryPow) {
	// do nothing
}

func (visitor *Null) ExprBinaryShiftLeft(_ *ast.ExprBinaryShiftLeft) {
	// do nothing
}

func (visitor *Null) ExprBinaryShiftRight(_ *ast.ExprBinaryShiftRight) {
	// do nothing
}

func (visitor *Null) ExprBinarySmaller(_ *ast.ExprBinarySmaller) {
	// do nothing
}

func (visitor *Null) ExprBinarySmallerOrEqual(_ *ast.ExprBinarySmallerOrEqual) {
	// do nothing
}

func (visitor *Null) ExprBinarySpaceship(_ *ast.ExprBinarySpaceship) {
	// do nothing
}

func (visitor *Null) ExprCastArray(_ *ast.ExprCastArray) {
	// do nothing
}

func (visitor *Null) ExprCastBool(_ *ast.ExprCastBool) {
	// do nothing
}

func (visitor *Null) ExprCastDouble(_ *ast.ExprCastDouble) {
	// do nothing
}

func (visitor *Null) ExprCastInt(_ *ast.ExprCastInt) {
	// do nothing
}

func (visitor *Null) ExprCastObject(_ *ast.ExprCastObject) {
	// do nothing
}

func (visitor *Null) ExprCastString(_ *ast.ExprCastString) {
	// do nothing
}

func (visitor *Null) ExprCastUnset(_ *ast.ExprCastUnset) {
	// do nothing
}

func (visitor *Null) ScalarDnumber(_ *ast.ScalarDnumber) {
	// do nothing
}

func (visitor *Null) ScalarEncapsed(_ *ast.ScalarEncapsed) {
	// do nothing
}

func (visitor *Null) ScalarEncapsedStringPart(_ *ast.ScalarEncapsedStringPart) {
	// do nothing
}

func (visitor *Null) ScalarEncapsedStringBrackets(_ *ast.ScalarEncapsedStringBrackets) {
	// do nothing
}

func (visitor *Null) ScalarEncapsedStringVar(_ *ast.ScalarEncapsedStringVar) {
	// do nothing
}

func (visitor *Null) ScalarHeredoc(_ *ast.ScalarHeredoc) {
	// do nothing
}

func (visitor *Null) ScalarLnumber(_ *ast.ScalarLnumber) {
	// do nothing
}

func (visitor *Null) ScalarMagicConstant(_ *ast.ScalarMagicConstant) {
	// do nothing
}

func (visitor *Null) ScalarString(_ *ast.ScalarString) {
	// do nothing
}

func (visitor *Null) NameName(_ *ast.Name) {
	// do nothing
}

func (visitor *Null) NameFullyQualified(_ *ast.NameFullyQualified) {
	// do nothing
}

func (visitor *Null) NameRelative(_ *ast.NameRelative) {
	// do nothing
}

func (visitor *Null) NameNamePart(_ *ast.NamePart) {
	// do nothing
}
