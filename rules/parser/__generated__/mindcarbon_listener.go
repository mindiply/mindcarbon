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

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterDivision is called when entering the division production.
	EnterDivision(c *DivisionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterExponentiation is called when entering the exponentiation production.
	EnterExponentiation(c *ExponentiationContext)

	// EnterIntConstant is called when entering the intConstant production.
	EnterIntConstant(c *IntConstantContext)

	// EnterStringConstant is called when entering the stringConstant production.
	EnterStringConstant(c *StringConstantContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFloatConstant is called when entering the floatConstant production.
	EnterFloatConstant(c *FloatConstantContext)

	// EnterSubtraction is called when entering the subtraction production.
	EnterSubtraction(c *SubtractionContext)

	// EnterMultiplication is called when entering the multiplication production.
	EnterMultiplication(c *MultiplicationContext)

	// EnterGrouping is called when entering the grouping production.
	EnterGrouping(c *GroupingContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitComputationDef is called when exiting the computationDef production.
	ExitComputationDef(c *ComputationDefContext)

	// ExitParamDef is called when exiting the paramDef production.
	ExitParamDef(c *ParamDefContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitDivision is called when exiting the division production.
	ExitDivision(c *DivisionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitExponentiation is called when exiting the exponentiation production.
	ExitExponentiation(c *ExponentiationContext)

	// ExitIntConstant is called when exiting the intConstant production.
	ExitIntConstant(c *IntConstantContext)

	// ExitStringConstant is called when exiting the stringConstant production.
	ExitStringConstant(c *StringConstantContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFloatConstant is called when exiting the floatConstant production.
	ExitFloatConstant(c *FloatConstantContext)

	// ExitSubtraction is called when exiting the subtraction production.
	ExitSubtraction(c *SubtractionContext)

	// ExitMultiplication is called when exiting the multiplication production.
	ExitMultiplication(c *MultiplicationContext)

	// ExitGrouping is called when exiting the grouping production.
	ExitGrouping(c *GroupingContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)
}
