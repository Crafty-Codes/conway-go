package main

import (
	"conway/go/object"
	"fmt"
)

func main() {
	space := object.NewSpace(5)

	for i := 0; i < 5; i++ {
		print(space)
		fmt.Println()

		space = survival(space)
	}
}

func print(space object.Space) {
	for r := range space {
		for c := range space[r] {
			if space[r][c].Vitality == object.DEAD {
				fmt.Print("0")
			} else {
				fmt.Print("1")
			}
		}
		fmt.Println()
	}
}

func survival(space object.Space) object.Space {
	newSpace := object.NewSpace(len(space))

	for r := range space {
		for c := range space[r] {
			life := getSurvivalCount(space, r, c)

			if life == 3 {
				newSpace[r][c] = object.Block{Vitality: object.ALIVE}
			} else if space[r][c].Vitality == object.ALIVE && life == 2 {
				newSpace[r][c] = object.Block{Vitality: object.ALIVE}
			}
		}
	}

	return newSpace
}

func getSurvivalCount(space object.Space, row int, column int) uint8 {
	var counter uint8 = 0

	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if r != row || c != column {
				if (r >= 0 && r < len(space)) && (c >= 0 && c < len(space)) {
					if space[r][c].Vitality == object.ALIVE {
						counter++
					}
				}
			}
		}
	}

	return counter
}
