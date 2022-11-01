/*
- Crie e utilize uma função anônima.
*/

package main

import "fmt"

func main() {
	helloWorld := func() string {
		return "Hello world"
	}

	fmt.Println(helloWorld())
}
