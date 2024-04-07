package script

type Store struct {
	Key, Value Script
}

func (s *Store) Run(storage map[any]any) any {
	switch key := s.Key.(type) {
	case *Selector:
		key.Value.Run(storage).(map[any]any)[key.Key] = s.Value.Run(storage)
	case *Variable:
		storage[key.Key] = s.Value.Run(storage)
	}
	return nil
}
