// Code generated from mindcarbon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindcarbon

import "github.com/antlr4-go/antlr/v4"

type BasemindcarbonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemindcarbonVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitComputationDef(ctx *ComputationDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitParamDef(ctx *ParamDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitDivision(ctx *DivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitExponentiation(ctx *ExponentiationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitIntConstant(ctx *IntConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitStringConstant(ctx *StringConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitFloatConstant(ctx *FloatConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitSubtraction(ctx *SubtractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitGrouping(ctx *GroupingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitAddition(ctx *AdditionContext) interface{} {
	return v.VisitChildren(ctx)
}
