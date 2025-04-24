// Code generated from mindcarbon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindcarbon

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by mindcarbonParser.
type mindcarbonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mindcarbonParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#computationDef.
	VisitComputationDef(ctx *ComputationDefContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#paramDef.
	VisitParamDef(ctx *ParamDefContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#bstat.
	VisitBstat(ctx *BstatContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#idResolution.
	VisitIdResolution(ctx *IdResolutionContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#mulOrDivOp.
	VisitMulOrDivOp(ctx *MulOrDivOpContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#stringConstant.
	VisitStringConstant(ctx *StringConstantContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#addOrMinOp.
	VisitAddOrMinOp(ctx *AddOrMinOpContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#expOp.
	VisitExpOp(ctx *ExpOpContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#grouping.
	VisitGrouping(ctx *GroupingContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#numberConstant.
	VisitNumberConstant(ctx *NumberConstantContext) interface{}
}
