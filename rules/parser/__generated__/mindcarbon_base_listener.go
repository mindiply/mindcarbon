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

// EnterBlock is called when production block is entered.
func (s *BasemindcarbonListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasemindcarbonListener) ExitBlock(ctx *BlockContext) {}

// EnterBstat is called when production bstat is entered.
func (s *BasemindcarbonListener) EnterBstat(ctx *BstatContext) {}

// ExitBstat is called when production bstat is exited.
func (s *BasemindcarbonListener) ExitBstat(ctx *BstatContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasemindcarbonListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasemindcarbonListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIdResolution is called when production idResolution is entered.
func (s *BasemindcarbonListener) EnterIdResolution(ctx *IdResolutionContext) {}

// ExitIdResolution is called when production idResolution is exited.
func (s *BasemindcarbonListener) ExitIdResolution(ctx *IdResolutionContext) {}

// EnterMulOrDivOp is called when production mulOrDivOp is entered.
func (s *BasemindcarbonListener) EnterMulOrDivOp(ctx *MulOrDivOpContext) {}

// ExitMulOrDivOp is called when production mulOrDivOp is exited.
func (s *BasemindcarbonListener) ExitMulOrDivOp(ctx *MulOrDivOpContext) {}

// EnterStringConstant is called when production stringConstant is entered.
func (s *BasemindcarbonListener) EnterStringConstant(ctx *StringConstantContext) {}

// ExitStringConstant is called when production stringConstant is exited.
func (s *BasemindcarbonListener) ExitStringConstant(ctx *StringConstantContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasemindcarbonListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasemindcarbonListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAddOrMinOp is called when production addOrMinOp is entered.
func (s *BasemindcarbonListener) EnterAddOrMinOp(ctx *AddOrMinOpContext) {}

// ExitAddOrMinOp is called when production addOrMinOp is exited.
func (s *BasemindcarbonListener) ExitAddOrMinOp(ctx *AddOrMinOpContext) {}

// EnterExpOp is called when production expOp is entered.
func (s *BasemindcarbonListener) EnterExpOp(ctx *ExpOpContext) {}

// ExitExpOp is called when production expOp is exited.
func (s *BasemindcarbonListener) ExitExpOp(ctx *ExpOpContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BasemindcarbonListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BasemindcarbonListener) ExitGrouping(ctx *GroupingContext) {}

// EnterNumberConstant is called when production numberConstant is entered.
func (s *BasemindcarbonListener) EnterNumberConstant(ctx *NumberConstantContext) {}

// ExitNumberConstant is called when production numberConstant is exited.
func (s *BasemindcarbonListener) ExitNumberConstant(ctx *NumberConstantContext) {}
