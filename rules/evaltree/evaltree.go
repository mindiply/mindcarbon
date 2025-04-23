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
	name          string
	quality       string
	unitOfMeasure string
}

type FunctionDefinitionEvaluator struct {
	name      string
	arguments map[string]ParameterDefinition
	body      Evaluator
}

func (f *FunctionDefinitionEvaluator) Evaluate(env Environment) Object {
	env.Set(f.name, f)
	return &NullObject{}
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
	name string
}

func (v *VariableEvaluator) Evaluate(env Environment) Object {
	variable, ok := env.Get(v.name)
	if !ok {
		return NewError("Variable %s not found", v.name)
	}
	return variable.Evaluate(env)
}

type AssignmentEvaluator struct {
	name  string
	value Evaluator
}

func (a *AssignmentEvaluator) Evaluate(env Environment) Object {
	value := a.value.Evaluate(env)
	if value.Type() == ERRORTYPE {
		return value
	}
	env.Set(a.name, ScalarEvaluator{Value: value})
	return &NullObject{}
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

func NewComputationDefEvaluator(name string, arguments map[string]ParameterDefinition, body Evaluator) Evaluator {
	return &FunctionDefinitionEvaluator{name: name, arguments: arguments, body: body}
}
