package ast

import (
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

// Root node
type Root struct {
	Position *position.Position
	Stmts    []Vertex
	EndTkn   *token.Token
}

func (node *Root) Accept(visitor Visitor) {
	visitor.Root(node)
}

func (node *Root) GetPosition() *position.Position {
	return node.Position
}

// Nullable node
type Nullable struct {
	Position    *position.Position
	QuestionTkn *token.Token
	Expr        Vertex
}

func (node *Nullable) Accept(visitor Visitor) {
	visitor.Nullable(node)
}

func (node *Nullable) GetPosition() *position.Position {
	return node.Position
}

// Union node holds a union type such as `int|string`, its member types
// separated by `|`.
type Union struct {
	Position      *position.Position
	Types         []Vertex
	SeparatorTkns []*token.Token
}

func (node *Union) Accept(visitor Visitor) {
	visitor.Union(node)
}

func (node *Union) GetPosition() *position.Position {
	return node.Position
}

// Intersection node holds an intersection type such as `Countable&Traversable`,
// its member types separated by `&`.
type Intersection struct {
	Position      *position.Position
	Types         []Vertex
	SeparatorTkns []*token.Token
}

func (node *Intersection) Accept(visitor Visitor) {
	visitor.Intersection(node)
}

func (node *Intersection) GetPosition() *position.Position {
	return node.Position
}

// Parameter node
type Parameter struct {
	Position     *position.Position
	AttrGroups   []Vertex
	Modifiers    []Vertex
	Type         Vertex
	AmpersandTkn *token.Token
	VariadicTkn  *token.Token
	Var          Vertex
	EqualTkn     *token.Token
	DefaultValue Vertex
}

func (node *Parameter) Accept(visitor Visitor) {
	visitor.Parameter(node)
}

func (node *Parameter) GetPosition() *position.Position {
	return node.Position
}

// Identifier node
type Identifier struct {
	Position      *position.Position
	IdentifierTkn *token.Token
	Value         []byte
}

func (node *Identifier) Accept(visitor Visitor) {
	visitor.Identifier(node)
}

func (node *Identifier) GetPosition() *position.Position {
	return node.Position
}

// Argument node
type Argument struct {
	Position     *position.Position
	Name         Vertex
	ColonTkn     *token.Token
	VariadicTkn  *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

func (node *Argument) Accept(visitor Visitor) {
	visitor.Argument(node)
}

func (node *Argument) GetPosition() *position.Position {
	return node.Position
}

// Attribute node
type Attribute struct {
	Position            *position.Position
	Name                Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (node *Attribute) Accept(visitor Visitor) {
	visitor.Attribute(node)
}

func (node *Attribute) GetPosition() *position.Position {
	return node.Position
}

// AttributeGroup node
type AttributeGroup struct {
	Position          *position.Position
	OpenAttributeTkn  *token.Token
	Attrs             []Vertex
	SeparatorTkns     []*token.Token
	CloseAttributeTkn *token.Token
}

func (node *AttributeGroup) Accept(visitor Visitor) {
	visitor.AttributeGroup(node)
}

func (node *AttributeGroup) GetPosition() *position.Position {
	return node.Position
}

// ScalarDnumber node
type ScalarDnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

func (node *ScalarDnumber) Accept(visitor Visitor) {
	visitor.ScalarDnumber(node)
}

func (node *ScalarDnumber) GetPosition() *position.Position {
	return node.Position
}

// ScalarEncapsed node
type ScalarEncapsed struct {
	Position      *position.Position
	OpenQuoteTkn  *token.Token
	Parts         []Vertex
	CloseQuoteTkn *token.Token
}

func (node *ScalarEncapsed) Accept(visitor Visitor) {
	visitor.ScalarEncapsed(node)
}

func (node *ScalarEncapsed) GetPosition() *position.Position {
	return node.Position
}

// ScalarEncapsedStringPart node
type ScalarEncapsedStringPart struct {
	Position       *position.Position
	EncapsedStrTkn *token.Token
	Value          []byte
}

func (node *ScalarEncapsedStringPart) Accept(visitor Visitor) {
	visitor.ScalarEncapsedStringPart(node)
}

func (node *ScalarEncapsedStringPart) GetPosition() *position.Position {
	return node.Position
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringVar struct {
	Position                  *position.Position
	DollarOpenCurlyBracketTkn *token.Token
	Name                      Vertex
	OpenSquareBracketTkn      *token.Token
	Dim                       Vertex
	CloseSquareBracketTkn     *token.Token
	CloseCurlyBracketTkn      *token.Token
}

func (node *ScalarEncapsedStringVar) Accept(visitor Visitor) {
	visitor.ScalarEncapsedStringVar(node)
}

func (node *ScalarEncapsedStringVar) GetPosition() *position.Position {
	return node.Position
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringBrackets struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Var                  Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *ScalarEncapsedStringBrackets) Accept(visitor Visitor) {
	visitor.ScalarEncapsedStringBrackets(node)
}

func (node *ScalarEncapsedStringBrackets) GetPosition() *position.Position {
	return node.Position
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Position        *position.Position
	OpenHeredocTkn  *token.Token
	Parts           []Vertex
	CloseHeredocTkn *token.Token
}

func (node *ScalarHeredoc) Accept(visitor Visitor) {
	visitor.ScalarHeredoc(node)
}

func (node *ScalarHeredoc) GetPosition() *position.Position {
	return node.Position
}

// ScalarLnumber node
type ScalarLnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

func (node *ScalarLnumber) Accept(visitor Visitor) {
	visitor.ScalarLnumber(node)
}

func (node *ScalarLnumber) GetPosition() *position.Position {
	return node.Position
}

// ScalarMagicConstant node
type ScalarMagicConstant struct {
	Position      *position.Position
	MagicConstTkn *token.Token
	Value         []byte
}

func (node *ScalarMagicConstant) Accept(visitor Visitor) {
	visitor.ScalarMagicConstant(node)
}

func (node *ScalarMagicConstant) GetPosition() *position.Position {
	return node.Position
}

// ScalarString node
type ScalarString struct {
	Position  *position.Position
	MinusTkn  *token.Token
	StringTkn *token.Token
	Value     []byte
}

func (node *ScalarString) Accept(visitor Visitor) {
	visitor.ScalarString(node)
}

func (node *ScalarString) GetPosition() *position.Position {
	return node.Position
}

// StmtBreak node
type StmtBreak struct {
	Position     *position.Position
	BreakTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtBreak) Accept(visitor Visitor) {
	visitor.StmtBreak(node)
}

func (node *StmtBreak) GetPosition() *position.Position {
	return node.Position
}

// StmtCase node
type StmtCase struct {
	Position         *position.Position
	CaseTkn          *token.Token
	Cond             Vertex
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (node *StmtCase) Accept(visitor Visitor) {
	visitor.StmtCase(node)
}

func (node *StmtCase) GetPosition() *position.Position {
	return node.Position
}

// StmtCatch node
type StmtCatch struct {
	Position             *position.Position
	CatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Types                []Vertex
	SeparatorTkns        []*token.Token
	Var                  Vertex
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtCatch) Accept(visitor Visitor) {
	visitor.StmtCatch(node)
}

func (node *StmtCatch) GetPosition() *position.Position {
	return node.Position
}

// StmtClass node
type StmtClass struct {
	Position                *position.Position
	AttrGroups              []Vertex
	Modifiers               []Vertex
	ClassTkn                *token.Token
	Name                    Vertex
	OpenParenthesisTkn      *token.Token
	Args                    []Vertex
	SeparatorTkns           []*token.Token
	CloseParenthesisTkn     *token.Token
	ExtendsTkn              *token.Token
	Extends                 Vertex
	ImplementsTkn           *token.Token
	Implements              []Vertex
	ImplementsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn     *token.Token
	Stmts                   []Vertex
	CloseCurlyBracketTkn    *token.Token
}

func (node *StmtClass) Accept(visitor Visitor) {
	visitor.StmtClass(node)
}

func (node *StmtClass) GetPosition() *position.Position {
	return node.Position
}

// StmtClassConstList node
type StmtClassConstList struct {
	Position      *position.Position
	AttrGroups    []Vertex
	Modifiers     []Vertex
	ConstTkn      *token.Token
	Type          Vertex
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtClassConstList) Accept(visitor Visitor) {
	visitor.StmtClassConstList(node)
}

func (node *StmtClassConstList) GetPosition() *position.Position {
	return node.Position
}

// StmtClassMethod node
type StmtClassMethod struct {
	Position            *position.Position
	AttrGroups          []Vertex
	Modifiers           []Vertex
	FunctionTkn         *token.Token
	AmpersandTkn        *token.Token
	Name                Vertex
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	Stmt                Vertex
}

func (node *StmtClassMethod) Accept(visitor Visitor) {
	visitor.StmtClassMethod(node)
}

func (node *StmtClassMethod) GetPosition() *position.Position {
	return node.Position
}

// StmtConstList node
type StmtConstList struct {
	Position      *position.Position
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtConstList) Accept(visitor Visitor) {
	visitor.StmtConstList(node)
}

func (node *StmtConstList) GetPosition() *position.Position {
	return node.Position
}

// StmtConstant node
type StmtConstant struct {
	Position *position.Position
	Name     Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *StmtConstant) Accept(visitor Visitor) {
	visitor.StmtConstant(node)
}

func (node *StmtConstant) GetPosition() *position.Position {
	return node.Position
}

// StmtContinue node
type StmtContinue struct {
	Position     *position.Position
	ContinueTkn  *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtContinue) Accept(visitor Visitor) {
	visitor.StmtContinue(node)
}

func (node *StmtContinue) GetPosition() *position.Position {
	return node.Position
}

// StmtDeclare node
type StmtDeclare struct {
	Position            *position.Position
	DeclareTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Consts              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndDeclareTkn       *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtDeclare) Accept(visitor Visitor) {
	visitor.StmtDeclare(node)
}

func (node *StmtDeclare) GetPosition() *position.Position {
	return node.Position
}

// StmtDefault node
type StmtDefault struct {
	Position         *position.Position
	DefaultTkn       *token.Token
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (node *StmtDefault) Accept(visitor Visitor) {
	visitor.StmtDefault(node)
}

func (node *StmtDefault) GetPosition() *position.Position {
	return node.Position
}

// StmtDo node
type StmtDo struct {
	Position            *position.Position
	DoTkn               *token.Token
	Stmt                Vertex
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtDo) Accept(visitor Visitor) {
	visitor.StmtDo(node)
}

func (node *StmtDo) GetPosition() *position.Position {
	return node.Position
}

// StmtEcho node
type StmtEcho struct {
	Position      *position.Position
	EchoTkn       *token.Token
	Exprs         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtEcho) Accept(visitor Visitor) {
	visitor.StmtEcho(node)
}

func (node *StmtEcho) GetPosition() *position.Position {
	return node.Position
}

// StmtElse node
type StmtElse struct {
	Position *position.Position
	ElseTkn  *token.Token
	ColonTkn *token.Token
	Stmt     Vertex
}

func (node *StmtElse) Accept(visitor Visitor) {
	visitor.StmtElse(node)
}

func (node *StmtElse) GetPosition() *position.Position {
	return node.Position
}

// StmtElseIf node
type StmtElseIf struct {
	Position            *position.Position
	ElseIfTkn           *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
}

func (node *StmtElseIf) Accept(visitor Visitor) {
	visitor.StmtElseIf(node)
}

func (node *StmtElseIf) GetPosition() *position.Position {
	return node.Position
}

// StmtExpression node
type StmtExpression struct {
	Position     *position.Position
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtExpression) Accept(visitor Visitor) {
	visitor.StmtExpression(node)
}

func (node *StmtExpression) GetPosition() *position.Position {
	return node.Position
}

// StmtFinally node
type StmtFinally struct {
	Position             *position.Position
	FinallyTkn           *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtFinally) Accept(visitor Visitor) {
	visitor.StmtFinally(node)
}

func (node *StmtFinally) GetPosition() *position.Position {
	return node.Position
}

// StmtFor node
type StmtFor struct {
	Position            *position.Position
	ForTkn              *token.Token
	OpenParenthesisTkn  *token.Token
	Init                []Vertex
	InitSeparatorTkns   []*token.Token
	InitSemiColonTkn    *token.Token
	Cond                []Vertex
	CondSeparatorTkns   []*token.Token
	CondSemiColonTkn    *token.Token
	Loop                []Vertex
	LoopSeparatorTkns   []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForTkn           *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtFor) Accept(visitor Visitor) {
	visitor.StmtFor(node)
}

func (node *StmtFor) GetPosition() *position.Position {
	return node.Position
}

// StmtForeach node
type StmtForeach struct {
	Position            *position.Position
	ForeachTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	AsTkn               *token.Token
	Key                 Vertex
	DoubleArrowTkn      *token.Token
	AmpersandTkn        *token.Token
	Var                 Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForeachTkn       *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtForeach) Accept(visitor Visitor) {
	visitor.StmtForeach(node)
}

func (node *StmtForeach) GetPosition() *position.Position {
	return node.Position
}

// StmtFunction node
type StmtFunction struct {
	Position             *position.Position
	AttrGroups           []Vertex
	FunctionTkn          *token.Token
	AmpersandTkn         *token.Token
	Name                 Vertex
	OpenParenthesisTkn   *token.Token
	Params               []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	ReturnType           Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtFunction) Accept(visitor Visitor) {
	visitor.StmtFunction(node)
}

func (node *StmtFunction) GetPosition() *position.Position {
	return node.Position
}

// StmtGlobal node
type StmtGlobal struct {
	Position      *position.Position
	GlobalTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtGlobal) Accept(visitor Visitor) {
	visitor.StmtGlobal(node)
}

func (node *StmtGlobal) GetPosition() *position.Position {
	return node.Position
}

// StmtGoto node
type StmtGoto struct {
	Position     *position.Position
	GotoTkn      *token.Token
	Label        Vertex
	SemiColonTkn *token.Token
}

func (node *StmtGoto) Accept(visitor Visitor) {
	visitor.StmtGoto(node)
}

func (node *StmtGoto) GetPosition() *position.Position {
	return node.Position
}

// StmtHaltCompiler node
type StmtHaltCompiler struct {
	Position            *position.Position
	HaltCompilerTkn     *token.Token
	OpenParenthesisTkn  *token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtHaltCompiler) Accept(visitor Visitor) {
	visitor.StmtHaltCompiler(node)
}

func (node *StmtHaltCompiler) GetPosition() *position.Position {
	return node.Position
}

// StmtIf node
type StmtIf struct {
	Position            *position.Position
	IfTkn               *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	ElseIf              []Vertex
	Else                Vertex
	EndIfTkn            *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtIf) Accept(visitor Visitor) {
	visitor.StmtIf(node)
}

func (node *StmtIf) GetPosition() *position.Position {
	return node.Position
}

// StmtInlineHtml node
type StmtInlineHtml struct {
	Position      *position.Position
	InlineHtmlTkn *token.Token
	Value         []byte
}

func (node *StmtInlineHtml) Accept(visitor Visitor) {
	visitor.StmtInlineHtml(node)
}

func (node *StmtInlineHtml) GetPosition() *position.Position {
	return node.Position
}

// StmtEnum node holds an enum declaration, optionally backed by a scalar type.
type StmtEnum struct {
	Position                *position.Position
	AttrGroups              []Vertex
	EnumTkn                 *token.Token
	Name                    Vertex
	ColonTkn                *token.Token
	Type                    Vertex
	ImplementsTkn           *token.Token
	Implements              []Vertex
	ImplementsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn     *token.Token
	Stmts                   []Vertex
	CloseCurlyBracketTkn    *token.Token
}

func (node *StmtEnum) Accept(visitor Visitor) {
	visitor.StmtEnum(node)
}

func (node *StmtEnum) GetPosition() *position.Position {
	return node.Position
}

// StmtEnumCase node holds one `case` of an enum, with a value when the enum is
// backed.
type StmtEnumCase struct {
	Position     *position.Position
	AttrGroups   []Vertex
	CaseTkn      *token.Token
	Name         Vertex
	EqualTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtEnumCase) Accept(visitor Visitor) {
	visitor.StmtEnumCase(node)
}

func (node *StmtEnumCase) GetPosition() *position.Position {
	return node.Position
}

// StmtInterface node
type StmtInterface struct {
	Position             *position.Position
	AttrGroups           []Vertex
	InterfaceTkn         *token.Token
	Name                 Vertex
	ExtendsTkn           *token.Token
	Extends              []Vertex
	ExtendsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtInterface) Accept(visitor Visitor) {
	visitor.StmtInterface(node)
}

func (node *StmtInterface) GetPosition() *position.Position {
	return node.Position
}

// StmtLabel node
type StmtLabel struct {
	Position *position.Position
	Name     Vertex
	ColonTkn *token.Token
}

func (node *StmtLabel) Accept(visitor Visitor) {
	visitor.StmtLabel(node)
}

func (node *StmtLabel) GetPosition() *position.Position {
	return node.Position
}

// StmtNamespace node
type StmtNamespace struct {
	Position             *position.Position
	NsTkn                *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

func (node *StmtNamespace) Accept(visitor Visitor) {
	visitor.StmtNamespace(node)
}

func (node *StmtNamespace) GetPosition() *position.Position {
	return node.Position
}

// StmtNop node
type StmtNop struct {
	Position     *position.Position
	SemiColonTkn *token.Token
}

func (node *StmtNop) Accept(visitor Visitor) {
	visitor.StmtNop(node)
}

func (node *StmtNop) GetPosition() *position.Position {
	return node.Position
}

// StmtProperty node
type StmtProperty struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *StmtProperty) Accept(visitor Visitor) {
	visitor.StmtProperty(node)
}

func (node *StmtProperty) GetPosition() *position.Position {
	return node.Position
}

// StmtPropertyList node
type StmtPropertyList struct {
	Position      *position.Position
	AttrGroups    []Vertex
	Modifiers     []Vertex
	Type          Vertex
	Props         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtPropertyList) Accept(visitor Visitor) {
	visitor.StmtPropertyList(node)
}

func (node *StmtPropertyList) GetPosition() *position.Position {
	return node.Position
}

// StmtReturn node
type StmtReturn struct {
	Position     *position.Position
	ReturnTkn    *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtReturn) Accept(visitor Visitor) {
	visitor.StmtReturn(node)
}

func (node *StmtReturn) GetPosition() *position.Position {
	return node.Position
}

// StmtStatic node
type StmtStatic struct {
	Position      *position.Position
	StaticTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtStatic) Accept(visitor Visitor) {
	visitor.StmtStatic(node)
}

func (node *StmtStatic) GetPosition() *position.Position {
	return node.Position
}

// StmtStaticVar node
type StmtStaticVar struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *StmtStaticVar) Accept(visitor Visitor) {
	visitor.StmtStaticVar(node)
}

func (node *StmtStaticVar) GetPosition() *position.Position {
	return node.Position
}

// StmtStmtList node
type StmtStmtList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtStmtList) Accept(visitor Visitor) {
	visitor.StmtStmtList(node)
}

func (node *StmtStmtList) GetPosition() *position.Position {
	return node.Position
}

// StmtSwitch node
type StmtSwitch struct {
	Position             *position.Position
	SwitchTkn            *token.Token
	OpenParenthesisTkn   *token.Token
	Cond                 Vertex
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	OpenCurlyBracketTkn  *token.Token
	CaseSeparatorTkn     *token.Token
	Cases                []Vertex
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
}

func (node *StmtSwitch) Accept(visitor Visitor) {
	visitor.StmtSwitch(node)
}

func (node *StmtSwitch) GetPosition() *position.Position {
	return node.Position
}

// StmtThrow node
type StmtThrow struct {
	Position     *position.Position
	ThrowTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (node *StmtThrow) Accept(visitor Visitor) {
	visitor.StmtThrow(node)
}

func (node *StmtThrow) GetPosition() *position.Position {
	return node.Position
}

// ExprThrow node is a `throw` used in expression position, such as
// `$value ?? throw new InvalidArgumentException()`.
type ExprThrow struct {
	Position *position.Position
	ThrowTkn *token.Token
	Expr     Vertex
}

func (node *ExprThrow) Accept(visitor Visitor) {
	visitor.ExprThrow(node)
}

func (node *ExprThrow) GetPosition() *position.Position {
	return node.Position
}

// StmtTrait node
type StmtTrait struct {
	Position             *position.Position
	AttrGroups           []Vertex
	TraitTkn             *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *StmtTrait) Accept(visitor Visitor) {
	visitor.StmtTrait(node)
}

func (node *StmtTrait) GetPosition() *position.Position {
	return node.Position
}

// StmtTraitUse node
type StmtTraitUse struct {
	Position             *position.Position
	UseTkn               *token.Token
	Traits               []Vertex
	SeparatorTkns        []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

func (node *StmtTraitUse) Accept(visitor Visitor) {
	visitor.StmtTraitUse(node)
}

func (node *StmtTraitUse) GetPosition() *position.Position {
	return node.Position
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	AsTkn          *token.Token
	Modifier       Vertex
	Alias          Vertex
	SemiColonTkn   *token.Token
}

func (node *StmtTraitUseAlias) Accept(visitor Visitor) {
	visitor.StmtTraitUseAlias(node)
}

func (node *StmtTraitUseAlias) GetPosition() *position.Position {
	return node.Position
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	InsteadofTkn   *token.Token
	Insteadof      []Vertex
	SeparatorTkns  []*token.Token
	SemiColonTkn   *token.Token
}

func (node *StmtTraitUsePrecedence) Accept(visitor Visitor) {
	visitor.StmtTraitUsePrecedence(node)
}

func (node *StmtTraitUsePrecedence) GetPosition() *position.Position {
	return node.Position
}

// StmtTry node
type StmtTry struct {
	Position             *position.Position
	TryTkn               *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	Catches              []Vertex
	Finally              Vertex
}

func (node *StmtTry) Accept(visitor Visitor) {
	visitor.StmtTry(node)
}

func (node *StmtTry) GetPosition() *position.Position {
	return node.Position
}

// StmtUnset node
type StmtUnset struct {
	Position            *position.Position
	UnsetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtUnset) Accept(visitor Visitor) {
	visitor.StmtUnset(node)
}

func (node *StmtUnset) GetPosition() *position.Position {
	return node.Position
}

// StmtUseList node
type StmtUseList struct {
	Position      *position.Position
	UseTkn        *token.Token
	Type          Vertex
	Uses          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (node *StmtUseList) Accept(visitor Visitor) {
	visitor.StmtUse(node)
}

func (node *StmtUseList) GetPosition() *position.Position {
	return node.Position
}

// StmtGroupUseList node
type StmtGroupUseList struct {
	Position              *position.Position
	UseTkn                *token.Token
	Type                  Vertex
	LeadingNsSeparatorTkn *token.Token
	Prefix                Vertex
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	Uses                  []Vertex
	SeparatorTkns         []*token.Token
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

func (node *StmtGroupUseList) Accept(visitor Visitor) {
	visitor.StmtGroupUse(node)
}

func (node *StmtGroupUseList) GetPosition() *position.Position {
	return node.Position
}

// StmtUse node
type StmtUse struct {
	Position       *position.Position
	Type           Vertex
	NsSeparatorTkn *token.Token
	Use            Vertex
	AsTkn          *token.Token
	Alias          Vertex
}

func (node *StmtUse) Accept(visitor Visitor) {
	visitor.StmtUseDeclaration(node)
}

func (node *StmtUse) GetPosition() *position.Position {
	return node.Position
}

// StmtWhile node
type StmtWhile struct {
	Position            *position.Position
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndWhileTkn         *token.Token
	SemiColonTkn        *token.Token
}

func (node *StmtWhile) Accept(visitor Visitor) {
	visitor.StmtWhile(node)
}

func (node *StmtWhile) GetPosition() *position.Position {
	return node.Position
}

// ExprArray node
type ExprArray struct {
	Position        *position.Position
	ArrayTkn        *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

func (node *ExprArray) Accept(visitor Visitor) {
	visitor.ExprArray(node)
}

func (node *ExprArray) GetPosition() *position.Position {
	return node.Position
}

// ExprArrayDimFetch node
type ExprArrayDimFetch struct {
	Position        *position.Position
	Var             Vertex
	OpenBracketTkn  *token.Token
	Dim             Vertex
	CloseBracketTkn *token.Token
}

func (node *ExprArrayDimFetch) Accept(visitor Visitor) {
	visitor.ExprArrayDimFetch(node)
}

func (node *ExprArrayDimFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprArrayItem node
type ExprArrayItem struct {
	Position       *position.Position
	EllipsisTkn    *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	AmpersandTkn   *token.Token
	Val            Vertex
}

func (node *ExprArrayItem) Accept(visitor Visitor) {
	visitor.ExprArrayItem(node)
}

func (node *ExprArrayItem) GetPosition() *position.Position {
	return node.Position
}

// ExprArrowFunction node
type ExprArrowFunction struct {
	Position            *position.Position
	StaticTkn           *token.Token
	FnTkn               *token.Token
	AmpersandTkn        *token.Token
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	DoubleArrowTkn      *token.Token
	Expr                Vertex
}

func (node *ExprArrowFunction) Accept(visitor Visitor) {
	visitor.ExprArrowFunction(node)
}

func (node *ExprArrowFunction) GetPosition() *position.Position {
	return node.Position
}

// ExprBitwiseNot node
type ExprBitwiseNot struct {
	Position *position.Position
	TildaTkn *token.Token
	Expr     Vertex
}

func (node *ExprBitwiseNot) Accept(visitor Visitor) {
	visitor.ExprBitwiseNot(node)
}

func (node *ExprBitwiseNot) GetPosition() *position.Position {
	return node.Position
}

// ExprBooleanNot node
type ExprBooleanNot struct {
	Position       *position.Position
	ExclamationTkn *token.Token
	Expr           Vertex
}

func (node *ExprBooleanNot) Accept(visitor Visitor) {
	visitor.ExprBooleanNot(node)
}

func (node *ExprBooleanNot) GetPosition() *position.Position {
	return node.Position
}

type ExprBrackets struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (node *ExprBrackets) Accept(visitor Visitor) {
	visitor.ExprBrackets(node)
}

func (node *ExprBrackets) GetPosition() *position.Position {
	return node.Position
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Const          Vertex
}

func (node *ExprClassConstFetch) Accept(visitor Visitor) {
	visitor.ExprClassConstFetch(node)
}

func (node *ExprClassConstFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprClone node
type ExprClone struct {
	Position *position.Position
	CloneTkn *token.Token
	Expr     Vertex
}

func (node *ExprClone) Accept(visitor Visitor) {
	visitor.ExprClone(node)
}

func (node *ExprClone) GetPosition() *position.Position {
	return node.Position
}

// ExprClosure node
type ExprClosure struct {
	Position               *position.Position
	StaticTkn              *token.Token
	FunctionTkn            *token.Token
	AmpersandTkn           *token.Token
	OpenParenthesisTkn     *token.Token
	Params                 []Vertex
	SeparatorTkns          []*token.Token
	CloseParenthesisTkn    *token.Token
	UseTkn                 *token.Token
	UseOpenParenthesisTkn  *token.Token
	Uses                   []Vertex
	UseSeparatorTkns       []*token.Token
	UseCloseParenthesisTkn *token.Token
	ColonTkn               *token.Token
	ReturnType             Vertex
	OpenCurlyBracketTkn    *token.Token
	Stmts                  []Vertex
	CloseCurlyBracketTkn   *token.Token
}

func (node *ExprClosure) Accept(visitor Visitor) {
	visitor.ExprClosure(node)
}

func (node *ExprClosure) GetPosition() *position.Position {
	return node.Position
}

// ExprClosureUse node
type ExprClosureUse struct {
	Position     *position.Position
	AmpersandTkn *token.Token
	Var          Vertex
}

func (node *ExprClosureUse) Accept(visitor Visitor) {
	visitor.ExprClosureUse(node)
}

func (node *ExprClosureUse) GetPosition() *position.Position {
	return node.Position
}

// ExprConstFetch node
type ExprConstFetch struct {
	Position *position.Position
	Const    Vertex
}

func (node *ExprConstFetch) Accept(visitor Visitor) {
	visitor.ExprConstFetch(node)
}

func (node *ExprConstFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprEmpty node
type ExprEmpty struct {
	Position            *position.Position
	EmptyTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (node *ExprEmpty) Accept(visitor Visitor) {
	visitor.ExprEmpty(node)
}

func (node *ExprEmpty) GetPosition() *position.Position {
	return node.Position
}

// ExprErrorSuppress node
type ExprErrorSuppress struct {
	Position *position.Position
	AtTkn    *token.Token
	Expr     Vertex
}

func (node *ExprErrorSuppress) Accept(visitor Visitor) {
	visitor.ExprErrorSuppress(node)
}

func (node *ExprErrorSuppress) GetPosition() *position.Position {
	return node.Position
}

// ExprEval node
type ExprEval struct {
	Position            *position.Position
	EvalTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (node *ExprEval) Accept(visitor Visitor) {
	visitor.ExprEval(node)
}

func (node *ExprEval) GetPosition() *position.Position {
	return node.Position
}

// ExprExit node
type ExprExit struct {
	Position            *position.Position
	ExitTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (node *ExprExit) Accept(visitor Visitor) {
	visitor.ExprExit(node)
}

func (node *ExprExit) GetPosition() *position.Position {
	return node.Position
}

// ExprFunctionCall node
type ExprFunctionCall struct {
	Position            *position.Position
	Function            Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (node *ExprFunctionCall) Accept(visitor Visitor) {
	visitor.ExprFunctionCall(node)
}

func (node *ExprFunctionCall) GetPosition() *position.Position {
	return node.Position
}

// ExprInclude node
type ExprInclude struct {
	Position   *position.Position
	IncludeTkn *token.Token
	Expr       Vertex
}

func (node *ExprInclude) Accept(visitor Visitor) {
	visitor.ExprInclude(node)
}

func (node *ExprInclude) GetPosition() *position.Position {
	return node.Position
}

// ExprIncludeOnce node
type ExprIncludeOnce struct {
	Position       *position.Position
	IncludeOnceTkn *token.Token
	Expr           Vertex
}

func (node *ExprIncludeOnce) Accept(visitor Visitor) {
	visitor.ExprIncludeOnce(node)
}

func (node *ExprIncludeOnce) GetPosition() *position.Position {
	return node.Position
}

// ExprInstanceOf node
type ExprInstanceOf struct {
	Position      *position.Position
	Expr          Vertex
	InstanceOfTkn *token.Token
	Class         Vertex
}

func (node *ExprInstanceOf) Accept(visitor Visitor) {
	visitor.ExprInstanceOf(node)
}

func (node *ExprInstanceOf) GetPosition() *position.Position {
	return node.Position
}

// ExprIsset node
type ExprIsset struct {
	Position            *position.Position
	IssetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (node *ExprIsset) Accept(visitor Visitor) {
	visitor.ExprIsset(node)
}

func (node *ExprIsset) GetPosition() *position.Position {
	return node.Position
}

// ExprList node
type ExprList struct {
	Position        *position.Position
	ListTkn         *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

func (node *ExprList) Accept(visitor Visitor) {
	visitor.ExprList(node)
}

func (node *ExprList) GetPosition() *position.Position {
	return node.Position
}

// ExprMethodCall node
type ExprMethodCall struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

func (node *ExprMethodCall) Accept(visitor Visitor) {
	visitor.ExprMethodCall(node)
}

func (node *ExprMethodCall) GetPosition() *position.Position {
	return node.Position
}

// ExprNullsafeMethodCall node is a `$object?->method()` call, which short-circuits
// to null when the object is null.
type ExprNullsafeMethodCall struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

func (node *ExprNullsafeMethodCall) Accept(visitor Visitor) {
	visitor.ExprNullsafeMethodCall(node)
}

func (node *ExprNullsafeMethodCall) GetPosition() *position.Position {
	return node.Position
}

// ExprNullsafePropertyFetch node is a `$object?->property` access, which
// short-circuits to null when the object is null.
type ExprNullsafePropertyFetch struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Prop                 Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *ExprNullsafePropertyFetch) Accept(visitor Visitor) {
	visitor.ExprNullsafePropertyFetch(node)
}

func (node *ExprNullsafePropertyFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprNew node
type ExprNew struct {
	Position            *position.Position
	NewTkn              *token.Token
	Class               Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (node *ExprNew) Accept(visitor Visitor) {
	visitor.ExprNew(node)
}

func (node *ExprNew) GetPosition() *position.Position {
	return node.Position
}

// ExprPostDec node
type ExprPostDec struct {
	Position *position.Position
	Var      Vertex
	DecTkn   *token.Token
}

func (node *ExprPostDec) Accept(visitor Visitor) {
	visitor.ExprPostDec(node)
}

func (node *ExprPostDec) GetPosition() *position.Position {
	return node.Position
}

// ExprPostInc node
type ExprPostInc struct {
	Position *position.Position
	Var      Vertex
	IncTkn   *token.Token
}

func (node *ExprPostInc) Accept(visitor Visitor) {
	visitor.ExprPostInc(node)
}

func (node *ExprPostInc) GetPosition() *position.Position {
	return node.Position
}

// ExprPreDec node
type ExprPreDec struct {
	Position *position.Position
	DecTkn   *token.Token
	Var      Vertex
}

func (node *ExprPreDec) Accept(visitor Visitor) {
	visitor.ExprPreDec(node)
}

func (node *ExprPreDec) GetPosition() *position.Position {
	return node.Position
}

// ExprPreInc node
type ExprPreInc struct {
	Position *position.Position
	IncTkn   *token.Token
	Var      Vertex
}

func (node *ExprPreInc) Accept(visitor Visitor) {
	visitor.ExprPreInc(node)
}

func (node *ExprPreInc) GetPosition() *position.Position {
	return node.Position
}

// ExprPrint node
type ExprPrint struct {
	Position *position.Position
	PrintTkn *token.Token
	Expr     Vertex
}

func (node *ExprPrint) Accept(visitor Visitor) {
	visitor.ExprPrint(node)
}

func (node *ExprPrint) GetPosition() *position.Position {
	return node.Position
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Prop                 Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *ExprPropertyFetch) Accept(visitor Visitor) {
	visitor.ExprPropertyFetch(node)
}

func (node *ExprPropertyFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprRequire node
type ExprRequire struct {
	Position   *position.Position
	RequireTkn *token.Token
	Expr       Vertex
}

func (node *ExprRequire) Accept(visitor Visitor) {
	visitor.ExprRequire(node)
}

func (node *ExprRequire) GetPosition() *position.Position {
	return node.Position
}

// ExprRequireOnce node
type ExprRequireOnce struct {
	Position       *position.Position
	RequireOnceTkn *token.Token
	Expr           Vertex
}

func (node *ExprRequireOnce) Accept(visitor Visitor) {
	visitor.ExprRequireOnce(node)
}

func (node *ExprRequireOnce) GetPosition() *position.Position {
	return node.Position
}

// ExprShellExec node
type ExprShellExec struct {
	Position         *position.Position
	OpenBacktickTkn  *token.Token
	Parts            []Vertex
	CloseBacktickTkn *token.Token
}

func (node *ExprShellExec) Accept(visitor Visitor) {
	visitor.ExprShellExec(node)
}

func (node *ExprShellExec) GetPosition() *position.Position {
	return node.Position
}

// ExprStaticCall node
type ExprStaticCall struct {
	Position             *position.Position
	Class                Vertex
	DoubleColonTkn       *token.Token
	OpenCurlyBracketTkn  *token.Token
	Call                 Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

func (node *ExprStaticCall) Accept(visitor Visitor) {
	visitor.ExprStaticCall(node)
}

func (node *ExprStaticCall) GetPosition() *position.Position {
	return node.Position
}

// ExprStaticPropertyFetch node
type ExprStaticPropertyFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Prop           Vertex
}

func (node *ExprStaticPropertyFetch) Accept(visitor Visitor) {
	visitor.ExprStaticPropertyFetch(node)
}

func (node *ExprStaticPropertyFetch) GetPosition() *position.Position {
	return node.Position
}

// ExprTernary node
type ExprTernary struct {
	Position    *position.Position
	Cond        Vertex
	QuestionTkn *token.Token
	IfTrue      Vertex
	ColonTkn    *token.Token
	IfFalse     Vertex
}

func (node *ExprTernary) Accept(visitor Visitor) {
	visitor.ExprTernary(node)
}

func (node *ExprTernary) GetPosition() *position.Position {
	return node.Position
}

// ExprMatch node
type ExprMatch struct {
	Position             *position.Position
	MatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Expr                 Vertex
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Arms                 []Vertex
	SeparatorTkns        []*token.Token
	CloseCurlyBracketTkn *token.Token
}

func (node *ExprMatch) Accept(visitor Visitor) {
	visitor.ExprMatch(node)
}

func (node *ExprMatch) GetPosition() *position.Position {
	return node.Position
}

// MatchArm node
type MatchArm struct {
	Position       *position.Position
	DefaultTkn     *token.Token
	Exprs          []Vertex
	SeparatorTkns  []*token.Token
	DoubleArrowTkn *token.Token
	ReturnExpr     Vertex
}

func (node *MatchArm) Accept(visitor Visitor) {
	visitor.MatchArm(node)
}

func (node *MatchArm) GetPosition() *position.Position {
	return node.Position
}

// ExprUnaryMinus node
type ExprUnaryMinus struct {
	Position *position.Position
	MinusTkn *token.Token
	Expr     Vertex
}

func (node *ExprUnaryMinus) Accept(visitor Visitor) {
	visitor.ExprUnaryMinus(node)
}

func (node *ExprUnaryMinus) GetPosition() *position.Position {
	return node.Position
}

// ExprUnaryPlus node
type ExprUnaryPlus struct {
	Position *position.Position
	PlusTkn  *token.Token
	Expr     Vertex
}

func (node *ExprUnaryPlus) Accept(visitor Visitor) {
	visitor.ExprUnaryPlus(node)
}

func (node *ExprUnaryPlus) GetPosition() *position.Position {
	return node.Position
}

// ExprVariable node
type ExprVariable struct {
	Position             *position.Position
	DollarTkn            *token.Token
	OpenCurlyBracketTkn  *token.Token
	Name                 Vertex
	CloseCurlyBracketTkn *token.Token
}

func (node *ExprVariable) Accept(visitor Visitor) {
	visitor.ExprVariable(node)
}

func (node *ExprVariable) GetPosition() *position.Position {
	return node.Position
}

// ExprYield node
type ExprYield struct {
	Position       *position.Position
	YieldTkn       *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	Val            Vertex
}

func (node *ExprYield) Accept(visitor Visitor) {
	visitor.ExprYield(node)
}

func (node *ExprYield) GetPosition() *position.Position {
	return node.Position
}

// ExprYieldFrom node
type ExprYieldFrom struct {
	Position     *position.Position
	YieldFromTkn *token.Token
	Expr         Vertex
}

func (node *ExprYieldFrom) Accept(visitor Visitor) {
	visitor.ExprYieldFrom(node)
}

func (node *ExprYieldFrom) GetPosition() *position.Position {
	return node.Position
}

// ExprCastArray node
type ExprCastArray struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastArray) Accept(visitor Visitor) {
	visitor.ExprCastArray(node)
}

func (node *ExprCastArray) GetPosition() *position.Position {
	return node.Position
}

// ExprCastBool node
type ExprCastBool struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastBool) Accept(visitor Visitor) {
	visitor.ExprCastBool(node)
}

func (node *ExprCastBool) GetPosition() *position.Position {
	return node.Position
}

// ExprCastDouble node
type ExprCastDouble struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastDouble) Accept(visitor Visitor) {
	visitor.ExprCastDouble(node)
}

func (node *ExprCastDouble) GetPosition() *position.Position {
	return node.Position
}

// ExprCastInt node
type ExprCastInt struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastInt) Accept(visitor Visitor) {
	visitor.ExprCastInt(node)
}

func (node *ExprCastInt) GetPosition() *position.Position {
	return node.Position
}

// ExprCastObject node
type ExprCastObject struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastObject) Accept(visitor Visitor) {
	visitor.ExprCastObject(node)
}

func (node *ExprCastObject) GetPosition() *position.Position {
	return node.Position
}

// ExprCastString node
type ExprCastString struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastString) Accept(visitor Visitor) {
	visitor.ExprCastString(node)
}

func (node *ExprCastString) GetPosition() *position.Position {
	return node.Position
}

// ExprCastUnset node
type ExprCastUnset struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (node *ExprCastUnset) Accept(visitor Visitor) {
	visitor.ExprCastUnset(node)
}

func (node *ExprCastUnset) GetPosition() *position.Position {
	return node.Position
}

// ExprAssign node
type ExprAssign struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssign) Accept(visitor Visitor) {
	visitor.ExprAssign(node)
}

func (node *ExprAssign) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignReference node
type ExprAssignReference struct {
	Position     *position.Position
	Var          Vertex
	EqualTkn     *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

func (node *ExprAssignReference) Accept(visitor Visitor) {
	visitor.ExprAssignReference(node)
}

func (node *ExprAssignReference) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignBitwiseAnd node
type ExprAssignBitwiseAnd struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignBitwiseAnd) Accept(visitor Visitor) {
	visitor.ExprAssignBitwiseAnd(node)
}

func (node *ExprAssignBitwiseAnd) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignBitwiseOr node
type ExprAssignBitwiseOr struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignBitwiseOr) Accept(visitor Visitor) {
	visitor.ExprAssignBitwiseOr(node)
}

func (node *ExprAssignBitwiseOr) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignBitwiseXor node
type ExprAssignBitwiseXor struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignBitwiseXor) Accept(visitor Visitor) {
	visitor.ExprAssignBitwiseXor(node)
}

func (node *ExprAssignBitwiseXor) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignCoalesce node
type ExprAssignCoalesce struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignCoalesce) Accept(visitor Visitor) {
	visitor.ExprAssignCoalesce(node)
}

func (node *ExprAssignCoalesce) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignConcat node
type ExprAssignConcat struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignConcat) Accept(visitor Visitor) {
	visitor.ExprAssignConcat(node)
}

func (node *ExprAssignConcat) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignDiv node
type ExprAssignDiv struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignDiv) Accept(visitor Visitor) {
	visitor.ExprAssignDiv(node)
}

func (node *ExprAssignDiv) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignMinus node
type ExprAssignMinus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignMinus) Accept(visitor Visitor) {
	visitor.ExprAssignMinus(node)
}

func (node *ExprAssignMinus) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignMod node
type ExprAssignMod struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignMod) Accept(visitor Visitor) {
	visitor.ExprAssignMod(node)
}

func (node *ExprAssignMod) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignMul node
type ExprAssignMul struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignMul) Accept(visitor Visitor) {
	visitor.ExprAssignMul(node)
}

func (node *ExprAssignMul) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignPlus node
type ExprAssignPlus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignPlus) Accept(visitor Visitor) {
	visitor.ExprAssignPlus(node)
}

func (node *ExprAssignPlus) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignPow node
type ExprAssignPow struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignPow) Accept(visitor Visitor) {
	visitor.ExprAssignPow(node)
}

func (node *ExprAssignPow) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignShiftLeft node
type ExprAssignShiftLeft struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignShiftLeft) Accept(visitor Visitor) {
	visitor.ExprAssignShiftLeft(node)
}

func (node *ExprAssignShiftLeft) GetPosition() *position.Position {
	return node.Position
}

// ExprAssignShiftRight node
type ExprAssignShiftRight struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (node *ExprAssignShiftRight) Accept(visitor Visitor) {
	visitor.ExprAssignShiftRight(node)
}

func (node *ExprAssignShiftRight) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryBitwiseAnd node
type ExprBinaryBitwiseAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryBitwiseAnd) Accept(visitor Visitor) {
	visitor.ExprBinaryBitwiseAnd(node)
}

func (node *ExprBinaryBitwiseAnd) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryBitwiseOr node
type ExprBinaryBitwiseOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryBitwiseOr) Accept(visitor Visitor) {
	visitor.ExprBinaryBitwiseOr(node)
}

func (node *ExprBinaryBitwiseOr) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryBitwiseXor node
type ExprBinaryBitwiseXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryBitwiseXor) Accept(visitor Visitor) {
	visitor.ExprBinaryBitwiseXor(node)
}

func (node *ExprBinaryBitwiseXor) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryBooleanAnd node
type ExprBinaryBooleanAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryBooleanAnd) Accept(visitor Visitor) {
	visitor.ExprBinaryBooleanAnd(node)
}

func (node *ExprBinaryBooleanAnd) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryBooleanOr node
type ExprBinaryBooleanOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryBooleanOr) Accept(visitor Visitor) {
	visitor.ExprBinaryBooleanOr(node)
}

func (node *ExprBinaryBooleanOr) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryCoalesce node
type ExprBinaryCoalesce struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryCoalesce) Accept(visitor Visitor) {
	visitor.ExprBinaryCoalesce(node)
}

func (node *ExprBinaryCoalesce) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryConcat node
type ExprBinaryConcat struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryConcat) Accept(visitor Visitor) {
	visitor.ExprBinaryConcat(node)
}

func (node *ExprBinaryConcat) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryDiv node
type ExprBinaryDiv struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryDiv) Accept(visitor Visitor) {
	visitor.ExprBinaryDiv(node)
}

func (node *ExprBinaryDiv) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryEqual node
type ExprBinaryEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryEqual) Accept(visitor Visitor) {
	visitor.ExprBinaryEqual(node)
}

func (node *ExprBinaryEqual) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryGreater node
type ExprBinaryGreater struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryGreater) Accept(visitor Visitor) {
	visitor.ExprBinaryGreater(node)
}

func (node *ExprBinaryGreater) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryGreaterOrEqual node
type ExprBinaryGreaterOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryGreaterOrEqual) Accept(visitor Visitor) {
	visitor.ExprBinaryGreaterOrEqual(node)
}

func (node *ExprBinaryGreaterOrEqual) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryIdentical node
type ExprBinaryIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryIdentical) Accept(visitor Visitor) {
	visitor.ExprBinaryIdentical(node)
}

func (node *ExprBinaryIdentical) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryLogicalAnd node
type ExprBinaryLogicalAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryLogicalAnd) Accept(visitor Visitor) {
	visitor.ExprBinaryLogicalAnd(node)
}

func (node *ExprBinaryLogicalAnd) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryLogicalOr node
type ExprBinaryLogicalOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryLogicalOr) Accept(visitor Visitor) {
	visitor.ExprBinaryLogicalOr(node)
}

func (node *ExprBinaryLogicalOr) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryLogicalXor node
type ExprBinaryLogicalXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryLogicalXor) Accept(visitor Visitor) {
	visitor.ExprBinaryLogicalXor(node)
}

func (node *ExprBinaryLogicalXor) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryMinus node
type ExprBinaryMinus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryMinus) Accept(visitor Visitor) {
	visitor.ExprBinaryMinus(node)
}

func (node *ExprBinaryMinus) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryMod node
type ExprBinaryMod struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryMod) Accept(visitor Visitor) {
	visitor.ExprBinaryMod(node)
}

func (node *ExprBinaryMod) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryMul node
type ExprBinaryMul struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryMul) Accept(visitor Visitor) {
	visitor.ExprBinaryMul(node)
}

func (node *ExprBinaryMul) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryNotEqual node
type ExprBinaryNotEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryNotEqual) Accept(visitor Visitor) {
	visitor.ExprBinaryNotEqual(node)
}

func (node *ExprBinaryNotEqual) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryNotIdentical node
type ExprBinaryNotIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryNotIdentical) Accept(visitor Visitor) {
	visitor.ExprBinaryNotIdentical(node)
}

func (node *ExprBinaryNotIdentical) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryPlus node
type ExprBinaryPlus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryPlus) Accept(visitor Visitor) {
	visitor.ExprBinaryPlus(node)
}

func (node *ExprBinaryPlus) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryPow node
type ExprBinaryPow struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryPow) Accept(visitor Visitor) {
	visitor.ExprBinaryPow(node)
}

func (node *ExprBinaryPow) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryShiftLeft node
type ExprBinaryShiftLeft struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryShiftLeft) Accept(visitor Visitor) {
	visitor.ExprBinaryShiftLeft(node)
}

func (node *ExprBinaryShiftLeft) GetPosition() *position.Position {
	return node.Position
}

// ExprBinaryShiftRight node
type ExprBinaryShiftRight struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinaryShiftRight) Accept(visitor Visitor) {
	visitor.ExprBinaryShiftRight(node)
}

func (node *ExprBinaryShiftRight) GetPosition() *position.Position {
	return node.Position
}

// ExprBinarySmaller node
type ExprBinarySmaller struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinarySmaller) Accept(visitor Visitor) {
	visitor.ExprBinarySmaller(node)
}

func (node *ExprBinarySmaller) GetPosition() *position.Position {
	return node.Position
}

// ExprBinarySmallerOrEqual node
type ExprBinarySmallerOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinarySmallerOrEqual) Accept(visitor Visitor) {
	visitor.ExprBinarySmallerOrEqual(node)
}

func (node *ExprBinarySmallerOrEqual) GetPosition() *position.Position {
	return node.Position
}

// ExprBinarySpaceship node
type ExprBinarySpaceship struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (node *ExprBinarySpaceship) Accept(visitor Visitor) {
	visitor.ExprBinarySpaceship(node)
}

func (node *ExprBinarySpaceship) GetPosition() *position.Position {
	return node.Position
}

type Name struct {
	Position      *position.Position
	Parts         []Vertex
	SeparatorTkns []*token.Token
}

func (node *Name) Accept(visitor Visitor) {
	visitor.NameName(node)
}

func (node *Name) GetPosition() *position.Position {
	return node.Position
}

type NameFullyQualified struct {
	Position       *position.Position
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

func (node *NameFullyQualified) Accept(visitor Visitor) {
	visitor.NameFullyQualified(node)
}

func (node *NameFullyQualified) GetPosition() *position.Position {
	return node.Position
}

type NameRelative struct {
	Position       *position.Position
	NsTkn          *token.Token
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

func (node *NameRelative) Accept(visitor Visitor) {
	visitor.NameRelative(node)
}

func (node *NameRelative) GetPosition() *position.Position {
	return node.Position
}

type NamePart struct {
	Position  *position.Position
	StringTkn *token.Token
	Value     []byte
}

func (node *NamePart) Accept(visitor Visitor) {
	visitor.NameNamePart(node)
}

func (node *NamePart) GetPosition() *position.Position {
	return node.Position
}
