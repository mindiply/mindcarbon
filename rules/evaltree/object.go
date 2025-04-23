package evaltree

import "fmt"

const (
	NULLTYPE   = "null"
	NUMBERTYPE = "number"
	STRINGTYPE = "string"
	ERRORTYPE  = "error"
)

type ObjectType string

type Object interface {
	Type() ObjectType

	// Value returns the Value encapsulated by the object.
	Value() interface{}

	Inspect() string
}

type NumberObject struct {
	FValue float64
}

func (n NumberObject) Type() ObjectType {
	return NUMBERTYPE
}

func (n NumberObject) Value() interface{} {
	return n.FValue
}

func (n NumberObject) Inspect() string {
	return fmt.Sprintf("Number {%f}", n.FValue)
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
	return nil
}

func (e ErrorObject) Inspect() string {
	return fmt.Sprintf("Error: {%s}", e.Message)
}

func NewError(format string, a ...interface{}) *ErrorObject {
	return &ErrorObject{Message: fmt.Sprintf(format, a...)}
}
