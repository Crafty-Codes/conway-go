package object

import "fmt"

type Space [][]Block

func NewSpace(row int, column int) [][]Block {
	space := make([][]Block, row)

	for r := range space {
		space[r] = make([]Block, column)
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
