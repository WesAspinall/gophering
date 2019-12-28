package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// add a record
	x["Wes_A"] = []string{`Coffee`, `iPhones`, `Lamborghinis`}

	// delete a record
	delete(x, "bond_james")

	for i, v := range x {
		fmt.Printf("This is the record for %v\n", i)
		for j, v2 := range v {
			fmt.Printf("\t %v %v\n", j+1, v2)
		}
	}

}
