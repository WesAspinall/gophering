package main

import "fmt"

func main() {
	xi := []int{4, 3, 34, 5, 6, 78, 9, 1}
	s := sum(xi...)
	fmt.Println(s)

	s2 := even(sum, xi...)
	fmt.Println(s2)

	s3 := odd(sum, xi...)
	fmt.Println(s3)
}

func sum(x ...int) int {

	sum := 0
	for _, v := range x {
		sum = sum + v
	}

	return sum

}

func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

// create a function odd that adds up the odd numbers
func odd(f func(x ...int) int, vi ...int) int {
	var odd_cache []int

	for _, v := range vi {
		if v%2 != 0 {
			odd_cache = append(odd_cache, v)
		}
	}

	return f(odd_cache...)
}
