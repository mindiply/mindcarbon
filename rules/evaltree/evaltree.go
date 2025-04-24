package evaltree

import (
	"fmt"
	"math"
)

type Evaluator interface {
	// Evaluate evaluates the expression and returns the result.
	Evaluate(env Environment) Object
}

type ScalarEvaluator struct {
	Value Object
}

func (s ScalarEvaluator) Evaluate(env Environment) Object {
	return s.Value
}

type ProgramEvaluator struct {
	Statements []Evaluator
}

func (p ProgramEvaluator) Evaluate(env Environment) Object {
	var v Object
	for _, statement := range p.Statements {
		v = statement.Evaluate(env)
		if v.Type() == ERRORTYPE {
			return v
		}
	}
	return v
}

type MathExpressionEvaluator struct {
	left     Evaluator
	right    Evaluator
	operator string
}

const TRUE = "true"
const FALSE = "false"

func (m *MathExpressionEvaluator) Evaluate(enf Environment) Object {
	leftValue := m.left.Evaluate(enf)
	rightValue := m.right.Evaluate(enf)
	left, lok := leftValue.(NumberObject)
	if !lok {
		return NewError(fmt.Sprintf("Cannot evaluate left expression"))
	}
	right, rok := rightValue.(NumberObject)
	if !rok {
		return NewError(fmt.Sprintf("Cannot evaluate right expression"))
	}

	switch m.operator {
	case "+":
		return &NumberObject{FValue: left.FValue + right.FValue}
	case "-":
		return &NumberObject{FValue: left.FValue - right.FValue}
	case "*":
		return &NumberObject{FValue: left.FValue / right.FValue}
	case "/":
		return &NumberObject{FValue: left.FValue / right.FValue}
	case "^":
		return &NumberObject{FValue: math.Pow(left.FValue, right.FValue)}
	default:
		return NewError(fmt.Sprintf("Unknown operator '%s'", m.operator))
	}
}

type ParameterDefinition struct {
	Name          string
	Quality       string
	UnitOfMeasure string
}

type FunctionDefinitionEvaluator struct {
	name      string
	arguments map[string]ParameterDefinition
	body      Evaluator
}

func (f *FunctionDefinitionEvaluator) Evaluate(env Environment) Object {
	env.Set(f.name, f)
	return NULL
}

type BlockEvaluator struct {
	statements []Evaluator
}

func (b *BlockEvaluator) Evaluate(enf Environment) Object {
	var result Object
	for _, statement := range b.statements {
		result = statement.Evaluate(enf)
		if result.Type() == ERRORTYPE {
			return result
		}
	}
	return result

}

type FunctionCallEvaluator struct {
	name      string
	arguments map[string]Evaluator
}

func (f *FunctionCallEvaluator) Evaluate(env Environment) Object {
	fnDef, ok := env.Get(f.name)
	if !ok {
		return NewError("Function %s not found", f.name)
	}
	functionDef, ok := fnDef.(*FunctionDefinitionEvaluator)
	if !ok {
		return NewError("Identifier %s is not a function", f.name)
	}
	fEnv := NewEnvironment(env)
	for name, arg := range f.arguments {
		if _, ok := functionDef.arguments[name]; !ok {
			return NewError("Argument %s not in the list of required arguments", name)
		}
		value := arg.Evaluate(env)
		var valEvaluator Evaluator
		valEvaluator = ScalarEvaluator{Value: value}
		fEnv.Set(name, valEvaluator)
	}

	for name, _ := range functionDef.arguments {
		if _, ok := fEnv.Get(name); !ok {
			return NewError("Argument %s is required", name)
		}
		// ToDo: check if the argument is of the correct type
	}
	return functionDef.body.Evaluate(fEnv)
}

type ParenthesizedEvaluator struct {
	innerExpression Evaluator
}

func (p *ParenthesizedEvaluator) Evaluate(env Environment) Object {
	return p.innerExpression.Evaluate(env)
}

type VariableEvaluator struct {
	Name string
}

func (v VariableEvaluator) Evaluate(env Environment) Object {
	variable, ok := env.Get(v.Name)
	if !ok {
		return NewError("Variable %s not found", v.Name)
	}
	return variable.Evaluate(env)
}

type AssignmentEvaluator struct {
	name          string
	bodyEvaluator Evaluator
}

func (a *AssignmentEvaluator) Evaluate(env Environment) Object {
	value := a.bodyEvaluator.Evaluate(env)
	if value.Type() == ERRORTYPE {
		return NewError("Cannot resolve expression being assigned to %s: %s", a.name, value.Value())
	}
	env.Set(a.name, a.bodyEvaluator)
	return NULL
}

func NewVariableEvaluator(name string) Evaluator {
	return &VariableEvaluator{Name: name}
}

func NewProgramEvaluator(statements []Evaluator) Evaluator {
	return ProgramEvaluator{Statements: statements}
}

func NewNumberEvaluator(value float64) Evaluator {
	return ScalarEvaluator{Value: &NumberObject{FValue: value}}
}

func NewStringEvaluator(value string) Evaluator {
	return ScalarEvaluator{Value: &StringObject{value: value}}
}

func NewParameterDefinition(name, quality, unitOfMeasure string) ParameterDefinition {
	return ParameterDefinition{Name: name, Quality: quality, UnitOfMeasure: unitOfMeasure}
}

func NewComputationDefEvaluator(name string, arguments map[string]ParameterDefinition, body Evaluator) Evaluator {
	return &FunctionDefinitionEvaluator{name: name, arguments: arguments, body: body}
}

func NewMathExpressionEvaluator(left, right Evaluator, operator string) Evaluator {
	return &MathExpressionEvaluator{left: left, right: right, operator: operator}
}

func NewFunctionCallEvaluator(name string, arguments map[string]Evaluator) Evaluator {
	if arguments == nil {
		arguments = make(map[string]Evaluator)
	}
	return &FunctionCallEvaluator{name: name, arguments: arguments}
}

func NewAssignmentEvaluator(name string, bodyEvaluator Evaluator) Evaluator {
	return &AssignmentEvaluator{name: name, bodyEvaluator: bodyEvaluator}
}

func NewBlockEvaluator(statements []Evaluator) Evaluator {
	return &BlockEvaluator{statements: statements}
}
