package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"Vanilla", "Chocolate", "Strawberry"},
	}
	p2 := person{
		first:      "Dr",
		last:       "Evil",
		favFlavors: []string{"Lava", "Bubblegum", "Rootbeer"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		// fmt.Println(v)
		fmt.Println(v.last)
		for _, v2 := range v.favFlavors {
			fmt.Printf("\t%v", v2)
		}
	}
}
