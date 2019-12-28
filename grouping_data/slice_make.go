package main

// when working with slices using a composite literal, if you want to add
// an item to an array, it go creates a new array and throws away the old one
// bc the length changes
// with make(), the length is set
// (type, len, cap)
import "fmt"

func main() {

	x := make([]int, 10, 100)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// x[10] = 2232 //will return index out of range
	x = append(x, 2232)

	fmt.Println(x)
	fmt.Println(len(x)) // length is now 11
	fmt.Println(cap(x))
}
