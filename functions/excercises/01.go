package main

import "fmt"

// Hands on exercise
// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

func main() {
	f := foo()
	x, y := bar()

	fmt.Println(f)
	fmt.Println(x, y)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	x := 43
	y := "bar"

	return x, y
}
