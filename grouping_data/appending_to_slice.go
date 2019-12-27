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
	fmt.Println(x)
}
