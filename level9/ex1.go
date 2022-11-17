/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import "sync"

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		println("Uma go-routine")
		wg.Done()
	}()

	go func() {
		go println("Outra go-routine")
		wg.Done()
	}()

	wg.Wait()
}
