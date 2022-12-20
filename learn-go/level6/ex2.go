/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}

	s := sum(list...)
	s2 := sumSlice(list)

	fmt.Println(s, s2)
}

func sum(list ...int) int {
	s := 0
	for _, v := range list {
		s += v
	}

	return s
}

func sumSlice(list []int) int {
	s := 0
	for _, v := range list {
		s += v
	}

	return s
}
