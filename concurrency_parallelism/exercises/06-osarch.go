package main

// go run
// go build
// go install - installs in go dir bin/
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
