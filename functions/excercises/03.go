package main

import "fmt"

func main() {

	defer hello()
	goodbye()

}

func hello() {
	fmt.Println("Hello from hello, this should be second.")
}

func goodbye() {
	fmt.Println("Hello from goodbye, this should be first.")
}
