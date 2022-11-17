/*
- Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var m sync.Mutex
	wg := sync.WaitGroup{}

	contador := 0
	total := 1000
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			m.Lock()
			x := contador
			runtime.Gosched()
			x++
			contador = x
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
