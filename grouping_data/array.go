package main

import "fmt"

func main() {
	// arrays have length as a part of their type
	// they are not commonly used in go
	var x [5]int

	//set the first element to 42
	x[0] = 42

	//print first element
	fmt.Println(x[0])
}
