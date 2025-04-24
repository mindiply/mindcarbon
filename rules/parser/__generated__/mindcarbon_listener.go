// Code generated from mindcarbon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindcarbon

import "github.com/antlr4-go/antlr/v4"

// mindcarbonListener is a complete listener for a parse tree produced by mindcarbonParser.
type mindcarbonListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterComputationDef is called when entering the computationDef production.
	EnterComputationDef(c *ComputationDefContext)

	// EnterParamDef is called when entering the paramDef production.
	EnterParamDef(c *ParamDefContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBstat is called when entering the bstat production.
	EnterBstat(c *BstatContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterIdResolution is called when entering the idResolution production.
	EnterIdResolution(c *IdResolutionContext)

	// EnterMulOrDivOp is called when entering the mulOrDivOp production.
	EnterMulOrDivOp(c *MulOrDivOpContext)

	// EnterStringConstant is called when entering the stringConstant production.
	EnterStringConstant(c *StringConstantContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAddOrMinOp is called when entering the addOrMinOp production.
	EnterAddOrMinOp(c *AddOrMinOpContext)

	// EnterExpOp is called when entering the expOp production.
	EnterExpOp(c *ExpOpContext)

	// EnterGrouping is called when entering the grouping production.
	EnterGrouping(c *GroupingContext)

	// EnterNumberConstant is called when entering the numberConstant production.
	EnterNumberConstant(c *NumberConstantContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitComputationDef is called when exiting the computationDef production.
	ExitComputationDef(c *ComputationDefContext)

	// ExitParamDef is called when exiting the paramDef production.
	ExitParamDef(c *ParamDefContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBstat is called when exiting the bstat production.
	ExitBstat(c *BstatContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitIdResolution is called when exiting the idResolution production.
	ExitIdResolution(c *IdResolutionContext)

	// ExitMulOrDivOp is called when exiting the mulOrDivOp production.
	ExitMulOrDivOp(c *MulOrDivOpContext)

	// ExitStringConstant is called when exiting the stringConstant production.
	ExitStringConstant(c *StringConstantContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAddOrMinOp is called when exiting the addOrMinOp production.
	ExitAddOrMinOp(c *AddOrMinOpContext)

	// ExitExpOp is called when exiting the expOp production.
	ExitExpOp(c *ExpOpContext)

	// ExitGrouping is called when exiting the grouping production.
	ExitGrouping(c *GroupingContext)

	// ExitNumberConstant is called when exiting the numberConstant production.
	ExitNumberConstant(c *NumberConstantContext)
}
