package main

import "fmt"

//Assign a func to a variable, then call that func

func main() {
	f := func(x int) int {
		return x + 2
	}

	fmt.Println(f(2))
}
