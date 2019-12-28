package main

import "fmt"

func main() {
	a := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(a)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("for item in index position", i, "we are now adding,", v, "to the total which is now ", sum)
	}
	fmt.Println("The total is ", sum)
	return sum
}
