package main

// task: fix race condition using atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// remember - AddUInt64, LoadUInt64
// docs: https://golang.org/pkg/sync/atomic/#LoadUint64

func main() {
	var counter uint64
	var wg sync.WaitGroup
	const gs = 50

	wg.Add(50)

	//number of goroutines
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddUint64(&counter, 1)
			r := atomic.LoadUint64(&counter)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter) // should be 50
}
