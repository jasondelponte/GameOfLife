package main

import (
	"time"
	"fmt"
	"flag"
)

var width = flag.Int("w", 15, "Width of the game")
var height = flag.Int("h", 15, "Height of the game")
var chance = flag.Int("c", 40, "Chance of life")

func main() {
	flag.Parse()

	life := NewLife(*height, *width, *chance, nil)

	ticker := time.NewTicker(250 * time.Millisecond)
	for {
		life.Step()
		fmt.Println(life)
		select {
		case <-ticker.C:
		}
	}
}
