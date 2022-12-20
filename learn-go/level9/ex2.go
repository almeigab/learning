/*
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
- Se precisar de dicas, veja: https://gobyexample.com/interfaces
*/

package main

import "fmt"

type pessoa struct {
	nome string
}

func (p *pessoa) falar() {
	fmt.Printf("Oi, eu sou %v!\n", p.nome)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(hs humanos) {
	hs.falar()
}

func main() {
	p1 := pessoa{"Gabriel"}
	dizerAlgumaCoisa(&p1)

	/*
	   discomente a linha abaixo para verificar
	   que você não pode utilizar um valor do tipo "pessoa"
	   na função "dizerAlgumaCoisa"
	*/
	// dizerAlgumaCoisa(p1)

}
