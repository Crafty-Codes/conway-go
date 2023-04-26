package main

import (
	"fmt"
	"time"

	"github.com/Crafty-Codes/conway-go/internal/logic"
	"github.com/Crafty-Codes/conway-go/internal/object"
)

func main() {
	space := object.NewSpace(5)

	space[2][1].Vitality = object.ALIVE
	space[2][2].Vitality = object.ALIVE
	space[2][3].Vitality = object.ALIVE

	for {
		object.Print(space)
		fmt.Println()

		space = logic.Survival(space)

		time.Sleep(1 * time.Second)
	}
}
