package main

import "fmt"

const (
	_ = 2022 + iota
	Y1
	Y2
	Y3
	Y4
)

func main() {
	fmt.Println(Y1, Y2, Y3, Y4)
}
