/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import "fmt"

func main() {
	x := struct {
		mapa  map[string]int
		fatia []int
	}{
		mapa: map[string]int{
			"tres":  3,
			"cinco": 5,
		},
		fatia: []int{4},
	}

	fmt.Println(x)
}
