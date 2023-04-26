package main

import (
	"conway/go/logic"
	"conway/go/object"
	"fmt"
	"time"
)

func main() {
	space := object.NewSpace(5)

	for true {
		print(space)
		fmt.Println()

		space = logic.Survival(space)

		time.Sleep(1 * time.Second)
	}
}

func print(space object.Space) {
	for r := range space {
		for c := range space[r] {
			space[r][c].Print()
		}
		fmt.Println()
	}
}
