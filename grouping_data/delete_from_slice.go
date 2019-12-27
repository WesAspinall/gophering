package main

import "fmt"

func main() {
	x := []int{1, 223, 4, 5}

	x = append(x, 8)

	y := []int{2323, 234023, 3}

	// y... is known as "unfurling" or destructuring
	// can't just put y
	// otherwise you'll get -> err cannot use y (type []int) as type int in append
	x = append(x, y...)

	fmt.Println(x) // this prints [1 223 4 5 8 2323 234023 3]

	// now get rid of the 5 and 8

	x = append(x[0:2], x[5:]...)

	fmt.Println(x)

}
