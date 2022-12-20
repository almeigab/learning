package main

import "fmt"

func main() {

	i := 11

	switch {
	case i < 10:
		fmt.Println("less than 10")
	case i > 10:
		fmt.Println("greater than 10")
	default:
		fmt.Println("equals 10")
	}
}
