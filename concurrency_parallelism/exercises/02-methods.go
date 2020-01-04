package main

import "fmt"

// This exercise will reinforce our understanding of method sets:
// create a type person struct
// attach a method speak to type person using a pointer receiver
// *person
// create a type human interface
// to implicitly implement the interface, a human must have the speak method
// create func “saySomething”
// have it take in a human as a parameter
// have it call the speak method
// show the following in your code
// you CAN pass a value of type *person into saySomething
// you CANNOT pass a value of type person into saySomething
// here is a hint if you need some help
// https://play.golang.org/p/FAwcQbNtMG

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() string {
	a := fmt.Sprint("Hello, my name is ", p.first, " ", p.last, ", and I am ", p.age, " years old.")
	return a
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		first: "James",
		last:  "Taylor",
		age:   71,
	}

	saySomething(&p1)
	p1.speak()
}
