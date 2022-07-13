package main

import "fmt"

func main() {
	x := [][]string{
		{"Gabriel", "Almeida", "Plays guitar"},
		{"Julio", "Cocorico", "Plays harmonica"},
		{"Yugi", "Oh", "Plays cards"},
	}

	for _, row := range x {
		for _, value := range row {
			fmt.Println(value)
		}
		fmt.Println()
	}

}
