/*
- Crie e utilize uma função anônima.
*/

package main

import "fmt"

func main() {
	fmt.Println(func() string {
		return "Hello world"
	}())
}
