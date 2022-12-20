package main

import "fmt"

func main() {
	bornYear := 1995
	for {
		fmt.Println(bornYear)
		bornYear++
		if bornYear > 2022 {
			break
		}

	}
}
