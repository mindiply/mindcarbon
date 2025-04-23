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

	// Visit a parse tree produced by mindcarbonParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#division.
	VisitDivision(ctx *DivisionContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#exponentiation.
	VisitExponentiation(ctx *ExponentiationContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#intConstant.
	VisitIntConstant(ctx *IntConstantContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#stringConstant.
	VisitStringConstant(ctx *StringConstantContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#floatConstant.
	VisitFloatConstant(ctx *FloatConstantContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#subtraction.
	VisitSubtraction(ctx *SubtractionContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#grouping.
	VisitGrouping(ctx *GroupingContext) interface{}

	// Visit a parse tree produced by mindcarbonParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}
}
