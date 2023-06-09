package logic

import "github.com/Crafty-Codes/conway-go/internal/object"

func Survival(space object.Space) object.Space {
	newSpace := make(object.Space, len(space))

	for r := range space {
		newSpace[r] = make([]object.Block, len(space[r]))
		for c := range space[r] {
			life := getSurvivalCount(space, r, c)

			if life == 3 {
				newSpace[r][c] = object.Block{Vitality: object.ALIVE}
			} else if space[r][c].Vitality == object.ALIVE && life == 2 {
				newSpace[r][c] = object.Block{Vitality: object.ALIVE}
			} else if life == 9 {
				newSpace[r][c] = object.Block{Vitality: object.WALL}
			} else {
				newSpace[r][c] = object.Block{Vitality: object.DEAD}
			}
		}
	}

	return newSpace
}

func getSurvivalCount(space object.Space, row int, column int) uint8 {
	var counter uint8 = 0
	// if a WALL was detected it will return 9
	// 9 was choosen because it will not be possible to have 9 living neighbours
	if space[row][column].Vitality == object.WALL {
		return 9
	}

	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if r != row || c != column {
				if (r >= 0 && r < len(space)) && (c >= 0 && c < len(space[r])) {
					if space[r][c].Vitality == object.ALIVE {
						counter++
					}
				}
			}
		}
	}

	return counter
}
