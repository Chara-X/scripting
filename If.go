package script

type If struct{ Condition, Body Script }

func (i *If) Run(storage map[any]any) any {
	if i.Condition.Run(storage).(bool) {
		return i.Body.Run(storage)
	}
	return nil
}
