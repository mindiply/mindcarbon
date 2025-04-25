package evaltree

type NativeFunction func(params map[string]Object) Object

var nativeFunctions = map[string]NativeFunction{
	"lookupFactorFor": func(params map[string]Object) Object {
		return NewNumberObject(0)
	},
	"anotherFunction": func(params map[string]Object) Object {
		return NewStringObject("example")
	},
}

func GetNativeFunctions() map[string]NativeFunction {
	return nativeFunctions
}
