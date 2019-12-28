package main

import "fmt"

type person struct {
	first_name  string
	last_name   string
	favorite_ic []string
}

func main() {
	p1 := person{
		first_name:  "Bill",
		last_name:   "Williams",
		favorite_ic: []string{"vanilla", "bubblegum"},
	}
	p2 := person{
		first_name: "James",
		last_name:  "Bond",
		favorite_ic: []string{
			"chocolate",
			"rum and coke",
			"martini"},
	}

	fmt.Println(p1.first_name, p1.last_name, p1.favorite_ic)
	fmt.Println(p2.first_name, p2.last_name)
	for i, v := range p2.favorite_ic {
		fmt.Println(i, v)
	}

}
