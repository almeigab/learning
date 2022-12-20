/*
- Callback: passe uma função como argumento a outra função.
*/

package main

import "fmt"

func funcao(funcaoInterna func() string) string {
	return fmt.Sprintf("resultado: %v", funcaoInterna())
}

func helloWorld() string {
	return "Hello world"
}

func main() {

	fmt.Println(funcao(helloWorld))
}
