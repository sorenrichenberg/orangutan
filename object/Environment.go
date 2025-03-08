package object

type Environment struct {
	storage map[string]Object
	outer   *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{storage: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	extended := NewEnvironment()
	extended.outer = outer
	return extended
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.storage[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.storage[name] = val
	return val
}
