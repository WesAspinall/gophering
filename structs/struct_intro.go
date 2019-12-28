package main

import "fmt"

// composing values of different types
type person struct {
	first string
	last  string
}

func main() {
	// "created a value p1 of type person"
	// approximation of an "instance"
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
