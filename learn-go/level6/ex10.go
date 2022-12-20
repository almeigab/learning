/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

package main

import "fmt"

func main() {
	func() {
		i := 0
		fmt.Printf("fora antes: %v\n", i)
		func() {
			i += 1
			fmt.Printf("dentro: %v\n", i)
		}()
		fmt.Printf("fora depois: %v\n", i)
	}()
}
