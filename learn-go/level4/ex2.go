package main

import "fmt"

func main() {
	slice := []int{2, 5, 3, 7, 1, 10, 3, 4, 2, 0}

	for _, value := range slice {
		fmt.Print(value, " ")
	}

	fmt.Printf("\n%T\n", slice)
}
