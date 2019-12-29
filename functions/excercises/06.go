package main

import "fmt"

// build and use an anonymous func
func main() {
	func(x int) {
		fmt.Println(x + 1)
	}(42)

}
