package evaltree

type NativeFunction func(params map[string]Object) Object

var LookupFactorFor NativeFunction = func(params map[string]Object) Object {
	return NewNumberObject(0)
}
