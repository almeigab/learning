package main

import "fmt"

type pessoa struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {
	pessoa1 := pessoa{
		nome:              "Gabriel",
		sobrenome:         "Almeida",
		sorvetesFavoritos: []string{"chocolate", "doce de leite"},
	}

	pessoa2 := pessoa{
		nome:              "Beatle",
		sobrenome:         "Juice",
		sorvetesFavoritos: []string{"flocos", "pistache"},
	}

	pessoas := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for _, p := range pessoas {
		fmt.Println(p.nome, p.sobrenome)

		for _, v := range p.sorvetesFavoritos {
			fmt.Println("\t -", v)
		}
	}
}
