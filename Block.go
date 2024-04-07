package script

type Block struct{ Statements []Script }

func (b *Block) Run(storage map[any]any) any {
	for _, child := range b.Statements {
		if res := child.Run(storage); res != nil {
			return res
		}
	}
	return nil
}
