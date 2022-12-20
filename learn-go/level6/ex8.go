/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/

package main

import "fmt"

func funcao() func() string {
	return func() string {
		return "Hello world"
	}
}

func main() {
	variavel := funcao()

	fmt.Println(variavel())
}
