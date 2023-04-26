package main

import (
	"fmt"
	"time"

	"github.com/Crafty-Codes/conway-go/internal/logic"
	"github.com/Crafty-Codes/conway-go/internal/object"
	"github.com/Crafty-Codes/conway-go/internal/reader"
)

func main() {
	space := reader.ReadFile()

	for {
		fmt.Println()
		object.Print(space)

		space = logic.Survival(space)

		time.Sleep(500 * time.Millisecond)
	}
}
