package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)  // type string
	fmt.Printf("%T\n", bs) // type uint8

	// we want to unmarshal this bs into a struct
	// unmarshal takes a slice of bytes and a pointer to a value
	// people := []person{} // composite literal, same as below
	var people []person // "declare var people with a type of slice of people" // the data will be unmarshaled into here

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
