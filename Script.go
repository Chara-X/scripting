package script

import (
	"encoding/json"
)

type Script interface{ Run(storage map[any]any) any }

func Compile(source string) Script { return compile(&reader{strings: lexer.FindAllString(source, -1)}) }
func compile(reader *reader) Script {
	var script Script
	switch prefix := reader.Read(); prefix {
	case "func":
		var keys []string
		reader.Read()
		for reader.Read() != ")" {
			reader.Unread()
			keys = append(keys, reader.Read())
		}
		return &Func{keys, compile(reader)}
	case "new":
		return &New{compile(reader)}
	case "if":
		return &If{compile(reader), compile(reader)}
	case "for":
		return &For{compile(reader), compile(reader)}
	case "{":
		var children []Script
		for reader.Read() != "}" {
			reader.Unread()
			children = append(children, compile(reader))
		}
		return &Block{children}
	default:
		var value any
		if json.Unmarshal([]byte(prefix), &value) == nil {
			return &Constant{value}
		}
		script = &Variable{prefix}
	}
	for {
		switch reader.Read() {
		case ".":
			script = &Selector{script, reader.Read()}
		case "=":
			return &Store{script, compile(reader)}
		case "(":
			var args []Script
			for reader.Read() != ")" {
				reader.Unread()
				args = append(args, compile(reader))
			}
			script = &Call{script, args}
		default:
			reader.Unread()
			return script
		}
	}
}

type reader struct {
	strings []string
	offset  int
}

func (r *reader) Unread() { r.offset-- }
func (r *reader) Read() string {
	r.offset++
	return r.strings[r.offset-1]
}
