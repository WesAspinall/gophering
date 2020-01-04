package main

// task: create race condition

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	const gs = 50

	wg.Add(50)

	//number of goroutines
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Printf("%v: %v\n", i, counter)
			wg.Done()
		}()
	}
	fmt.Println(counter) // should be 50
}
