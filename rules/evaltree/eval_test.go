package evaltree

import "testing"

func TestScalarEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{evaluator: NewNumberEvaluator(1.0), expected: NewNumberObject(1.0)},
		{evaluator: NewStringEvaluator("test"), expected: NewStringObject("test")},
		{evaluator: NewNumberEvaluator(3.14), expected: NewNumberObject(3.14)},
		{evaluator: NewStringEvaluator("hello"), expected: NewStringObject("hello")},
		{evaluator: NewFunctionDefinitionEvaluator("test",
			map[string]ParameterDefinition{}, NewNumberEvaluator((42))), expected: NULL},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(env)
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}
}

func TestMathExpressionEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(2), NewNumberEvaluator(3), "+"), expected: NewNumberObject(5)},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(5), NewNumberEvaluator(2), "-"), expected: NewNumberObject(3)},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(2), NewNumberEvaluator(3), "*"), expected: NewNumberObject(6)},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(6), NewNumberEvaluator(3), "/"), expected: NewNumberObject(2)},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(2), NewNumberEvaluator(3), "^"), expected: NewNumberObject(8)},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(5), NewNumberEvaluator(2), "%"), expected: NewError("Invalid operator '%%'")},
		{evaluator: NewMathExpressionEvaluator(NewStringEvaluator("A"), NewNumberEvaluator(2), "+"), expected: NewError("You can only perform math operations on numbers")},
		{evaluator: NewMathExpressionEvaluator(NewNumberEvaluator(8), NewStringEvaluator("wow"), "+"), expected: NewError("You can only perform math operations on numbers")},
		{evaluator: NewMathExpressionEvaluator(NewStringEvaluator("Pow"), NewStringEvaluator("wow"), "+"), expected: NewError("You can only perform math operations on numbers")},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(env)
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}
}

func TestFunctionDefinitionEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{evaluator: NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{}, NewNumberEvaluator(0)),
			expected: NewFunctionDefinitionObject("testA", map[string]ParameterDefinition{}, NewNumberEvaluator(0))},
		{evaluator: NewFunctionDefinitionEvaluator("testB", map[string]ParameterDefinition{"a": {
			Name:          "a",
			Quality:       "length",
			UnitOfMeasure: "m",
		}}, NewNumberEvaluator(0)), expected: NewFunctionDefinitionObject("testB", map[string]ParameterDefinition{"a": {
			Name:          "a",
			Quality:       "length",
			UnitOfMeasure: "m",
		}}, NewNumberEvaluator(0))},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(env)
		if result.Type() != NULLTYPE {
			t.Errorf("Expected type %s, got %s", NULLTYPE, result.Type())
		}
		def := test.evaluator.(*FunctionDefinitionEvaluator)
		v, ok := env.Get(def.name)
		if !ok {
			t.Errorf("Expected function %s to be defined", def.name)
			continue
		} else if v.Evaluate(env).Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), v.Evaluate(env).Inspect())
		}
	}
}

func TestFunctionCallEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", nil, NewNumberEvaluator(0)),
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), nil),
			}),
			expected: NewNumberObject(0),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), map[string]Evaluator{}),
			}),
			expected: NewError("Cannot resolve to a function definition"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{}, NewNumberEvaluator(0)),
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), map[string]Evaluator{
					"a": NewNumberEvaluator(1),
				}),
			}),
			expected: NewError("Argument %s not in the list of required arguments", "a"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{
					"a": NewParameterDefinition("a", "length", "m"),
				}, NewNumberEvaluator(0)),
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), map[string]Evaluator{}),
			}),
			expected: NewError("Argument %s is required", "a"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{
					"a": NewParameterDefinition("a", "length", "m"),
				}, NewBlockEvaluator([]Evaluator{
					NewVariableEvaluator("a"),
				})),
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), map[string]Evaluator{
					"a": NewNumberEvaluator(1),
				}),
			}),
			expected: NewNumberObject(1),
		},
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{
					"a": NewParameterDefinition("a", "length", "m"),
				}, NewBlockEvaluator([]Evaluator{
					NewAssignmentEvaluator("a", NewMathExpressionEvaluator(
						NewStringEvaluator("bad "),
						NewStringEvaluator("end"),
						"/",
					)),
					NewVariableEvaluator("a"),
				})),
				NewFunctionCallEvaluator(NewVariableEvaluator("testA"), map[string]Evaluator{
					"a": NewNumberEvaluator(1),
				}),
			}),
			expected: NewError("Cannot resolve expression being assigned to %s: You can only perform math operations on numbers", "a"),
		},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(NewEnvironment(env))
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}
}

func TestAssignmentEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewAssignmentEvaluator("a", NewNumberEvaluator(1)),
				NewVariableEvaluator("a"),
			}),
			expected: NewNumberObject(1),
		},
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewVariableEvaluator("a"),
			}),
			expected: NewError("variable %s not found", "a"),
		},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(NewEnvironment(env))
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}
}

func TestBlockEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewAssignmentEvaluator("a", NewNumberEvaluator(1)),
				NewVariableEvaluator("a"),
				NewMathExpressionEvaluator(
					NewVariableEvaluator("a"),
					NewStringEvaluator("bad end"),
					"+",
				),
			}),
			expected: NewError("You can only perform math operations on numbers"),
		},
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewVariableEvaluator("a"),
			}),
			expected: NewError("variable %s not found", "a"),
		},
	}
	env := NewEnvironment(nil)
	for _, test := range tests {
		result := test.evaluator.Evaluate(NewEnvironment(env))
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}
}

func TestNativeFunction_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{
			evaluator: NewFunctionCallEvaluator(NewVariableEvaluator("LookupFactorFor"), map[string]Evaluator{}),
			expected:  NewNumberObject(0),
		},
	}
	env := NewEnvironment(nil)
	env.Set("LookupFactorFor", NewNativeFunctionCallEvaluator("LookupFactorFor", LookupFactorFor))
	for _, test := range tests {
		result := test.evaluator.Evaluate(env)
		if result.Type() != test.expected.Type() {
			t.Errorf("Expected type %s, got %s", test.expected.Type(), result.Type())
		}
		if result.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), result.Inspect())
		}
	}

}
