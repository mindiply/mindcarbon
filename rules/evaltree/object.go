package evaltree

import (
	"fmt"
)

const (
	NULLTYPE       = "null"
	NUMBERTYPE     = "number"
	STRINGTYPE     = "string"
	ERRORTYPE      = "error"
	FN_DEF_TYPE    = "functionDefinition"
	NATIVE_FN_TYPE = "nativeFunction"
)

type ObjectType string

type Object interface {
	Type() ObjectType

	// Value returns the Value encapsulated by the object.
	Value() interface{}

	Inspect() string
}

type NumberObject struct {
	value float64
}

func (n NumberObject) Type() ObjectType {
	return NUMBERTYPE
}

func (n NumberObject) Value() interface{} {
	return n.value
}

func (n NumberObject) Inspect() string {
	return fmt.Sprintf("Number {%f}", n.value)
}

func NewNumberObject(value float64) *NumberObject {
	return &NumberObject{value: value}
}

type StringObject struct {
	value string
}

func (s StringObject) Type() ObjectType {
	return STRINGTYPE
}

func (s StringObject) Value() interface{} {
	return s.value
}

func (s StringObject) Inspect() string {
	return fmt.Sprintf("String {%s}", s.value)
}

func NewStringObject(value string) *StringObject {
	return &StringObject{value: value}
}

type NullObject struct {
}

var NULL = NullObject{}

func (n NullObject) Type() ObjectType {
	return NULLTYPE
}

func (n NullObject) Value() interface{} {
	return NULL
}

func (n NullObject) Inspect() string {
	return "null"
}

func IsNull(obj Object) bool {
	return obj == NULL
}

type ErrorObject struct {
	Message string
}

func (e ErrorObject) Type() ObjectType {
	return ERRORTYPE
}

func (e ErrorObject) Value() interface{} {
	return e.Message
}

func (e ErrorObject) Inspect() string {
	return fmt.Sprintf("Error: {%s}", e.Message)
}

func NewError(format string, a ...interface{}) *ErrorObject {
	return &ErrorObject{Message: fmt.Sprintf(format, a...)}
}

type FunctionDefinitionObject struct {
	Name       string
	Parameters map[string]ParameterDefinition
	Body       Evaluator
}

func (f FunctionDefinitionObject) Type() ObjectType {
	return FN_DEF_TYPE
}

func (f FunctionDefinitionObject) Value() interface{} {
	return f
}

func (f FunctionDefinitionObject) Inspect() string {
	return fmt.Sprintf("FunctionDefinition {%s}", f.Name)
}

func NewFunctionDefinitionObject(name string, parameters map[string]ParameterDefinition, body Evaluator) *FunctionDefinitionObject {
	return &FunctionDefinitionObject{Name: name, Parameters: parameters, Body: body}
}

type NativeFunctionDefinitionObject struct {
	Name string
	Fn   NativeFunction
}

func (n NativeFunctionDefinitionObject) Type() ObjectType {
	return NATIVE_FN_TYPE
}

func (n NativeFunctionDefinitionObject) Value() interface{} {
	return n
}

func (n NativeFunctionDefinitionObject) Inspect() string {
	return fmt.Sprintf("NativeFunctionDefinition {%s}", n.Name)
}

func NewNativeFunctionDefinitionObject(name string, fn NativeFunction) *NativeFunctionDefinitionObject {
	return &NativeFunctionDefinitionObject{Name: name, Fn: fn}
}
