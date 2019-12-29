package main

// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument
import "fmt"

func main() {
	fmt.Println(foo(bar, "suh dude"))
}

func bar() string {
	return "hello from bar"
}

func foo(b func() string, s string) string {

	x := b()

	return fmt.Sprint("this is from bar: ", x, ", and this is the string you passed in: ", s)
}
