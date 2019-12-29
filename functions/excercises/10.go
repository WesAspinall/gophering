package main

import "fmt"

// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:

func main() {
	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
