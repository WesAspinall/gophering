package main

import "fmt"

func main() {
	a := factorialLoop(4)
	fmt.Println(a)
}

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

func factorialLoop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
