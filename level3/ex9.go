package main

import "fmt"

func main() {
	favoriteSport := "soccer"

	switch favoriteSport {

	case "basketball":
		fmt.Println("Basket")
	case "football":
		fmt.Println("Touchdown")
	case "soccer":
		fmt.Println("Goal")
	default:
		fmt.Println("Unknown sport")
	}
}
