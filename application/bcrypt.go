package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := []byte(`password123`)
	bs, err := bcrypt.GenerateFromPassword(s, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)         // the slice of bytes aka, what we will store in a db
	fmt.Println(string(bs)) //if you want to see the string

	loginPword := []byte(`password123`)
	err = bcrypt.CompareHashAndPassword(bs, loginPword)
	if err != nil {
		fmt.Println("You can't login")
		return
	}

	fmt.Println("You're logged in!")
}
