package main

import (
	"encoding/json"
	"fmt"
)

// marshal user to json
type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	// put Users in []

	users_slice := []User{u1, u2, u3}
	bs, err := json.Marshal(users_slice)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(string(bs))
}
