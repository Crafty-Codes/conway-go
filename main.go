package main

import (
	"conway/go/object"
	"fmt"
)

func main() {
	space := object.NewSpace(5)

	print(space)

}

func print(space [][]object.Block) {
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
