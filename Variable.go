package script

type Variable struct{ Key string }

func (v *Variable) Run(storage map[any]any) any { return storage[v.Key] }
