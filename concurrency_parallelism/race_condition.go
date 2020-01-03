package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// starts go routines
		go func() {
			v := counter
			// time.Spleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println("Goroutines:", runtime.NumGoroutine()) // count doesn't get updated as expected
		}() //race condition occurs when the counter is updated
	}

	//fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
