package main

import "fmt"

func main() {
	fmt.Println("h")
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	a := func(x int) {
		fmt.Println("The year big brother started watching us:", x)
	}
	a(1984)
}
