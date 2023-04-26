package object

import "fmt"

type Block struct {
	Vitality Vital
}

type Vital string

const (
	ALIVE Vital = "░"
	DEAD  Vital = "█"
	WALL  Vital = "▓"
)

func (block Block) Print() {
	fmt.Print(block.Vitality)
}
