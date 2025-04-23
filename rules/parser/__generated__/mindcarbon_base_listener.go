// Code generated from mindcarbon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindcarbon

import "github.com/antlr4-go/antlr/v4"

// BasemindcarbonListener is a complete listener for a parse tree produced by mindcarbonParser.
type BasemindcarbonListener struct{}

var _ mindcarbonListener = &BasemindcarbonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemindcarbonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemindcarbonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemindcarbonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemindcarbonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasemindcarbonListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasemindcarbonListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasemindcarbonListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasemindcarbonListener) ExitStatement(ctx *StatementContext) {}

// EnterComputationDef is called when production computationDef is entered.
func (s *BasemindcarbonListener) EnterComputationDef(ctx *ComputationDefContext) {}

// ExitComputationDef is called when production computationDef is exited.
func (s *BasemindcarbonListener) ExitComputationDef(ctx *ComputationDefContext) {}

// EnterParamDef is called when production paramDef is entered.
func (s *BasemindcarbonListener) EnterParamDef(ctx *ParamDefContext) {}

// ExitParamDef is called when production paramDef is exited.
func (s *BasemindcarbonListener) ExitParamDef(ctx *ParamDefContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasemindcarbonListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasemindcarbonListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterDivision is called when production division is entered.
func (s *BasemindcarbonListener) EnterDivision(ctx *DivisionContext) {}

// ExitDivision is called when production division is exited.
func (s *BasemindcarbonListener) ExitDivision(ctx *DivisionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasemindcarbonListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasemindcarbonListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterExponentiation is called when production exponentiation is entered.
func (s *BasemindcarbonListener) EnterExponentiation(ctx *ExponentiationContext) {}

// ExitExponentiation is called when production exponentiation is exited.
func (s *BasemindcarbonListener) ExitExponentiation(ctx *ExponentiationContext) {}

// EnterIntConstant is called when production intConstant is entered.
func (s *BasemindcarbonListener) EnterIntConstant(ctx *IntConstantContext) {}

// ExitIntConstant is called when production intConstant is exited.
func (s *BasemindcarbonListener) ExitIntConstant(ctx *IntConstantContext) {}

// EnterStringConstant is called when production stringConstant is entered.
func (s *BasemindcarbonListener) EnterStringConstant(ctx *StringConstantContext) {}

// ExitStringConstant is called when production stringConstant is exited.
func (s *BasemindcarbonListener) ExitStringConstant(ctx *StringConstantContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasemindcarbonListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasemindcarbonListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFloatConstant is called when production floatConstant is entered.
func (s *BasemindcarbonListener) EnterFloatConstant(ctx *FloatConstantContext) {}

// ExitFloatConstant is called when production floatConstant is exited.
func (s *BasemindcarbonListener) ExitFloatConstant(ctx *FloatConstantContext) {}

// EnterSubtraction is called when production subtraction is entered.
func (s *BasemindcarbonListener) EnterSubtraction(ctx *SubtractionContext) {}

// ExitSubtraction is called when production subtraction is exited.
func (s *BasemindcarbonListener) ExitSubtraction(ctx *SubtractionContext) {}

// EnterMultiplication is called when production multiplication is entered.
func (s *BasemindcarbonListener) EnterMultiplication(ctx *MultiplicationContext) {}

// ExitMultiplication is called when production multiplication is exited.
func (s *BasemindcarbonListener) ExitMultiplication(ctx *MultiplicationContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BasemindcarbonListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BasemindcarbonListener) ExitGrouping(ctx *GroupingContext) {}

// EnterAddition is called when production addition is entered.
func (s *BasemindcarbonListener) EnterAddition(ctx *AdditionContext) {}

// ExitAddition is called when production addition is exited.
func (s *BasemindcarbonListener) ExitAddition(ctx *AdditionContext) {}
