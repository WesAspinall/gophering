package main

// do same thing as ex1 except  with a slice

import "fmt"

func main() {

	slice := []int{1, 2, 34, 5, 6, 7, 8, 9, 10}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", slice)
}
