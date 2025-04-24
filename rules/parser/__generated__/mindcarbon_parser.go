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
		"", "", "", "", "", "", "'''", "", "", "'/'", "'*'", "'+'", "'-'", "'^'",
		"'='", "':'", "','", "'('", "')'", "'{'", "'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "COMPUTATION", "ID", "QUOTED_STRING", "ESC_SEQ", "UNICODE", "SINGLE_QUOTE",
		"INT", "FLOAT", "DIV", "MUL", "ADD", "MIN", "EXP", "EQUALS", "COLUMN",
		"COMMA", "LPAREN", "RPAREN", "LCURLY", "RCURLY", "SM", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "computationDef", "paramDef", "block", "bstat",
		"assignment", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 3, 2, 41,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 52, 8,
		3, 1, 4, 1, 4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 68, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 3, 7, 94, 8, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 99, 8, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 114, 8, 7, 10, 7, 12, 7, 117,
		9, 7, 1, 7, 0, 1, 14, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 3, 1, 0, 7, 8, 1,
		0, 9, 10, 1, 0, 11, 12, 128, 0, 19, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 29,
		1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 67, 1, 0, 0, 0, 12,
		69, 1, 0, 0, 0, 14, 102, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0,
		0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1,
		1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 23, 3, 14, 7, 0, 23, 24, 5, 21, 0,
		0, 24, 28, 1, 0, 0, 0, 25, 28, 3, 4, 2, 0, 26, 28, 3, 12, 6, 0, 27, 22,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 3, 1, 0, 0, 0,
		29, 30, 5, 1, 0, 0, 30, 31, 5, 2, 0, 0, 31, 40, 5, 17, 0, 0, 32, 37, 3,
		6, 3, 0, 33, 34, 5, 16, 0, 0, 34, 36, 3, 6, 3, 0, 35, 33, 1, 0, 0, 0, 36,
		39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 41, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 40, 32, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42,
		1, 0, 0, 0, 42, 43, 5, 18, 0, 0, 43, 44, 3, 8, 4, 0, 44, 5, 1, 0, 0, 0,
		45, 46, 5, 2, 0, 0, 46, 47, 5, 15, 0, 0, 47, 51, 5, 2, 0, 0, 48, 49, 5,
		17, 0, 0, 49, 50, 5, 2, 0, 0, 50, 52, 5, 18, 0, 0, 51, 48, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 7, 1, 0, 0, 0, 53, 57, 5, 19, 0, 0, 54, 56, 3,
		10, 5, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 3, 14,
		7, 0, 61, 62, 5, 20, 0, 0, 62, 9, 1, 0, 0, 0, 63, 68, 3, 12, 6, 0, 64,
		65, 3, 14, 7, 0, 65, 66, 5, 21, 0, 0, 66, 68, 1, 0, 0, 0, 67, 63, 1, 0,
		0, 0, 67, 64, 1, 0, 0, 0, 68, 11, 1, 0, 0, 0, 69, 70, 5, 2, 0, 0, 70, 71,
		5, 14, 0, 0, 71, 72, 3, 14, 7, 0, 72, 73, 5, 21, 0, 0, 73, 13, 1, 0, 0,
		0, 74, 75, 6, 7, -1, 0, 75, 76, 5, 17, 0, 0, 76, 77, 3, 14, 7, 0, 77, 78,
		5, 18, 0, 0, 78, 103, 1, 0, 0, 0, 79, 80, 5, 2, 0, 0, 80, 93, 5, 17, 0,
		0, 81, 82, 5, 2, 0, 0, 82, 83, 5, 15, 0, 0, 83, 90, 3, 14, 7, 0, 84, 85,
		5, 16, 0, 0, 85, 86, 5, 2, 0, 0, 86, 87, 5, 15, 0, 0, 87, 89, 3, 14, 7,
		0, 88, 84, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 81, 1, 0, 0, 0,
		93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 103, 5, 18, 0, 0, 96, 103,
		5, 2, 0, 0, 97, 99, 5, 12, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 100, 1, 0, 0, 0, 100, 103, 7, 0, 0, 0, 101, 103, 5, 3, 0, 0, 102, 74,
		1, 0, 0, 0, 102, 79, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 98, 1, 0, 0,
		0, 102, 101, 1, 0, 0, 0, 103, 115, 1, 0, 0, 0, 104, 105, 10, 7, 0, 0, 105,
		106, 5, 13, 0, 0, 106, 114, 3, 14, 7, 8, 107, 108, 10, 6, 0, 0, 108, 109,
		7, 1, 0, 0, 109, 114, 3, 14, 7, 7, 110, 111, 10, 5, 0, 0, 111, 112, 7,
		2, 0, 0, 112, 114, 3, 14, 7, 6, 113, 104, 1, 0, 0, 0, 113, 107, 1, 0, 0,
		0, 113, 110, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 15, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 13, 19, 27,
		37, 40, 51, 57, 67, 90, 93, 98, 102, 113, 115,
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
	mindcarbonParserSINGLE_QUOTE  = 6
	mindcarbonParserINT           = 7
	mindcarbonParserFLOAT         = 8
	mindcarbonParserDIV           = 9
	mindcarbonParserMUL           = 10
	mindcarbonParserADD           = 11
	mindcarbonParserMIN           = 12
	mindcarbonParserEXP           = 13
	mindcarbonParserEQUALS        = 14
	mindcarbonParserCOLUMN        = 15
	mindcarbonParserCOMMA         = 16
	mindcarbonParserLPAREN        = 17
	mindcarbonParserRPAREN        = 18
	mindcarbonParserLCURLY        = 19
	mindcarbonParserRCURLY        = 20
	mindcarbonParserSM            = 21
	mindcarbonParserCOMMENT       = 22
	mindcarbonParserWS            = 23
)

// mindcarbonParser rules.
const (
	mindcarbonParserRULE_program        = 0
	mindcarbonParserRULE_statement      = 1
	mindcarbonParserRULE_computationDef = 2
	mindcarbonParserRULE_paramDef       = 3
	mindcarbonParserRULE_block          = 4
	mindcarbonParserRULE_bstat          = 5
	mindcarbonParserRULE_assignment     = 6
	mindcarbonParserRULE_expr           = 7
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135566) != 0 {
		{
			p.SetState(16)
			p.Statement()
		}

		p.SetState(21)
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
	Expr() IExprContext
	SM() antlr.TerminalNode
	ComputationDef() IComputationDefContext
	Assignment() IAssignmentContext

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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.expr(0)
		}
		{
			p.SetState(23)
			p.Match(mindcarbonParserSM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.ComputationDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Assignment()
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
	Block() IBlockContext
	ID() antlr.TerminalNode
	AllParamDef() []IParamDefContext
	ParamDef(i int) IParamDefContext
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

func (s *ComputationDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(mindcarbonParserCOMPUTATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ComputationDefContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(mindcarbonParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindcarbonParserID {
		{
			p.SetState(32)
			p.ParamDef()
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == mindcarbonParserCOMMA {
			{
				p.SetState(33)
				p.Match(mindcarbonParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(34)
				p.ParamDef()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(42)
		p.Match(mindcarbonParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Block()
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
		p.SetState(45)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ParamDefContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(mindcarbonParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)

		var _m = p.Match(mindcarbonParserID)

		localctx.(*ParamDefContext).type_ = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindcarbonParserLPAREN {
		{
			p.SetState(48)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)

			var _m = p.Match(mindcarbonParserID)

			localctx.(*ParamDefContext).unit = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	Expr() IExprContext
	RCURLY() antlr.TerminalNode
	AllBstat() []IBstatContext
	Bstat(i int) IBstatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLCURLY, 0)
}

func (s *BlockContext) Expr() IExprContext {
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

func (s *BlockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRCURLY, 0)
}

func (s *BlockContext) AllBstat() []IBstatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBstatContext); ok {
			len++
		}
	}

	tst := make([]IBstatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBstatContext); ok {
			tst[i] = t.(IBstatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Bstat(i int) IBstatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBstatContext); ok {
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

	return t.(IBstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mindcarbonParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(mindcarbonParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(54)
				p.Bstat()
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.expr(0)
	}
	{
		p.SetState(61)
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

// IBstatContext is an interface to support dynamic dispatch.
type IBstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Expr() IExprContext
	SM() antlr.TerminalNode

	// IsBstatContext differentiates from other interfaces.
	IsBstatContext()
}

type BstatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBstatContext() *BstatContext {
	var p = new(BstatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_bstat
	return p
}

func InitEmptyBstatContext(p *BstatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindcarbonParserRULE_bstat
}

func (*BstatContext) IsBstatContext() {}

func NewBstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BstatContext {
	var p = new(BstatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindcarbonParserRULE_bstat

	return p
}

func (s *BstatContext) GetParser() antlr.Parser { return s.parser }

func (s *BstatContext) Assignment() IAssignmentContext {
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

func (s *BstatContext) Expr() IExprContext {
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

func (s *BstatContext) SM() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserSM, 0)
}

func (s *BstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterBstat(s)
	}
}

func (s *BstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitBstat(s)
	}
}

func (s *BstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitBstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindcarbonParser) Bstat() (localctx IBstatContext) {
	localctx = NewBstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mindcarbonParserRULE_bstat)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.expr(0)
		}
		{
			p.SetState(65)
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
	p.EnterRule(localctx, 12, mindcarbonParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(mindcarbonParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(mindcarbonParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.expr(0)
	}
	{
		p.SetState(72)
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

type IdResolutionContext struct {
	ExprContext
}

func NewIdResolutionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdResolutionContext {
	var p = new(IdResolutionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdResolutionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdResolutionContext) ID() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, 0)
}

func (s *IdResolutionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterIdResolution(s)
	}
}

func (s *IdResolutionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitIdResolution(s)
	}
}

func (s *IdResolutionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitIdResolution(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulOrDivOpContext struct {
	ExprContext
	op antlr.Token
}

func NewMulOrDivOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulOrDivOpContext {
	var p = new(MulOrDivOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulOrDivOpContext) GetOp() antlr.Token { return s.op }

func (s *MulOrDivOpContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulOrDivOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOrDivOpContext) AllExpr() []IExprContext {
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

func (s *MulOrDivOpContext) Expr(i int) IExprContext {
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

func (s *MulOrDivOpContext) MUL() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMUL, 0)
}

func (s *MulOrDivOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserDIV, 0)
}

func (s *MulOrDivOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterMulOrDivOp(s)
	}
}

func (s *MulOrDivOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitMulOrDivOp(s)
	}
}

func (s *MulOrDivOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitMulOrDivOp(s)

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
	fnName  antlr.Token
	prmName antlr.Token
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetFnName() antlr.Token { return s.fnName }

func (s *FunctionCallContext) GetPrmName() antlr.Token { return s.prmName }

func (s *FunctionCallContext) SetFnName(v antlr.Token) { s.fnName = v }

func (s *FunctionCallContext) SetPrmName(v antlr.Token) { s.prmName = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserRPAREN, 0)
}

func (s *FunctionCallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(mindcarbonParserID)
}

func (s *FunctionCallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(mindcarbonParserID, i)
}

func (s *FunctionCallContext) AllCOLUMN() []antlr.TerminalNode {
	return s.GetTokens(mindcarbonParserCOLUMN)
}

func (s *FunctionCallContext) COLUMN(i int) antlr.TerminalNode {
	return s.GetToken(mindcarbonParserCOLUMN, i)
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

type AddOrMinOpContext struct {
	ExprContext
	op antlr.Token
}

func NewAddOrMinOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddOrMinOpContext {
	var p = new(AddOrMinOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddOrMinOpContext) GetOp() antlr.Token { return s.op }

func (s *AddOrMinOpContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddOrMinOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOrMinOpContext) AllExpr() []IExprContext {
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

func (s *AddOrMinOpContext) Expr(i int) IExprContext {
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

func (s *AddOrMinOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserADD, 0)
}

func (s *AddOrMinOpContext) MIN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMIN, 0)
}

func (s *AddOrMinOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterAddOrMinOp(s)
	}
}

func (s *AddOrMinOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitAddOrMinOp(s)
	}
}

func (s *AddOrMinOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitAddOrMinOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOpContext struct {
	ExprContext
}

func NewExpOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOpContext {
	var p = new(ExpOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOpContext) AllExpr() []IExprContext {
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

func (s *ExpOpContext) Expr(i int) IExprContext {
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

func (s *ExpOpContext) EXP() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserEXP, 0)
}

func (s *ExpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterExpOp(s)
	}
}

func (s *ExpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitExpOp(s)
	}
}

func (s *ExpOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitExpOp(s)

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

type NumberConstantContext struct {
	ExprContext
}

func NewNumberConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberConstantContext {
	var p = new(NumberConstantContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumberConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserFLOAT, 0)
}

func (s *NumberConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserINT, 0)
}

func (s *NumberConstantContext) MIN() antlr.TerminalNode {
	return s.GetToken(mindcarbonParserMIN, 0)
}

func (s *NumberConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.EnterNumberConstant(s)
	}
}

func (s *NumberConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindcarbonListener); ok {
		listenerT.ExitNumberConstant(s)
	}
}

func (s *NumberConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindcarbonVisitor:
		return t.VisitNumberConstant(s)

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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, mindcarbonParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
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
			p.SetState(75)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.expr(0)
		}
		{
			p.SetState(77)
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
			p.SetState(79)

			var _m = p.Match(mindcarbonParserID)

			localctx.(*FunctionCallContext).fnName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(mindcarbonParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == mindcarbonParserID {
			{
				p.SetState(81)

				var _m = p.Match(mindcarbonParserID)

				localctx.(*FunctionCallContext).prmName = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(82)
				p.Match(mindcarbonParserCOLUMN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(83)
				p.expr(0)
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == mindcarbonParserCOMMA {
				{
					p.SetState(84)
					p.Match(mindcarbonParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)

					var _m = p.Match(mindcarbonParserID)

					localctx.(*FunctionCallContext).prmName = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.Match(mindcarbonParserCOLUMN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(87)
					p.expr(0)
				}

				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(95)
			p.Match(mindcarbonParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIdResolutionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(mindcarbonParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewNumberConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == mindcarbonParserMIN {
			{
				p.SetState(97)
				p.Match(mindcarbonParserMIN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(100)
			_la = p.GetTokenStream().LA(1)

			if !(_la == mindcarbonParserINT || _la == mindcarbonParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		localctx = NewStringConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
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
	p.SetState(115)
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
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(mindcarbonParserEXP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expr(8)
				}

			case 2:
				localctx = NewMulOrDivOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(108)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulOrDivOpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindcarbonParserDIV || _la == mindcarbonParserMUL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulOrDivOpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(109)
					p.expr(7)
				}

			case 3:
				localctx = NewAddOrMinOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mindcarbonParserRULE_expr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(111)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddOrMinOpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindcarbonParserADD || _la == mindcarbonParserMIN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddOrMinOpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(112)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(117)
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
	case 7:
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
