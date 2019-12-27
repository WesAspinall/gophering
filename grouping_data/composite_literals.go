package main

import "fmt"

func main() {
	// x := type{values} // composite literal syntax // values of the same type
	x := []int{1, 2, 3}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[2])
}
