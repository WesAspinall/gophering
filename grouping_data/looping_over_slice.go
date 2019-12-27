package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 42}

	fmt.Println(len(x))
	//fmt.Prinln(cap(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	// for  index, value ... in range x
	for i, v := range x {
		fmt.Printf("the value at index %v is %v\n", i, v)
	}

}
