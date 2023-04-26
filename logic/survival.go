package logic

import "conway/go/object"

func Survival(space object.Space) object.Space {
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
