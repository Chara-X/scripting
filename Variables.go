package script

import (
	"fmt"
	"regexp"
	"time"
)

var (
	lexer = regexp.MustCompile(`[A-Za-z]\w*|-?\d+(\.\d+)?|".*?"|\(|\)|{|}|\.|=`)
	os    = map[any]any{
		"return": func(res any) any { return res },
		"break":  func() complex64 { return complex64(0) },
		"eq":     func(x, y any) bool { return x == y },
		"gt":     func(x, y float64) bool { return x > y },
		"lt":     func(x, y float64) bool { return x < y },
		"add":    func(x, y float64) any { return x + y },
		"go":     func(f func(args ...any) any) { go f() },
		"delete": func(m map[any]any, key any) { delete(m, key) },
		"slice":  func(elems ...any) []any { return elems },
		"append": func(slice []any, elems ...any) []any { return append(slice, elems...) },
		"set":    func(slice []any, index float64, elem any) { slice[int(index)] = elem },
		"get":    func(slice []any, index float64) any { return slice[int(index)] },
		"chan":   func(cap float64) chan any { return make(chan any, int(cap)) },
		"send":   func(c chan any, val any) { c <- val },
		"recv":   func(c chan any) any { return <-c },
		"close":  func(c chan any) { close(c) },
		"len": func(v any) float64 {
			switch v := v.(type) {
			case map[any]any:
				return float64(len(v))
			case []any:
				return float64(len(v))
			case chan any:
				return float64(len(v))
			case string:
				return float64(len(v))
			}
			panic(0)
		},
		"fmt": map[any]any{
			"Println": func(args ...any) { fmt.Println(args...) },
			"Scanln": func() string {
				var arg string
				fmt.Scanln(&arg)
				return arg
			},
		},
		"time": map[any]any{
			"Now": func() any {
				var time = time.Now()
				return map[any]any{
					"Year":   time.Year,
					"Format": time.Format,
				}
			},
		},
	}
)
