package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/mindiply/mindcarbon/rules/evaltree"
	parser "github.com/mindiply/mindcarbon/rules/parser/__generated__"
	"strconv"
)

func DmlToEvalTree(dml string) (evaltree.Evaluator, error) {
	textStream := antlr.NewInputStream(dml)
	lexer := parser.NewmindcarbonLexer(textStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	dmlParser := parser.NewmindcarbonParser(stream)
	listener := NewEvalTreeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, dmlParser.Program())
	if len(listener.errors) > 0 {
		return nil, fmt.Errorf("errors during parsing: %v", listener.errors)
	}
	return listener.ProgramEvaluator, nil
}

type EvalTreeListener struct {
	parser.BasemindcarbonListener
	evaluators       map[interface{}]evaltree.Evaluator
	ProgramEvaluator evaltree.Evaluator
	errors           []error
}

func NewEvalTreeListener() *EvalTreeListener {
	return &EvalTreeListener{
		evaluators: make(map[interface{}]evaltree.Evaluator),
		errors:     make([]error, 0),
	}
}

func (s *EvalTreeListener) ExitProgram(ctx *parser.ProgramContext) {
	statements := make([]evaltree.Evaluator, 0)

	for _, statement := range ctx.AllStatement() {
		statementEvaluator, ok := s.evaluators[statement]
		if !ok {
			return
		}
		statements = append(statements, statementEvaluator)
	}
	program := evaltree.NewProgramEvaluator(statements)
	s.evaluators[ctx] = program
	s.ProgramEvaluator = program
}

func (s *EvalTreeListener) ExitStatement(ctx *parser.StatementContext) {
	if ctx.Assignment() != nil {
		assignmentEvaluator, ok := s.evaluators[ctx.Assignment()]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("assignment expression required"))
			return
		}
		s.evaluators[ctx] = assignmentEvaluator
	} else if ctx.ComputationDef() != nil {
		computationDefEvaluator, ok := s.evaluators[ctx.ComputationDef()]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("computation definition required"))
			return
		}
		s.evaluators[ctx] = computationDefEvaluator
	} else if ctx.Expr() != nil {
		exprEvaluator, ok := s.evaluators[ctx.Expr()]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("expression required"))
			return
		}
		s.evaluators[ctx] = exprEvaluator
	} else {
		s.errors = append(s.errors, fmt.Errorf("unknown statement type"))
	}
}

func (s *EvalTreeListener) ExitComputationDef(ctx *parser.ComputationDefContext) {

}

func (s *EvalTreeListener) ExitFloatConstant(ctx *parser.FloatConstantContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		s.errors = append(s.errors, err)
		return
	}
	s.evaluators[ctx] = evaltree.NewNumberEvaluator(val)
}

func (s *EvalTreeListener) ExitIntConstant(ctx *parser.IntConstantContext) {
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		s.errors = append(s.errors, err)
		return
	}
	s.evaluators[ctx] = evaltree.NewNumberEvaluator(float64(val))
}

func (s *EvalTreeListener) ExitStringConstant(ctx *parser.StringConstantContext) {
	strToken := ctx.QUOTED_STRING()
	if strToken == nil {
		s.errors = append(s.errors, fmt.Errorf("string constant required"))
		return
	}
	str := strToken.GetText()
	str = str[1 : len(str)-1] // Remove the quotes
	s.evaluators[ctx] = evaltree.NewStringEvaluator(str)
}

func (s *EvalTreeListener) ExitIdentifier(ctx *parser.IdentifierContext) {
	strCtx := ctx.ID()
	if strCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("identifier required"))
		return
	}
	s.evaluators[ctx] = evaltree.NewStringEvaluator(strCtx.GetText())
}

func (s *EvalTreeListener) ExitParamDef(ctx *parser.ParamDefContext) {

}

func (s *EvalTreeListener) ExitAssignment(ctx *parser.AssignmentContext) {

}

func (s *EvalTreeListener) ExitDivision(ctx *parser.DivisionContext) {

}
