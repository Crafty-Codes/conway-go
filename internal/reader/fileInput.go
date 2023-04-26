package reader

import (
	"bufio"
	"log"
	"os"

	"github.com/Crafty-Codes/conway-go/internal/object"
)

func ReadFile() object.Space {
	// open file
	f, err := os.Open("map.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var scannedList object.Space
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]object.Block, len(text))
		for i, v := range scanner.Text() {
			if v == '0' {
				row[i].Vitality = object.DEAD
			} else if v == '1' {
				row[i].Vitality = object.ALIVE
			} else {
				row[i].Vitality = object.WALL
			}
		}
		scannedList = append(scannedList, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scannedList
}
