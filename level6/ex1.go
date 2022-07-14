/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/
package main

import "fmt"

func main() {
	fmt.Println(returnInt())
	fmt.Println(returnIntStr())
}

func returnInt() int {
	return 2
}

func returnIntStr() (i int, s string) {
	i = 4
	s = "abc"
	return
}
