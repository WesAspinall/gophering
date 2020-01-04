package main

// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup // variable is of type WaitGroup from package sync
	wg.Add(2)             // add 2 wait groups

	go func() {
		fmt.Println("hello from foo")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from bar")
		wg.Done()
	}()
	fmt.Println("Middle Number of Goroutines :", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End Number of Goroutines :", runtime.NumGoroutine())
	fmt.Println("End Number of CPUs:", runtime.NumCPU())
	fmt.Println("about to exit")
}
