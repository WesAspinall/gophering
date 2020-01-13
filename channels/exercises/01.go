package main

import (
	"fmt"
)

// get this code working
// func main() {
// 	c := make(chan int)

// 	c <- 42

// 	fmt.Println(<-c)
// }

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
