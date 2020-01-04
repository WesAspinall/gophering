package main

// task: fix race condition with mutex

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	counter := 0
	const gs = 50
	wg.Add(50)

	var mu sync.Mutex

	//number of goroutines
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Printf("Counter: %v\n", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("The final value of the counter is:", counter) // should be 50
}
