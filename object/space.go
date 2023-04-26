package object

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
