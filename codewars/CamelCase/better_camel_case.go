package codewars

//https://www.codewars.com/kata/camelcase-method/go

import (
	. "strings"
)

func camelCase(s string) string {
	return Join(Split(Title(s), " "), "")
}
