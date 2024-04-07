package script

type Selector struct {
	Value Script
	Key   string
}

func (s *Selector) Run(storage map[any]any) any { return s.Value.Run(storage).(map[any]any)[s.Key] }
