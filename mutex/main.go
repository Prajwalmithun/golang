/*
	Code explains,
	Problem = Race condition, when multiple goroutines accessing the shared variable.
	Solution = Mutex (Locking while operation is happening. Unlocking after the operation is done.)

	Useful command to check if our code has any race condition
	Command:
	$ go run --race main.go

	If the output of the above command has the following error, then there is race condition in the code.
	WARNING: DATA RACE
	exit status 66

*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	// wait group
	wg := &sync.WaitGroup{}
	// mutex: solution for race condition. Locking the operation when 1 subroutine performing the operation
	mut := &sync.Mutex{}

	var shared_value = 0

	// adding goroutines to the wait group
	wg.Add(3)

	// ananymous go routines
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Printf("Increment\n")

		// locking
		mut.Lock()
		shared_value = shared_value + 1

		fmt.Printf("value of shared variable is %v\n\n", shared_value)

		// Unlocking
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Printf("Decrement\n")

		// locking
		mut.Lock()
		shared_value = shared_value - 1
		// unlocking
		fmt.Printf("value of shared variable is %v\n\n", shared_value)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Printf("Muliplication\n")

		// locking
		mut.Lock()
		shared_value = shared_value * 1
		// unlocking
		fmt.Printf("value of shared variable is %v\n\n", shared_value)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	// go func(wg *sync.WaitGroup, mut *sync.Mutex){
	// 	fmt.Printf("Nothing\n")
	// 	fmt.Printf("value of shared variable is %v\n",shared_value)
	// 	wg.Done()
	// }(wg,mut)

	// waiting untill all the goroutines finish their execution
	wg.Wait()

	//mut.Lock()
	fmt.Printf("Value after all computation %v\n", shared_value)
	//mut.Unlock()
}
