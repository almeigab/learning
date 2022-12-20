package main

import "fmt"

func main() {
	slice := []int{2, 5, 3, 7, 1, 10, 3, 4, 2, 0}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
