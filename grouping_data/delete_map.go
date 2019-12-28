package main

import "fmt"

func main() {
	//create a map of strings and ints
	x := map[string]int{
		"James Bond":      32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(x)

	delete(x, "James Bond")

	fmt.Println(x)

	//again, comma ok idiom...
	// if moneypenny exists, print v, then delete
	if v, ok := x["Miss Moneypenny"]; ok {
		fmt.Println("value", v)
		delete(x, "Miss Moneypenny")
	}
}
