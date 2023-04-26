package main

import (
	"time"

	"github.com/Crafty-Codes/conway-go/internal/logic"
	"github.com/Crafty-Codes/conway-go/internal/object"
	"github.com/Crafty-Codes/conway-go/internal/reader"
)

func main() {
	space := reader.ReadFile()

	for {
		object.Print(space)

		space = logic.Survival(space)

		time.Sleep(250 * time.Millisecond)
	}
}
