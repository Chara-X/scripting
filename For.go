package script

type For struct {
	Condition, Body Script
}

func (f *For) Run(storage map[any]any) any {
	for f.Condition.Run(storage).(bool) {
		if res := f.Body.Run(storage); res != nil {
			if _, ok := res.(complex64); ok {
				return nil
			}
			return res
		}
	}
	return nil
}
