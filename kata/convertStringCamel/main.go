package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	r := make([]rune, 0)

	for i := 0; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
			v := string(s[i])
			v = strings.ToUpper(v)
			r = append(r, rune(v[0]))
			continue
		}
		r = append(r, rune(s[i]))
	}

	return string(r)
}

func main() {

	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}
