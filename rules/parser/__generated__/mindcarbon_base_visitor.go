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

func (v *BasemindcarbonVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitBstat(ctx *BstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitIdResolution(ctx *IdResolutionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitMulOrDivOp(ctx *MulOrDivOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitStringConstant(ctx *StringConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitAddOrMinOp(ctx *AddOrMinOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitExpOp(ctx *ExpOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitGrouping(ctx *GroupingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindcarbonVisitor) VisitNumberConstant(ctx *NumberConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
