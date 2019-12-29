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

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: false,
	}

	fmt.Println(sa1.speak())
}

func (s secretAgent) speak() string {
	var verb string
	if s.ltk {
		verb = "do"
	} else {
		verb = "do not"
	}
	return fmt.Sprint("Hello, I am ", s.first, " ", s.last, " and I ", verb, " have a license to kill.")
}
