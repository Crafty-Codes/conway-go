package main

import (
	"conway/go/logic"
	"conway/go/object"
	"fmt"
	"time"
)

func main() {
	space := object.NewSpace(5)

	space[2][1].Vitality = object.ALIVE
	space[2][2].Vitality = object.ALIVE
	space[2][3].Vitality = object.ALIVE

	for true {
		object.Print(space)
		fmt.Println()

		space = logic.Survival(space)

		time.Sleep(1 * time.Second)
	}
}
