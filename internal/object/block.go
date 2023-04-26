package object

import "fmt"

type Block struct {
	Vitality Vital
}

type Vital int

const (
	ALIVE Vital = 0
	DEAD  Vital = 1
)

func (block Block) Print() {
	if block.Vitality == ALIVE {
		fmt.Print("░")
	} else {
		fmt.Print("█")
	}
}
