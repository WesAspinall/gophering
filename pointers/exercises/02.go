package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Tom"
	fmt.Println(p)
}

func main() {
	p1 := person{
		"Bob",
		"Saget",
		53,
	}

	changeMe(&p1)
	fmt.Println(p1.first)
}
