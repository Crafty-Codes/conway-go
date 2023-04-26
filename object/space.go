package object

import "fmt"

type plane interface {
	Print()
}

type Space [][]Block

func NewSpace(size int) [][]Block {
	space := make([][]Block, size)

	for r := range space {
		space[r] = make([]Block, size)
		for c := range space[r] {
			space[r][c] = Block{Vitality: DEAD}
		}
	}

	return space
}

func Print(space Space) {
	for _, r := range space {
		for _, block := range r {
			block.Print()
		}
		fmt.Println()
	}
}
