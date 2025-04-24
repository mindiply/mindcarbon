package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/mindiply/mindcarbon/rules/evaltree"
	parser "github.com/mindiply/mindcarbon/rules/parser/__generated__"
	"strconv"
	"strings"
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
	paramDefs        map[interface{}]evaltree.ParameterDefinition
	ProgramEvaluator evaltree.Evaluator
	errors           []error
}

func NewEvalTreeListener() *EvalTreeListener {
	return &EvalTreeListener{
		evaluators: make(map[interface{}]evaltree.Evaluator),
		errors:     make([]error, 0),
		paramDefs:  make(map[interface{}]evaltree.ParameterDefinition),
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
	aCtx := ctx.Assignment()
	cCtx := ctx.ComputationDef()
	eCtx := ctx.Expr()
	if aCtx != nil {
		assignmentEvaluator, ok := s.evaluators[aCtx]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("assignment expression required"))
			return
		}
		s.evaluators[ctx] = assignmentEvaluator
	} else if cCtx != nil {
		computationDefEvaluator, ok := s.evaluators[cCtx]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("computation definition required"))
			return
		}
		s.evaluators[ctx] = computationDefEvaluator
	} else if eCtx != nil {
		exprEvaluator, ok := s.evaluators[eCtx]
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
	nameCtx := ctx.GetName()
	if nameCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("computation name required"))
		return
	}
	name := nameCtx.GetText()
	name = strings.TrimSpace(name)
	if name == "" {
		s.errors = append(s.errors, fmt.Errorf("computation name is a non empty id made of a first letter followed by letters, numnbers or underscore"))
		return
	}

	allPrmCtxs := ctx.AllParamDef()
	params := make(map[string]evaltree.ParameterDefinition)
	for _, prmCtx := range allPrmCtxs {
		prmDef, ok := s.paramDefs[prmCtx]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("parameter definition required"))
			return
		}
		params[prmDef.Name] = prmDef
	}

	blockCtx := ctx.Block()
	if blockCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("block required for calculation definition %s", name))
		return
	}
	blockEvaluator, ok := s.evaluators[blockCtx]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("block evaluator required for calculation definition %s", name))
		return
	}
	s.evaluators[ctx] = evaltree.NewComputationDefEvaluator(name, params, blockEvaluator)
}

func (s *EvalTreeListener) ExitParamDef(ctx *parser.ParamDefContext) {
	nameCtx := ctx.GetName()
	if nameCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("parameter name required in function definition"))
		return
	}
	qualityCtx := ctx.GetType_()
	if qualityCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("parameter type required in function definition"))
		return
	}
	name := nameCtx.GetText()
	name = strings.TrimSpace(name)
	if name == "" {
		s.errors = append(s.errors, fmt.Errorf("parameter name required in function definition"))
		return
	}
	quality := qualityCtx.GetText()
	quality = strings.TrimSpace(quality)
	if quality == "" {
		s.errors = append(s.errors, fmt.Errorf("parameter type required in function definition"))
		return
	}

	unitCtx := ctx.GetUnit()
	unit := ""
	if unitCtx != nil {
		unit = unitCtx.GetText()
		unit = strings.TrimSpace(unit)
	}
	s.paramDefs[ctx] = evaltree.NewParameterDefinition(
		name,
		quality,
		unit)
}

func (s *EvalTreeListener) ExitBlock(ctx *parser.BlockContext) {
	additionalStatements := ctx.AllBstat()

	statements := make([]evaltree.Evaluator, len(additionalStatements)+1)
	for i, bstat := range additionalStatements {
		statementEvaluator, ok := s.evaluators[bstat]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("statement expression required"))
			return
		}
		statements[i] = statementEvaluator
	}
	finalExprCtx := ctx.Expr()
	if finalExprCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("expression required"))
		return
	}
	finalExpr, ok := s.evaluators[finalExprCtx]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("return expression required"))
		return
	}
	statements[len(statements)-1] = finalExpr
	s.evaluators[ctx] = evaltree.NewBlockEvaluator(statements)
}

func (s *EvalTreeListener) ExitBstat(ctx *parser.BstatContext) {
	eCtx := ctx.Expr()
	if eCtx != nil {
		exprEvaluator, ok := s.evaluators[eCtx]
		if !ok {
			s.errors = append(s.errors, fmt.Errorf("expression statement required"))
			return
		}
		s.evaluators[ctx] = exprEvaluator
	} else {
		aCtx := ctx.Assignment()
		if aCtx != nil {
			assignmentEvaluator, ok := s.evaluators[aCtx]
			if !ok {
				s.errors = append(s.errors, fmt.Errorf("assignment statement required in block"))
				return
			}
			s.evaluators[ctx] = assignmentEvaluator
		} else {
			s.errors = append(s.errors, fmt.Errorf("unknown statement type in block"))
		}
	}
}

func (s *EvalTreeListener) ExitNumberConstant(ctx *parser.NumberConstantContext) {
	nStr := ""
	intCtx := ctx.INT()
	if intCtx != nil {
		nStr = intCtx.GetText()
	}
	floatCtx := ctx.FLOAT()
	if floatCtx != nil {
		nStr = floatCtx.GetText()
	}
	nStr = strings.Trim(nStr, " \t")
	if nStr == "" {
		s.errors = append(s.errors, fmt.Errorf("number constant required"))
		return
	}
	negCtx := ctx.MIN()

	n, err := strconv.ParseFloat(nStr, 64)
	if err != nil {
		s.errors = append(s.errors, err)
	}
	if negCtx != nil {
		n = -n
	}
	s.evaluators[ctx] = evaltree.NewNumberEvaluator(n)
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

func (s *EvalTreeListener) ExitIdResolution(ctx *parser.IdResolutionContext) {
	strCtx := ctx.ID()
	if strCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("identifier required"))
		return
	}
	s.evaluators[ctx] = evaltree.NewVariableEvaluator(strCtx.GetText())
}

func (s *EvalTreeListener) ExitAssignment(ctx *parser.AssignmentContext) {
	idTerm := ctx.ID()
	if idTerm == nil {
		s.errors = append(s.errors, fmt.Errorf("identifier required"))
		return
	}
	exprCtx := ctx.Expr()
	if exprCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("expression required at the right of an assignment"))
		return
	}
	expr := s.evaluators[exprCtx]
	if expr == nil {
		s.errors = append(s.errors, fmt.Errorf("expression not found at the right of an assignment"))
		return
	}
	id := idTerm.GetText()
	id = strings.Trim(id, " \t")
	if id == "" {
		s.errors = append(s.errors, fmt.Errorf("identifier should start with a letter and include letters, numbers and underscore only"))
		return
	}
	s.evaluators[ctx] = evaltree.NewAssignmentEvaluator(id, expr)
}

func (s *EvalTreeListener) ExitGrouping(ctx *parser.GroupingContext) {
	eCtx := ctx.Expr()
	if eCtx == nil {
		s.errors = append(s.errors, fmt.Errorf("grouping expression required"))
		return
	}

	exprEvaluator, ok := s.evaluators[eCtx]
	if !ok {
		ectxText := eCtx.GetText()
		s.errors = append(s.errors, fmt.Errorf("expression required for grouping, not found for [ %s ]", ectxText))
		return
	}
	s.evaluators[ctx] = exprEvaluator
}

func (s *EvalTreeListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	fnNamToken := ctx.GetFnName()
	if fnNamToken == nil {
		s.errors = append(s.errors, fmt.Errorf("function name required"))
		return
	}
	fnNameEvaluator, ok := s.evaluators[fnNamToken]
	if (!ok) || (fnNameEvaluator == nil) {
		s.errors = append(s.errors, fmt.Errorf("function name required for function"))
		return
	}

	argsExprs := ctx.AllExpr()
	argsIds := ctx.AllID()
	prms := make(map[string]evaltree.Evaluator)
	nNames := len(argsIds)
	nArgsEvaluators := len(argsExprs) - 1
	if nArgsEvaluators != nNames {
		s.errors = append(s.errors, fmt.Errorf("number of arguments does not match number of parameters"))
		return
	}
	for i := 1; i < len(argsExprs); i++ {
		argName := argsIds[i-1].GetText()
		argName = strings.Trim(argName, " \t")
		if argName == "" {
			s.errors = append(s.errors, fmt.Errorf("argument name required"))
			return
		}
		prms[argName] = s.evaluators[argsExprs[i]]
	}
	s.evaluators[ctx] = evaltree.NewFunctionCallEvaluator(fnNameEvaluator, prms)
}

func (s *EvalTreeListener) ExitMulOrDivOp(ctx *parser.MulOrDivOpContext) {
	expressions := ctx.AllExpr()
	if len(expressions) != 2 {
		s.errors = append(s.errors, fmt.Errorf("binary operation requires two expressions"))
		return
	}
	left, ok := s.evaluators[expressions[0]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("left expression required for binary operation"))
		return
	}
	right, ok := s.evaluators[expressions[1]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("right expression required for binary operation"))
		return
	}
	op := ctx.GetOp()
	if op == nil {
		s.errors = append(s.errors, fmt.Errorf("operator required for binary operation"))
		return
	}
	s.evaluators[ctx] = evaltree.NewMathExpressionEvaluator(left, right, op.GetText())
}

func (s *EvalTreeListener) ExitAddOrMinOp(ctx *parser.AddOrMinOpContext) {
	expressions := ctx.AllExpr()
	if len(expressions) != 2 {
		s.errors = append(s.errors, fmt.Errorf("binary operation requires two expressions"))
		return
	}
	left, ok := s.evaluators[expressions[0]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("left expression required for binary operation"))
		return
	}
	right, ok := s.evaluators[expressions[1]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("right expression required for binary operation"))
		return
	}
	op := ctx.GetOp()
	if op == nil {
		s.errors = append(s.errors, fmt.Errorf("operator required for binary operation"))
		return
	}
	s.evaluators[ctx] = evaltree.NewMathExpressionEvaluator(left, right, op.GetText())
}

func (s *EvalTreeListener) ExitExpOp(ctx *parser.ExpOpContext) {
	expressions := ctx.AllExpr()
	if len(expressions) != 2 {
		s.errors = append(s.errors, fmt.Errorf("binary operation requires two expressions"))
		return
	}
	left, ok := s.evaluators[expressions[0]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("left expression required for binary operation"))
		return
	}
	right, ok := s.evaluators[expressions[1]]
	if !ok {
		s.errors = append(s.errors, fmt.Errorf("right expression required for binary operation"))
		return
	}
	s.evaluators[ctx] = evaltree.NewMathExpressionEvaluator(left, right, "^")
}
