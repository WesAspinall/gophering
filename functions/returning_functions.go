package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T", x)
	i := x()
	fmt.Printf("\n%v\n", i)

	f := foo()
	fmt.Println(f())
}

func bar() func() int {
	return func() int {
		return 451
	}
}

func foo() func() int {
	return func() int {
		return 452
	}
}
