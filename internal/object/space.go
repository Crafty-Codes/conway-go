package object

import "fmt"

type Space [][]Block

func NewSpace(row int, column int) Space {
	space := make(Space, row)

	for r := range space {
		space[r] = make([]Block, column)
		for c := range space[r] {
			space[r][c] = Block{Vitality: DEAD}
		}
	}

	return space
}

func Print(space Space) {
	printString := ""
	for _, r := range space {
		for _, block := range r {
			printString = printString + string(block.Vitality)
		}
		printString = printString + "\n"
	}
	fmt.Println(printString)
}
