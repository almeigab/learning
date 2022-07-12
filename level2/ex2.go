package main

import "fmt"

func main() {
	x := 4
	y := 7

	eq := x == y
	ne := x != y
	lt := x < y
	lte := x <= y
	gt := x > y
	gte := x >= y

	fmt.Println(x, "equal", y, ":", eq)
	fmt.Println(x, "not equal", y, ":", ne)
	fmt.Println(x, "less than", y, ":", lt)
	fmt.Println(x, "less than or equal", y, ":", lte)
	fmt.Println(x, "greater than", y, ":", gt)
	fmt.Println(x, "greater than or equal", y, ":", gte)
}
