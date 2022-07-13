package main

import "fmt"

func main() {
	x := map[string][]string{
		"Almeida_Gabriel": {
			"Plays guitar",
			"Watch movies",
			"Ride a bike",
		},
		"Cocorico_Julio": {
			"Plays harmonica",
			"Farming",
			"Listen to country rock music",
		},
	}

	for index, value := range x {
		fmt.Println(index)
		for i, v := range value {
			fmt.Println("\t", i, " - ", v)
		}
	}
}
