package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Chara-X/scripting"
)

func main() {
	var storage = map[any]any{}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".cs" {
			var source, _ = os.ReadFile(path)
			script.Compile(string(source)).Run(storage)
		}
		return err
	})
	fmt.Println(storage["main"].(func(args ...any) any)())
}
