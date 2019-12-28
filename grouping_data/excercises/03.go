package main

import "fmt"

//create these
//first [42 43 44 45 46]
//second [47 48 49 50 51]
//third [44 45 46 47 48]
//fourth [43 44 45 46 47]

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	first_sl := x[0:5]
	second_sl := x[5:]
	third_sl := x[2:7]
	fourth_sl := x[1:6]

	fmt.Println(first_sl)
	fmt.Println(second_sl)
	fmt.Println(third_sl)
	fmt.Println(fourth_sl)

}
