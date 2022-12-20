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

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.sorvetesFavoritos {
		fmt.Println("\t -", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.sorvetesFavoritos {
		fmt.Println("\t -", v)
	}
}
