package eval

import (
	"github.com/mindiply/mindcarbon/rules/evaltree"
	"strings"
	"testing"
)

func TestNewInterpreter(t *testing.T) {
	i := NewInterpreter()
	if i == nil {
		t.Errorf("NewInterpreter() returned nil")
		return
	}
	if i.RootEnv() == nil {
		t.Errorf("RootEnv() returned nil")
		return
	}
	_, err := i.EvaluateDML("Not a DML statement")
	expectedErr := "errors during parsing"
	if !(err != nil && strings.HasPrefix(err.Error(), expectedErr)) {
		errorMsg := ""
		if err != nil {
			errorMsg = err.Error()
		}
		t.Errorf("Expected %s, got %s", expectedErr, errorMsg)
	}

	v, err := i.EvaluateDML("")
	if err != nil {
		t.Errorf("Excpected no error with an ampty DML statement, got %s", err.Error())
		return
	}
	if v.Type() != evaltree.NULLTYPE {
		t.Errorf("Expected NULL type, got %s", v.Type())
		return
	}

	_, err = i.EvaluateDML("NULL+2;")
	expectedErr = "execution error: You can only perform math operations on numbers"
	if err == nil {
		t.Errorf(`Expected an error with "NULL+2", got nil`)
	} else if err.Error() != expectedErr {
		t.Errorf("Expected %s, got %s", expectedErr, err.Error())
	}
}

func TestScalars(t *testing.T) {
	tests := []struct {
		dml string
		obj evaltree.Object
	}{
		{"1;", evaltree.NewNumberObject(1)},
		{"1.0;", evaltree.NewNumberObject(1.0)},
		{"-1;", evaltree.NewNumberObject(-1)},
		{`"hello";`, evaltree.NewStringObject("hello")},
		{`"\"hello\"";`, evaltree.NewStringObject("\\\"hello\\\"")},
		{"\"Hi!\nHow's it hanging\t\t?\";", evaltree.NewStringObject("Hi!\nHow's it hanging\t\t?")},
	}
	runTests(t, tests)
}

func TestMath(t *testing.T) {
	tests := []struct {
		dml string
		obj evaltree.Object
	}{
		{"1 + 2;", evaltree.NewNumberObject(3)},
		{"1 - 2;", evaltree.NewNumberObject(-1)},
		{"1 * 2;", evaltree.NewNumberObject(2)},
		{"1 / 2;", evaltree.NewNumberObject(0.5)},
		{"1 + 2 * 3;", evaltree.NewNumberObject(7)},
		{"1 + 2 - 3;", evaltree.NewNumberObject(0)},
		{"1 + 2 / 3;", evaltree.NewNumberObject(1 + (2 / 3.0))},
		{"2^3;", evaltree.NewNumberObject(8)},
		{"2^3^2;", evaltree.NewNumberObject(512)},
		{"(2^3)^2;", evaltree.NewNumberObject(64)},
		{"(0);", evaltree.NewNumberObject(0)},
		{"((((-1000))))*2;", evaltree.NewNumberObject(-2000)},
	}
	runTests(t, tests)
}

func TestNonFuncAssignments(t *testing.T) {
	tests := []struct {
		dml string
		obj evaltree.Object
	}{
		{"a = 1;a;", evaltree.NewNumberObject(1)},
		{"a = 1.0;a;", evaltree.NewNumberObject(1.0)},
		{"a = -1;b = a;b;", evaltree.NewNumberObject(-1)},
		{"a = 2/2;b = 3+a;b;", evaltree.NewNumberObject(4)},
		{`a = "hello";a;`, evaltree.NewStringObject("hello")},
	}
	runTests(t, tests)
}

func TestRuntimeErrors(t *testing.T) {
	tests := [][]string{
		{`a = "hello";b=2;b+a;`, "execution error: You can only perform math operations on numbers"},
		{`a = "hello";b=a;b+a;`, "execution error: You can only perform math operations on numbers"},
	}
	i := NewInterpreter()
	for _, test := range tests {
		v, e := i.EvaluateDML(test[0])
		if v.Type() != evaltree.ERRORTYPE {
			t.Errorf("Expected error type, got %s", v.Type())
		}
		if e == nil {
			t.Errorf("Expected error with %q, got nil", test[0])
		} else if e.Error() != test[1] {
			t.Errorf("Expected error message \"%q\", got \"%q\"", test[1], e.Error())
		}
	}
}

func TestSimpleFunctions(t *testing.T) {
	tests := []struct {
		dml string
		obj evaltree.Object
	}{
		{"computation test() { 1 }\n\ntest();", evaltree.NewNumberObject(1)},
		{"computation test() { \"Hello\" }\n\ntest();", evaltree.NewStringObject("Hello")},
		{"computation a() {a=2;b=3.2;a*b} computation test() { a() }\n\ntest();", evaltree.NewNumberObject(6.4)},
	}
	runTests(t, tests)
}

func TestFunctionVariables(t *testing.T) {
	tests := []struct {
		dml string
		obj evaltree.Object
	}{
		{"computation test() { 1 } a=test;\n\nb=test();b;", evaltree.NewNumberObject(1)},
		{"computation test() { \"Hello\" }\na=test;\na();", evaltree.NewStringObject("Hello")},
		{"computation a() {a=2;b=3.2;a*b} computation test() { b=a; b() }\n\ntest();", evaltree.NewNumberObject(6.4)},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests []struct {
	dml string
	obj evaltree.Object
}) {
	i := NewInterpreter()
	for _, test := range tests {
		v, e := i.EvaluateDML(test.dml)
		if e != nil {
			t.Errorf("Error evaluating DML %q: %v", test.dml, e)
			continue
		}
		if v == nil {
			t.Errorf("EvaluateDML(%q) returned nil", test.dml)
			continue
		}
		if v.Type() != test.obj.Type() {
			t.Errorf("EvaluateDML(%q) returned %T, expected %T", test.dml, v, test.obj)
			continue
		}
		if v.Inspect() != test.obj.Inspect() {
			t.Errorf("EvaluateDML(%q) returned %q, expected %q", test.dml, v.Inspect(), test.obj.Inspect())
			continue
		}
	}
}
