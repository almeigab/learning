/*
- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	contador := 0
	total := 1000
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			x := contador
			runtime.Gosched()
			x++
			contador = x
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
