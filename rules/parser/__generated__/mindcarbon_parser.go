// Code generated from mindcarbon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindcarbon

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type mindcarbonParser struct {
	*antlr.BaseParser
}

var MindcarbonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mindcarbonParserInit() {
	staticData := &MindcarbonParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'''", "", "", "'/'", "'*'", "'+'", "'-'",
		"'^'", "'='", "':'", "','", "'('", "')'", "'{'", "'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "COMPUTATION", "ID", "QUOTED_STRING", "ESC_SEQ", "UNICODE", "HEX",
		"SINGLE_QUOTE", "INT", "FLOAT", "DIV", "MUL", "ADD", "MIN", "EXP", "EQUALS",
		"COLUMN", "COMMA", "LPAREN", "RPAREN", "LCURLY", "RCURLY", "SM", "COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "computationDef", "paramDef", "assignment",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 113, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 24, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 32, 8, 2, 10, 2, 12, 2, 35, 9, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 42, 8, 2, 10, 2, 12, 2, 45, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 56, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 73, 8, 5, 10,
		5, 12, 5, 76, 9, 5, 3, 5, 78, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 83, 8, 5, 1,
		5, 1, 5, 3, 5, 87, 8, 5, 1, 5, 1, 5, 3, 5, 91, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 108, 8, 5, 10, 5, 12, 5, 111, 9, 5, 1, 5, 0, 1, 10, 6, 0, 2, 4, 6,
		8, 10, 0, 0, 127, 0, 15, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 25, 1, 0, 0,
		0, 6, 49, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 90, 1, 0, 0, 0, 12, 14, 3,
		2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15,
		16, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 24, 3, 8, 4,
		0, 19, 24, 3, 4, 2, 0, 20, 21, 3, 10, 5, 0, 21, 22, 5, 22, 0, 0, 22, 24,
		1, 0, 0, 0, 23, 18, 1, 0, 0, 0, 23, 19, 1, 0, 0, 0, 23, 20, 1, 0, 0, 0,
		24, 3, 1, 0, 0, 0, 25, 26, 5, 1, 0, 0, 26, 27, 5, 2, 0, 0, 27, 36, 5, 18,
		0, 0, 28, 33, 3, 6, 3, 0, 29, 30, 5, 17, 0, 0, 30, 32, 3, 6, 3, 0, 31,
		29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 28, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 19, 0, 0, 39, 43, 5, 20, 0,
		0, 40, 42, 3, 8, 4, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		46, 47, 3, 10, 5, 0, 47, 48, 5, 21, 0, 0, 48, 5, 1, 0, 0, 0, 49, 50, 5,
		2, 0, 0, 50, 51, 5, 16, 0, 0, 51, 55, 5, 2, 0, 0, 52, 53, 5, 18, 0, 0,
		53, 54, 5, 2, 0, 0, 54, 56, 5, 19, 0, 0, 55, 52, 1, 0, 0, 0, 55, 56, 1,
		0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 2, 0, 0, 58, 59, 5, 15, 0, 0, 59,
		60, 3, 10, 5, 0, 60, 61, 5, 22, 0, 0, 61, 9, 1, 0, 0, 0, 62, 63, 6, 5,
		-1, 0, 63, 64, 5, 18, 0, 0, 64, 65, 3, 10, 5, 0, 65, 66, 5, 19, 0, 0, 66,
		91, 1, 0, 0, 0, 67, 68, 5, 2, 0, 0, 68, 77, 5, 18, 0, 0, 69, 74, 3, 10,
		5, 0, 70, 71, 5, 17, 0, 0, 71, 73, 3, 10, 5, 0, 72, 70, 1, 0, 0, 0, 73,
		76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 78, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 91, 5, 19, 0, 0, 80, 91, 5, 2, 0, 0, 81, 83, 5, 13, 0,
		0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 91,
		5, 9, 0, 0, 85, 87, 5, 13, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 91, 5, 8, 0, 0, 89, 91, 5, 3, 0, 0, 90, 62, 1,
		0, 0, 0, 90, 67, 1, 0, 0, 0, 90, 80, 1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90,
		86, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 109, 1, 0, 0, 0, 92, 93, 10, 10,
		0, 0, 93, 94, 5, 14, 0, 0, 94, 108, 3, 10, 5, 11, 95, 96, 10, 9, 0, 0,
		96, 97, 5, 10, 0, 0, 97, 108, 3, 10, 5, 10, 98, 99, 10, 8, 0, 0, 99, 100,
		5, 11, 0, 0, 100, 108, 3, 10, 5, 9, 101, 102, 10, 7, 0, 0, 102, 103, 5,
		12, 0, 0, 103, 108, 3, 10, 5, 8, 104, 105, 10, 6, 0, 0, 105, 106, 5, 13,
		0, 0, 106, 108, 3, 10, 5, 7, 107, 92, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0,
		107, 98, 1, 0, 0, 0, 107, 101, 1, 0, 0, 0, 107, 104, 1, 0, 0, 0, 108, 111,
		1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 11, 1, 0,
		0, 0, 111, 109, 1, 0, 0, 0, 13, 15, 23, 33, 36, 43, 55, 74, 77, 82, 86,
		90, 107, 109,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// mindcarbonParserInit initializes any static state used to implement mindcarbonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewmindcarbonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MindcarbonParserInit() {
	staticData := &MindcarbonParserStaticData
	staticData.once.Do(mindcarbonParserInit)
}

// NewmindcarbonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewmindcarbonParser(input antlr.TokenStream) *mindcarbonParser {
	MindcarbonParserInit()
	this := new(mindcarbonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MindcarbonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "mindcarbon.g4"

	return this
}

// mindcarbonParser tokens.
const (
	mindcarbonParserEOF           = antlr.TokenEOF
	mindcarbonParserCOMPUTATION   = 1
	mindcarbonParserID            = 2
	mindcarbonParserQUOTED_STRING = 3
	mindcarbonParserESC_SEQ       = 4
	mindcarbonParserUNICODE       = 5
	mindcarbonParserHEX           = 6
	mindcarbonParserSINGLE_QUOTE  = 7
	mindcarbonParserINT           = 8
	mindcarbonParserFLOAT         = 9
	mindcarbonParserDIV           = 10
	mindcarbonParserMUL           = 11
	mindcarbonParserADD           = 12
	mindcarbonParserMIN           = 13
	mindcarbonParserEXP           = 14
	mindcarbonParserEQUALS        = 15
	mindcarbonParserCOLUMN        = 16
	mindcarbonParserCOMMA         = 17
	mindcarbonParserLPAREN        = 18
	mindcarbonParserRPAREN        = 19
	mindcarbonParserLCURLY        = 20
	mindcarbonParserRCURLY        = 21
	mindcarbonParserSM            = 22
	mindcarbonParserCOMMENT       = 23
	mindcarbonParserWS            = 24
)

// mindcarbonParser rules.
const (
	mindcarbonParserRULE_program        = 0
	mindcarbonParserRULE_statement      = 1
	mindcarbonParserRULE_computationDef = 2
	mindcarbonParserRULE_paramDef       = 3
	mindcarbonParserRULE_assignment     = 4
	mindcarbonParserRULE_expr           = 5
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mindcarbonParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&271118) != 0 {
		{
			p.SetState(12)
			p.Statement()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ComputationDef() IComputationDefContext
	Expr() IExprContext
	SM() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ComputationDef() IComputationDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComputationDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComputationDefContext)
}

func (s *StatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatementContext) SM() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserSM, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mindcarbonParserRULE_statement)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.ComputationDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
			p.expr(0)
		}
		{
			p.SetState(21)
			p.Match(mindcarbonParserSM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComputationDefContext is an interface to support dynamic dispatch.
type IComputationDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	COMPUTATION() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LCURLY() antlr.TerminalNode
	Expr() IExprContext
	RCURLY() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllParamDef() []IParamDefContext
	ParamDef(i int) IParamDefContext
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsComputationDefContext differentiates from other interfaces.
	IsComputationDefContext()
}

type ComputationDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyComputationDefContext() *ComputationDefContext {
	var p = new(ComputationDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_computationDef
	return p
}

func InitEmptyComputationDefContext(p *ComputationDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_computationDef
}

func (*ComputationDefContext) IsComputationDefContext() {}

func NewComputationDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputationDefContext {
	var p = new(ComputationDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_computationDef

	return p
}

func (s *ComputationDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ComputationDefContext) GetName() antlr.Token { return s.name }

func (s *ComputationDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ComputationDefContext) COMPUTATION() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserCOMPUTATION, 0)
}

func (s *ComputationDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLPAREN, 0)
}

func (s *ComputationDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRPAREN, 0)
}

func (s *ComputationDefContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLCURLY, 0)
}

func (s *ComputationDefContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComputationDefContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRCURLY, 0)
}

func (s *ComputationDefContext) ID() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, 0)
}

func (s *ComputationDefContext) AllParamDef() []IParamDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDefContext); ok {
			len++
		}
	}

	tst := make([]IParamDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDefContext); ok {
			tst[i] = t.(IParamDefContext)
			i++
		}
	}

	return tst
}

func (s *ComputationDefContext) ParamDef(i int) IParamDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamDefContext)
}

func (s *ComputationDefContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ComputationDefContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ComputationDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mindcarbonParserCOMMA)
}

func (s *ComputationDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mindcarbonParserCOMMA, i)
}

func (s *ComputationDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputationDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComputationDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterComputationDef(s)
	}
}

func (s *ComputationDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitComputationDef(s)
	}
}

func (s *ComputationDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitComputationDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) ComputationDef() (localctx IComputationDefContext) {
	localctx = NewComputationDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mindcarbonParserRULE_computationDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(mindcarbonParserCOMPUTATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ComputationDefContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(mindcarbonParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindcarbonParserID {
		{
			p.SetState(28)
			p.ParamDef()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == mindcarbonParserCOMMA {
			{
				p.SetState(29)
				p.Match(mindcarbonParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(30)
				p.ParamDef()
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(38)
		p.Match(mindcarbonParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(mindcarbonParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(40)
				p.Assignment()
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.expr(0)
	}
	{
		p.SetState(47)
		p.Match(mindcarbonParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamDefContext is an interface to support dynamic dispatch.
type IParamDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// GetUnit returns the unit token.
	GetUnit() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// SetUnit sets the unit token.
	SetUnit(antlr.Token)

	// Getter signatures
	COLUMN() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsParamDefContext differentiates from other interfaces.
	IsParamDefContext()
}

type ParamDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	type_  antlr.Token
	unit   antlr.Token
}

func NewEmptyParamDefContext() *ParamDefContext {
	var p = new(ParamDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_paramDef
	return p
}

func InitEmptyParamDefContext(p *ParamDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_paramDef
}

func (*ParamDefContext) IsParamDefContext() {}

func NewParamDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDefContext {
	var p = new(ParamDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_paramDef

	return p
}

func (s *ParamDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDefContext) GetName() antlr.Token { return s.name }

func (s *ParamDefContext) GetType_() antlr.Token { return s.type_ }

func (s *ParamDefContext) GetUnit() antlr.Token { return s.unit }

func (s *ParamDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ParamDefContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *ParamDefContext) SetUnit(v antlr.Token) { s.unit = v }

func (s *ParamDefContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserCOLUMN, 0)
}

func (s *ParamDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(mindcarbonParserID)
}

func (s *ParamDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, i)
}

func (s *ParamDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLPAREN, 0)
}

func (s *ParamDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRPAREN, 0)
}

func (s *ParamDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterParamDef(s)
	}
}

func (s *ParamDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitParamDef(s)
	}
}

func (s *ParamDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitParamDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) ParamDef() (localctx IParamDefContext) {
	localctx = NewParamDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mindcarbonParserRULE_paramDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ParamDefContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(mindcarbonParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ParamDefContext).type_ = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindcarbonParserLPAREN {
		{
			p.SetState(52)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)

			var _m = p.Match(mindcarbonParserID)

			localctx.(*ParamDefContext).unit = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(mindcarbonParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext
	SM() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, 0)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserEQUALS, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) SM() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserSM, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mindcarbonParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(mindcarbonParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(mindcarbonParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.expr(0)
	}
	{
		p.SetState(60)
		p.Match(mindcarbonParserSM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DivisionContext struct {
	ExprContext
}

func NewDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivisionContext {
	var p = new(DivisionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DivisionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivisionContext) DIV() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserDIV, 0)
}

func (s *DivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterDivision(s)
	}
}

func (s *DivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitDivision(s)
	}
}

func (s *DivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentiationContext struct {
	ExprContext
}

func NewExponentiationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentiationContext {
	var p = new(ExponentiationContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExponentiationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentiationContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExponentiationContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExponentiationContext) EXP() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserEXP, 0)
}

func (s *ExponentiationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterExponentiation(s)
	}
}

func (s *ExponentiationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitExponentiation(s)
	}
}

func (s *ExponentiationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitExponentiation(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntConstantContext struct {
	ExprContext
}

func NewIntConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntConstantContext {
	var p = new(IntConstantContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserINT, 0)
}

func (s *IntConstantContext) MIN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMIN, 0)
}

func (s *IntConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterIntConstant(s)
	}
}

func (s *IntConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitIntConstant(s)
	}
}

func (s *IntConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitIntConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringConstantContext struct {
	ExprContext
}

func NewStringConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringConstantContext {
	var p = new(StringConstantContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringConstantContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserQUOTED_STRING, 0)
}

func (s *StringConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterStringConstant(s)
	}
}

func (s *StringConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitStringConstant(s)
	}
}

func (s *StringConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitStringConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
	fnName antlr.Token
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetFnName() antlr.Token { return s.fnName }

func (s *FunctionCallContext) SetFnName(v antlr.Token) { s.fnName = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRPAREN, 0)
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mindcarbonParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mindcarbonParserCOMMA, i)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatConstantContext struct {
	ExprContext
}

func NewFloatConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatConstantContext {
	var p = new(FloatConstantContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserFLOAT, 0)
}

func (s *FloatConstantContext) MIN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMIN, 0)
}

func (s *FloatConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterFloatConstant(s)
	}
}

func (s *FloatConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitFloatConstant(s)
	}
}

func (s *FloatConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitFloatConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubtractionContext struct {
	ExprContext
}

func NewSubtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractionContext {
	var p = new(SubtractionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SubtractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SubtractionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubtractionContext) MIN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMIN, 0)
}

func (s *SubtractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterSubtraction(s)
	}
}

func (s *SubtractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitSubtraction(s)
	}
}

func (s *SubtractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitSubtraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationContext struct {
	ExprContext
}

func NewMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationContext {
	var p = new(MultiplicationContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicationContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplicationContext) MUL() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMUL, 0)
}

func (s *MultiplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterMultiplication(s)
	}
}

func (s *MultiplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitMultiplication(s)
	}
}

func (s *MultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupingContext struct {
	ExprContext
}

func NewGroupingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupingContext {
	var p = new(GroupingContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLPAREN, 0)
}

func (s *GroupingContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupingContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRPAREN, 0)
}

func (s *GroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (s *GroupingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitGrouping(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditionContext struct {
	ExprContext
}

func NewAdditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionContext {
	var p = new(AdditionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AdditionContext) ADD() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserADD, 0)
}

func (s *AdditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterAddition(s)
	}
}

func (s *AdditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitAddition(s)
	}
}

func (s *AdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *mindcarbonParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, mindcarbonParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGroupingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(63)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.expr(0)
		}
		{
			p.SetState(65)
			p.Match(mindcarbonParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)

			var _m = p.Match(mindcarbonParserID)

			localctx.(*FunctionCallContext).fnName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&271116) != 0 {
			{
				p.SetState(69)
				p.expr(0)
			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == mindcarbonParserCOMMA {
				{
					p.SetState(70)
					p.Match(mindcarbonParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.expr(0)
				}

				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(79)
			p.Match(mindcarbonParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(mindcarbonParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFloatConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == mindcarbonParserMIN {
			{
				p.SetState(81)
				p.Match(mindcarbonParserMIN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(84)
			p.Match(mindcarbonParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIntConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == mindcarbonParserMIN {
			{
				p.SetState(85)
				p.Match(mindcarbonParserMIN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(88)
			p.Match(mindcarbonParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Match(mindcarbonParserQUOTED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentiationContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					p.Match(mindcarbonParserEXP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.expr(11)
				}

			case 2:
				localctx = NewDivisionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(mindcarbonParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.expr(10)
				}

			case 3:
				localctx = NewMultiplicationContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.Match(mindcarbonParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.expr(9)
				}

			case 4:
				localctx = NewAdditionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(102)
					p.Match(mindcarbonParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(103)
					p.expr(8)
				}

			case 5:
				localctx = NewSubtractionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(mindcarbonParserMIN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *mindcarbonParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *mindcarbonParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
