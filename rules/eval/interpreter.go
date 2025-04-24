package eval

import (
	"fmt"
	"github.com/mindiply/mindcarbon/rules/evaltree"
	"github.com/mindiply/mindcarbon/rules/parser"
)

type Interpreter interface {
	EvaluateDML(dml string) (evaltree.Object, error)
	RootEnv() evaltree.Environment
}

// Interpreter is the main struct for the interpreter
type interpreter struct {
	// RootEnv is the root environment for the interpreter
	rootEnv evaltree.Environment
}

// EvaluateDML evaluates the DML string and returns the result
func (i *interpreter) EvaluateDML(dml string) (evaltree.Object, error) {
	program, err := parser.DmlToEvalTree(dml)
	if err != nil {
		return nil, err
	}
	if program == nil {
		return nil, fmt.Errorf("program is nil")
	}
	result := program.Evaluate(i.rootEnv)
	if result == nil {
		return evaltree.NULL, fmt.Errorf("unexpected error: result is nil")
	} else if result.Type() == evaltree.ERRORTYPE {
		return result, fmt.Errorf("execution error: %s", result.Value())
	}
	return result, nil
}

// RootEnv returns the root environment
func (i *interpreter) RootEnv() evaltree.Environment {
	return i.rootEnv
}

// NewInterpreter creates a new interpreter
func NewInterpreter() Interpreter {
	rootEnv := evaltree.NewEnvironment(nil)
	// Add built-in functions to the root environment
	rootEnv.Set("LookupFactorFor", evaltree.NewNativeFunctionCallEvaluator("LookupFactorFor", evaltree.LookupFactorFor))
	return &interpreter{
		rootEnv: rootEnv,
	}
}
