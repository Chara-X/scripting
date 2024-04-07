package script

import "maps"

type Func struct {
	Keys []string
	Body Script
}

func (f *Func) Run(storage map[any]any) any {
	return func(args ...any) any {
		var scope = maps.Clone(os)
		scope["this"] = storage
		for i, key := range f.Keys {
			scope[key] = args[i]
		}
		return f.Body.Run(scope)
	}
}
