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

	x["OH_Yugi"] = []string{
		"Play cards",
		"Duel",
	}

	delete(x, "Cocorico_Julio")

	for index, value := range x {
		fmt.Println(index)
		for i, v := range value {
			fmt.Println("\t", i, " - ", v)
		}
	}
}
