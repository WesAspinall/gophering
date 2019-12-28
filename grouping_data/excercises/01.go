package main

import "fmt"

// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr)
}
