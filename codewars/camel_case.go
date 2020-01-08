package codewars

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	t := strings.TrimSpace(s)
	sl := strings.Split(t, " ")

	if s == "" {
		return ""
	} else {
		for i := 0; i < len(sl); i++ {
			firstLetter := strings.ToUpper(string(sl[i][0]))
			restOfWord := strings.ToLower(string(sl[i][1:]))

			a := fmt.Sprintf("%s%s", firstLetter, restOfWord)
			sl[i] = a
		}

	}

	camelCase := fmt.Sprint(strings.Join(sl, ""))

	return camelCase

}
