package main

import "fmt"

// pointers is the memory address where the value is stored
// pointers are good to use if you have a large chunk of data... you can just to pass the address
//if you need to change something that's at a certain location

func main() {
	a := 4
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a) // type is a pointer to an int

	b := &a
	fmt.Println(b) //same address as a

	*b = 42 // dereferencing

	fmt.Println(a) //the value at a is now 42

}
