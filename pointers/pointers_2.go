package main

import "fmt"

func main() {
	a := 4
	fmt.Println(a)
	fmt.Println(&a)
	mutate(&a) // pass in the mem address
	fmt.Println(a)
	fmt.Println(&a)
}

func mutate(x *int) {
	*x = 42 // dereference
}
