package main

import "fmt"

func main() {
	defer hello()
	goodbye()
}

func hello() {
	fmt.Println("hello from hello")
}

func goodbye() {
	fmt.Println("hello from goodbye")
}
