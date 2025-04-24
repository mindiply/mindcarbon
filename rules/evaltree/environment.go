package evaltree

// Environment is an interface that represents a mapping from string keys to evauluators.
// A key will represent a named object, like a variable or a function.
// The environment can be used to look up the Value associated with a key, or to set a new Value for a key.
// The lookup will be done in the current environment, and if the key is not found, it will look up in the parent environment.
type Environment interface {
	// Get retrieves the Value associated with the given key.
	Get(key string) (Evaluator, bool)

	// Set associates the given key with the Value.
	Set(key string, value Evaluator)
}

type EnvironmentObject struct {
	namedObjects map[string]Evaluator
	parentEnv    Environment
}

func (e *EnvironmentObject) Get(key string) (Evaluator, bool) {
	if obj, ok := e.namedObjects[key]; ok {
		return obj, true
	}
	if e.parentEnv != nil {
		return e.parentEnv.Get(key)
	}
	return nil, false
}

func (e *EnvironmentObject) Set(key string, value Evaluator) {
	e.namedObjects[key] = value
}

func NewEnvironment(parent Environment) Environment {
	return &EnvironmentObject{
		namedObjects: make(map[string]Evaluator),
		parentEnv:    parent,
	}
}
