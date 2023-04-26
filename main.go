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
		object.Print(space)
		fmt.Println()

		space = logic.Survival(space)

		time.Sleep(1 * time.Second)
	}
}
