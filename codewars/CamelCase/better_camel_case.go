package codewars

import (
	. "strings"
)

func camelCase(s string) string {
	return Join(Split(Title(s), " "), "")
}
