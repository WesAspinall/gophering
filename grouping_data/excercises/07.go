package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{a, b}

	// fmt.Println(x)

	// range over records, then print the datat in each record
	for i, v := range x {
		fmt.Println(i, v)
		for _, thing := range v {
			fmt.Println(thing)
		}
	}
}
