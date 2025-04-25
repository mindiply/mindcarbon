package evaltree

import "testing"

func TestScalarEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		evaluator Evaluator
		expected  Object
	}{
		{evaluator: NewLiteralEvaluator(1.0), expected: NewNumberObject(1.0)},
		{evaluator: NewLiteralEvaluator("test"), expected: NewStringObject("test")},
		{evaluator: NewLiteralEvaluator(3.14), expected: NewNumberObject(3.14)},
		{evaluator: NewLiteralEvaluator("hello"), expected: NewStringObject("hello")},
		{evaluator: NewFunctionDefinitionEvaluator("test",
			map[string]ParameterDefinition{}, NewLiteralEvaluator((42))), expected: NULL},
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
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(2), NewLiteralEvaluator(3), "+"), expected: NewNumberObject(5)},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(5), NewLiteralEvaluator(2), "-"), expected: NewNumberObject(3)},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(2), NewLiteralEvaluator(3), "*"), expected: NewNumberObject(6)},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(6), NewLiteralEvaluator(3), "/"), expected: NewNumberObject(2)},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(2), NewLiteralEvaluator(3), "^"), expected: NewNumberObject(8)},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(5), NewLiteralEvaluator(2), "%"), expected: NewError("Invalid operator '%%'")},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator("A"), NewLiteralEvaluator(2), "+"), expected: NewError("You can only perform math operations on numbers")},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator(8), NewLiteralEvaluator("wow"), "+"), expected: NewError("You can only perform math operations on numbers")},
		{evaluator: NewMathExpressionEvaluator(NewLiteralEvaluator("Pow"), NewLiteralEvaluator("wow"), "+"), expected: NewError("You can only perform math operations on numbers")},
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
		{evaluator: NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{}, NewLiteralEvaluator(0)),
			expected: NewFunctionDefinitionObject("testA", map[string]ParameterDefinition{}, NewLiteralEvaluator(0))},
		{evaluator: NewFunctionDefinitionEvaluator("testB", map[string]ParameterDefinition{"a": {
			Name:          "a",
			Quality:       "length",
			UnitOfMeasure: "m",
		}}, NewLiteralEvaluator(0)), expected: NewFunctionDefinitionObject("testB", map[string]ParameterDefinition{"a": {
			Name:          "a",
			Quality:       "length",
			UnitOfMeasure: "m",
		}}, NewLiteralEvaluator(0))},
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
		} else if v.Inspect() != test.expected.Inspect() {
			t.Errorf("Expected %s, got %s", test.expected.Inspect(), v.Inspect())
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
				NewFunctionDefinitionEvaluator("testA", nil, NewLiteralEvaluator(0)),
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), nil),
			}),
			expected: NewNumberObject(0),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), map[string]Evaluator{}),
			}),
			expected: NewError("Cannot resolve to a function definition"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{}, NewLiteralEvaluator(0)),
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), map[string]Evaluator{
					"a": NewLiteralEvaluator(1),
				}),
			}),
			expected: NewError("Argument %s not in the list of required arguments", "a"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{
					"a": NewParameterDefinition("a", "length", "m"),
				}, NewLiteralEvaluator(0)),
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), map[string]Evaluator{}),
			}),
			expected: NewError("Argument %s is required", "a"),
		},
		{
			evaluator: NewProgramEvaluator([]Evaluator{
				NewFunctionDefinitionEvaluator("testA", map[string]ParameterDefinition{
					"a": NewParameterDefinition("a", "length", "m"),
				}, NewBlockEvaluator([]Evaluator{
					NewIdentifierEvaluator("a"),
				})),
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), map[string]Evaluator{
					"a": NewLiteralEvaluator(1),
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
						NewLiteralEvaluator("bad "),
						NewLiteralEvaluator("end"),
						"/",
					)),
					NewIdentifierEvaluator("a"),
				})),
				NewFunctionCallEvaluator(NewIdentifierEvaluator("testA"), map[string]Evaluator{
					"a": NewLiteralEvaluator(1),
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
				NewAssignmentEvaluator("a", NewLiteralEvaluator(1)),
				NewIdentifierEvaluator("a"),
			}),
			expected: NewNumberObject(1),
		},
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewIdentifierEvaluator("a"),
			}),
			expected: NewError("identifier %s not found", "a"),
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
				NewAssignmentEvaluator("a", NewLiteralEvaluator(1)),
				NewIdentifierEvaluator("a"),
				NewMathExpressionEvaluator(
					NewIdentifierEvaluator("a"),
					NewLiteralEvaluator("bad end"),
					"+",
				),
			}),
			expected: NewError("You can only perform math operations on numbers"),
		},
		{
			evaluator: NewBlockEvaluator([]Evaluator{
				NewIdentifierEvaluator("a"),
			}),
			expected: NewError("identifier %s not found", "a"),
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
			evaluator: NewFunctionCallEvaluator(NewIdentifierEvaluator("lookupFactorFor"), map[string]Evaluator{}),
			expected:  NewNumberObject(0),
		},
	}
	env := NewEnvironment(nil)
	for name, fn := range GetNativeFunctions() {
		env.Set(name, NewNativeFunctionDefinitionObject(name, fn))
	}

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
