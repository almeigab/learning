/*
- Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	wg := sync.WaitGroup{}

	contador := int32(0)
	total := 1000
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {

			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
