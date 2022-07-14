/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) demonstraPessoa() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func main() {
	p := pessoa{"Gabriel", "Almeida", 27}
	p.demonstraPessoa()
}
