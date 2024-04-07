package script

type Constant struct{ Value any }

func (c *Constant) Run(storage map[any]any) any { return c.Value }
