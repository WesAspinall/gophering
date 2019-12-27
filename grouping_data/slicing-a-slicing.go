package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 5, 6, 42}

	// from 0, up to but not including 3
	//[0:3]
	fmt.Println(x[0:3])
	fmt.Println(x[3:len(x)])

	// reviewing...
	//looping without using range
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v\n", x[i])
	}

	// looping with range
	for i, v := range x {
		fmt.Println(i, v)
	}
}
