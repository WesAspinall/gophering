package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// keyword identifier type
type human interface {
	speak() // any type that has the method speak is also type human
} // "If you have this method, then you're my type"

// ** a value can be of more than one type

// func bar(h human) {
// 	fmt.Println("I was passed into bar", h)
// }

func bar(h human) {
	switch h.(type) { //special switch on type, asserting
	case person:
		fmt.Print("I was passed into bar ", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar ", h.(secretAgent).first)
	}
}

func main() {
	// this is now of type human as well as type secret agent
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	bar(sa1) // sa1 is secretAgent and type human
	bar(sa2) // sa2 is secretAgent and type human
	bar(p1)  //  p1 is person and human

	//fyi ... conversion
	type hotdog int
	var x hotdog = 42
	fmt.Println(x)        //42
	fmt.Printf("%T\n", x) // main.hotdog

	var y = int(x)
	fmt.Println(y)        // 42
	fmt.Printf("%T\n", y) //int
}
