package script

import "maps"

type New struct{ Body Script }

func (n *New) Run(storage map[any]any) any {
	var scope = maps.Clone(os)
	scope["this"] = storage
	n.Body.Run(scope)
	return scope
}
