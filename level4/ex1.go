package main

import "fmt"

func main() {
	arr := [5]int{2, 5, 3, 7, 1}

	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Printf("\n%T\n", arr)
}
