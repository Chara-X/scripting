package script

import (
	"reflect"
)

type Call struct {
	Value Script
	Args  []Script
}

func (c *Call) Run(storage map[any]any) any {
	var args []reflect.Value
	for _, arg := range c.Args {
		args = append(args, reflect.ValueOf(arg.Run(storage)))
	}
	if res := reflect.ValueOf(c.Value.Run(storage)).Call(args); len(res) > 0 {
		return res[0].Interface()
	}
	return nil
}
