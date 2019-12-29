package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
	}() //anonymous func

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42) //anonymous func executes with
}

func foo() {
	fmt.Println("hello from foo")
}
