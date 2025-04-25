package evaltree

import (
	"math"
)

type Evaluator interface {
	// Evaluate evaluates the expression and returns the result.
	Evaluate(env Environment) Object
}
type ProgramEvaluator struct {
	Statements []Evaluator
}

func (p ProgramEvaluator) Evaluate(env Environment) Object {
	var v Object = NULL
	for _, statement := range p.Statements {
		v = statement.Evaluate(env)
		if v.Type() == ERRORTYPE {
			return v
		}
	}
	return v
}

type ScalarEvaluator struct {
	value Object
}

func (s ScalarEvaluator) Evaluate(env Environment) Object {
	if s.value == nil {
		return NewError("ScalarEvaluator is nil")
	}
	return s.value
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
	left, lok := leftValue.Value().(float64)
	if !lok {
		return NewError("You can only perform math operations on numbers")
	}
	right, rok := rightValue.Value().(float64)
	if !rok {
		return NewError("You can only perform math operations on numbers")
	}

	switch m.operator {
	case "+":
		return NewNumberObject(left + right)
	case "-":
		return NewNumberObject(left - right)
	case "*":
		return NewNumberObject(left * right)
	case "/":
		return NewNumberObject(left / right)
	case "^":
		return NewNumberObject(math.Pow(left, right))
	default:
		return NewError("Invalid operator '%s'", m.operator)
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
	env.Set(f.name, NewFunctionDefinitionObject(f.name, f.arguments, f.body))
	return NULL
}

type BlockEvaluator struct {
	statements []Evaluator
}

func (b *BlockEvaluator) Evaluate(enf Environment) Object {
	var result Object = NULL
	for _, statement := range b.statements {
		result = statement.Evaluate(enf)
		if result.Type() == ERRORTYPE {
			return result
		}
	}
	return result

}

type FunctionCallEvaluator struct {
	nameEvaluator Evaluator
	arguments     map[string]Evaluator
}

func (f *FunctionCallEvaluator) Evaluate(env Environment) Object {
	fnDef := f.nameEvaluator.Evaluate(env)
	if fnDef.Type() != FN_DEF_TYPE && fnDef.Type() != NATIVE_FN_TYPE {
		return NewError("Cannot resolve to a function definition")
	}
	functionDef, ok := fnDef.(*FunctionDefinitionObject)
	if ok {
		fEnv := NewEnvironment(env)
		for name, arg := range f.arguments {
			if _, ok := functionDef.Parameters[name]; !ok {
				return NewError("Argument %s not in the list of required arguments", name)
			}
			fEnv.Set(name, arg.Evaluate(env))
		}

		for name, _ := range functionDef.Parameters {
			if _, ok := fEnv.Get(name); !ok {
				return NewError("Argument %s is required", name)
			}
			// ToDo: check if the argument is of the correct type
		}
		return functionDef.Body.Evaluate(fEnv)
	}
	nativeDef, ok := fnDef.(*NativeFunctionDefinitionObject)
	if !ok {
		return NewError("Cannot resolve to a valid function definition")
	}
	resolvedArguments := make(map[string]Object)
	for name, arg := range f.arguments {
		resolvedArguments[name] = arg.Evaluate(env)
	}
	return nativeDef.Fn(resolvedArguments)
}

type IdentifierEvaluator struct {
	Name string
}

func (v IdentifierEvaluator) Evaluate(env Environment) Object {
	value, ok := env.Get(v.Name)
	if !ok {
		return NewError("identifier %s not found", v.Name)
	}
	return value
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
	env.Set(a.name, value)
	return NULL
}

func NewIdentifierEvaluator(name string) Evaluator {
	return &IdentifierEvaluator{Name: name}
}

func NewProgramEvaluator(statements []Evaluator) Evaluator {
	return ProgramEvaluator{Statements: statements}
}

func NewParameterDefinition(name, quality, unitOfMeasure string) ParameterDefinition {
	return ParameterDefinition{Name: name, Quality: quality, UnitOfMeasure: unitOfMeasure}
}

func NewFunctionDefinitionEvaluator(name string, arguments map[string]ParameterDefinition, body Evaluator) Evaluator {
	return &FunctionDefinitionEvaluator{name: name, arguments: arguments, body: body}
}

func NewMathExpressionEvaluator(left, right Evaluator, operator string) Evaluator {
	return &MathExpressionEvaluator{left: left, right: right, operator: operator}
}

func NewFunctionCallEvaluator(nameEvaluator Evaluator, arguments map[string]Evaluator) Evaluator {
	if arguments == nil {
		arguments = make(map[string]Evaluator)
	}
	return &FunctionCallEvaluator{nameEvaluator: nameEvaluator, arguments: arguments}
}

func NewAssignmentEvaluator(name string, bodyEvaluator Evaluator) Evaluator {
	return &AssignmentEvaluator{name: name, bodyEvaluator: bodyEvaluator}
}

func NewBlockEvaluator(statements []Evaluator) Evaluator {
	return &BlockEvaluator{statements: statements}
}

type NativeFunctionCallEvaluator struct {
	nativeFnObject *NativeFunctionDefinitionObject
}

func (n *NativeFunctionCallEvaluator) Evaluate(env Environment) Object {
	return n.nativeFnObject
}

func NewNativeFunctionCallEvaluator(name string, nativeFunction NativeFunction) Evaluator {
	return &NativeFunctionCallEvaluator{nativeFnObject: NewNativeFunctionDefinitionObject(
		name, nativeFunction)}
}

func NewLiteralEvaluator(value interface{}) Evaluator {
	if value == nil {
		return NewError("Cannot evaluate nil literal")
	}
	switch v := value.(type) {
	case string:
		return NewStringObject(v)
	case int:
		return NewNumberObject(float64(v))
	case int8:
		return NewNumberObject(float64(v))
	case int16:
		return NewNumberObject(float64(v))
	case int32:
		return NewNumberObject(float64(v))
	case int64:
		return NewNumberObject(float64(v))
	case uint:
		return NewNumberObject(float64(v))
	case uint8:
		return NewNumberObject(float64(v))
	case uint16:
		return NewNumberObject(float64(v))
	case uint32:
		return NewNumberObject(float64(v))
	case uint64:
		return NewNumberObject(float64(v))
	case float32:
		return NewNumberObject(float64(v))
	case float64:
		return NewNumberObject(v)
	default:
		return NewError("Unsupported literal type")
	}
}
