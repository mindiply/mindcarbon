package parser

import (
	"fmt"
	"github.com/mindiply/mindcarbon/rules/evaltree"
	"reflect"
	"strings"
	"testing"
)

func TestNumberScalarStatements(t *testing.T) {
	input := `
	20;
	20.2;
	-20;
	-1.1;
	0;
`
	values := []float64{20, 20.2, -20, -1.1, 0}
	p, err := parseProgram(input)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != len(values) {
		t.Errorf("Expected %d statements, got %d", len(values), len(p.Statements))
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

func TestStringScalarStatements(t *testing.T) {
	input := `
	"hello";
	"world";
	"hello world";
	"What
is
this?";
	""
`
	values := []string{"hello", "world", "hello world", "What\nis\nthis?", ""}
	p, err := parseProgram(input)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != 5 {
		t.Errorf("Expected 5 statements, got %d", len(p.Statements))
		return
	}
	for i, statement := range p.Statements {
		scalar, ok := statement.(evaltree.ScalarEvaluator)
		if !ok {
			t.Errorf("Expected ScalarEvaluator")
			return
		}
		v := scalar.Evaluate(evaltree.NewEnvironment(nil))
		str, ok := v.(*evaltree.StringObject)
		if !ok {
			t.Errorf("Expected StringObject")
			continue
		}
		if str.Value() != values[i] {
			t.Errorf("Expected %s, got %s", values[i], str.Value())
		}
	}
}

func TestIdResolutionStatements(t *testing.T) {
	input := `
	hello;
	i;
	j_23o;
`
	varNames := []string{"hello", "i", "j_23o"}
	p, err := parseProgram(input)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != 3 {
		t.Errorf("Expected 3 statements, got %d", len(p.Statements))
		return
	}
	for i, statement := range p.Statements {
		assign, ok := statement.(*evaltree.IdentifierEvaluator)
		if !ok {
			t.Errorf("Expected IdentifierEvaluator")
			continue
		}
		if assign.Name != varNames[i] {
			t.Errorf("Expected %s, got %s", varNames[i], assign.Name)
			continue
		}
	}
}

func parseProgram(dml string) (*evaltree.ProgramEvaluator, error) {
	evaluator, err := DmlToEvalTree(dml)
	if err != nil {
		return nil, err
	}
	if evaluator == nil {
		return nil, fmt.Errorf("evaluator is nil")
	}
	p, ok := evaluator.(evaltree.ProgramEvaluator)
	if !ok {
		return nil, fmt.Errorf("evaluator is not a ProgramEvaluator")
	}
	return &p, nil
}

func TestGroupingExpressionStatements(t *testing.T) {
	input := `
	((2.2));
	(2.2);
	(2);
	("a");
	((("hello")));
`
	exprConstants := []interface{}{evaltree.NewLiteralEvaluator(2.2),
		evaltree.NewLiteralEvaluator(2.2),
		evaltree.NewLiteralEvaluator(2),
		evaltree.NewLiteralEvaluator("a"),
		evaltree.NewLiteralEvaluator("hello")}
	p, err := parseProgram(input)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != 5 {
		t.Errorf("Expected 5 statements, got %d", len(p.Statements))
		return
	}
	for i, statement := range p.Statements {
		if !reflect.DeepEqual(statement, exprConstants[i]) {
			t.Errorf("Expected %v, got %v", exprConstants[i], statement)
			continue
		}
	}
}

func TestMathExpressionStatements(t *testing.T) {
	operations := []string{"^", "+", "-", "*", "/"}
	values := [][]float64{
		{1, 0},
		{2.2, 4},
		{2, 1.5},
		{0, 13000},
		{250, 8},
	}
	for _, op := range operations {
		input := ""
		var exprConstants []interface{}
		for _, pair := range values {
			input += fmt.Sprintf("%f %s %f;\n", pair[0], op, pair[1])
			exprConstants = append(exprConstants, evaltree.NewMathExpressionEvaluator(
				evaltree.NewLiteralEvaluator(pair[0]),
				evaltree.NewLiteralEvaluator(pair[1]),
				op,
			))
		}

		p, err := parseProgram(input)
		if err != nil {
			t.Errorf("Error parsing DML: %v for op %s", err, op)
			return
		}
		if len(p.Statements) != len(exprConstants) {
			t.Errorf("Expected %d statements, got %d for op %s", len(exprConstants), len(p.Statements), op)
			return
		}
		for i, statement := range p.Statements {
			if !reflect.DeepEqual(statement, exprConstants[i]) {
				t.Errorf("Expected %v, got %v for op %s", exprConstants[i], statement, op)
				continue
			}
		}
	}

}

func TestFunctionCallStatements(t *testing.T) {
	dmls := []string{"f();", "g(a: 1);", "h(a: 1, b_2: 2);", "i(foo: 1, bar: 2, bz: 3);", "j(a: 1, b: 2.0, c:\"Hello\", d: 0.4);"}
	dml := strings.Join(dmls, "\n")
	fnCallsEvaluators := []evaltree.Evaluator{
		evaltree.NewFunctionCallEvaluator(evaltree.NewIdentifierEvaluator("f"), nil),
		evaltree.NewFunctionCallEvaluator(evaltree.NewIdentifierEvaluator("g"), map[string]evaltree.Evaluator{
			"a": evaltree.NewLiteralEvaluator(1)}),
		evaltree.NewFunctionCallEvaluator(evaltree.NewIdentifierEvaluator("h"), map[string]evaltree.Evaluator{
			"a":   evaltree.NewLiteralEvaluator(1),
			"b_2": evaltree.NewLiteralEvaluator(2),
		}),
		evaltree.NewFunctionCallEvaluator(evaltree.NewIdentifierEvaluator("i"), map[string]evaltree.Evaluator{
			"foo": evaltree.NewLiteralEvaluator(1),
			"bar": evaltree.NewLiteralEvaluator(2),
			"bz":  evaltree.NewLiteralEvaluator(3),
		}),
		evaltree.NewFunctionCallEvaluator(evaltree.NewIdentifierEvaluator("j"), map[string]evaltree.Evaluator{
			"a": evaltree.NewLiteralEvaluator(1),
			"b": evaltree.NewLiteralEvaluator(2.0),
			"c": evaltree.NewLiteralEvaluator("Hello"),
			"d": evaltree.NewLiteralEvaluator(0.4),
		}),
	}
	p, err := parseProgram(dml)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != len(fnCallsEvaluators) {
		t.Errorf("Expected %d statements, got %d", len(fnCallsEvaluators), len(p.Statements))
		return
	}
	for i, statement := range p.Statements {
		if !reflect.DeepEqual(statement, fnCallsEvaluators[i]) {
			t.Errorf("Expected %v, got %v for fn with index %d", fnCallsEvaluators[i], statement, i)
			continue
		}
	}
}

func TestAssignmentStatements(t *testing.T) {
	dmls := []string{"a = 1;", "b = 2.0;", "c = \"Hello\";", "d = 0.4;",
		"m=2+3^4;",
		`computation test() {
		100
	}
	e = test;
`}
	assignEvaluators := []evaltree.Evaluator{
		evaltree.NewAssignmentEvaluator("a", evaltree.NewLiteralEvaluator(1)),
		evaltree.NewAssignmentEvaluator("b", evaltree.NewLiteralEvaluator(2)),
		evaltree.NewAssignmentEvaluator("c", evaltree.NewLiteralEvaluator("Hello")),
		evaltree.NewAssignmentEvaluator("d", evaltree.NewLiteralEvaluator(0.4)),
		evaltree.NewAssignmentEvaluator("m",
			evaltree.NewMathExpressionEvaluator(evaltree.NewLiteralEvaluator(2),
				evaltree.NewMathExpressionEvaluator(
					evaltree.NewLiteralEvaluator(3),
					evaltree.NewLiteralEvaluator(4),
					"^",
				),
				"+",
			)),
		evaltree.NewAssignmentEvaluator("e", evaltree.NewIdentifierEvaluator("test")),
	}
	for i := 0; i < len(dmls); i++ {
		p, err := parseProgram(dmls[i])
		if err != nil {
			t.Errorf("Error parsing DML: %v at index %d", err, i)
			return
		}
		lastStmt := p.Statements[len(p.Statements)-1]
		if !reflect.DeepEqual(lastStmt, assignEvaluators[i]) {
			t.Errorf("Expected %v, got %v for assignment with index %d", assignEvaluators[i], lastStmt, i)
			continue
		}
	}
}

func TestComputationDefinitionStatements(t *testing.T) {
	dml := `
	computation test(a: length (km), b: count) {
		c = "Hello";
		d = 0.4;
		m=a*b-d;
		m
	}
`
	computationDefEvaluator := evaltree.NewFunctionDefinitionEvaluator(
		"test",
		map[string]evaltree.ParameterDefinition{
			"a": evaltree.NewParameterDefinition("a", "length", "km"),
			"b": evaltree.NewParameterDefinition("b", "count", ""),
		},
		evaltree.NewBlockEvaluator(
			[]evaltree.Evaluator{
				evaltree.NewAssignmentEvaluator("c", evaltree.NewLiteralEvaluator("Hello")),
				evaltree.NewAssignmentEvaluator("d", evaltree.NewLiteralEvaluator(0.4)),
				evaltree.NewAssignmentEvaluator("m",
					evaltree.NewMathExpressionEvaluator(
						evaltree.NewMathExpressionEvaluator(
							evaltree.NewIdentifierEvaluator("a"),
							evaltree.NewIdentifierEvaluator("b"),
							"*",
						), evaltree.NewIdentifierEvaluator("d"), "-")),
				evaltree.NewIdentifierEvaluator("m"),
			},
		),
	)
	p, err := parseProgram(dml)
	if err != nil {
		t.Errorf("Error parsing DML: %v", err)
		return
	}
	if len(p.Statements) != 1 {
		t.Errorf("Expected 1 statement, got %d", len(p.Statements))
		return
	}
	lastStmt := p.Statements[0]
	if !reflect.DeepEqual(lastStmt, computationDefEvaluator) {
		t.Errorf("Expected %v, got %v for computation definition", computationDefEvaluator, lastStmt)
		return
	}
}
