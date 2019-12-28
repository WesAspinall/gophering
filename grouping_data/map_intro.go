package main

import "fmt"

// maps are key-value stores
// unordered list

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27, // need trailing comma
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // returns 0 value

	// "comma ok" idiom...
	// commonly used to check if value exists
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("The key exists", v)
	}

	// loop through map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//slice of int
	xi := []int{4, 5, 6, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
