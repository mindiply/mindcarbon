package parser

import (
	"github.com/mindiply/mindcarbon/rules/evaltree"
	"testing"
)

func TestNumberScalarStatements(t *testing.T) {
	input := `
	20;
	20.2;
	-20;
	-1.1
`
	values := []float64{20, 20.2, -20, -1.1}
	evaluator, err := DmlToEvalTree(input)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if evaluator == nil {
		t.Errorf("Evaluator is nil")
		return
	}
	p, ok := evaluator.(evaltree.ProgramEvaluator)
	if !ok {
		t.Errorf("Evaluator is not a ProgramEvaluator")
		return
	}
	if len(p.Statements) != 4 {
		t.Errorf("Expected 3 statements, got %d", len(p.Statements))
		return
	}
	for i, statement := range p.Statements {
		scalar, ok := statement.(evaltree.ScalarEvaluator)
		if !ok {
			t.Errorf("Expected ScalarEvaluator")
			return
		}
		v := scalar.Evaluate(evaltree.NewEnvironment(nil))
		number, ok := v.(*evaltree.NumberObject)
		if !ok {
			t.Errorf("Expected NumberObject")
			continue
		}
		if number.Value() != values[i] {
			t.Errorf("Expected %f, got %f", values[i], number.Value())
		}
	}
}
