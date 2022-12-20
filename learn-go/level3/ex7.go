package main

import "fmt"

func main() {
	i := 10

	if i < 10 {
		fmt.Println("less than 10")
	} else if i > 10 {
		fmt.Println("greater than 10")
	} else {
		fmt.Println("equals 10")
	}
}
