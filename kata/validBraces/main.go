package main

import (
	"fmt"
	// "strings"
)

func ValidBraces(str string) bool {
	countIn := 0
	countOu := 0
	sl := make([]rune, 0)
	for i, v := range str {
		switch v {
		case '(':
			sl = append(sl, v)
			countIn++
		case '[':
			sl = append(sl, v)
			countIn++
		case '{':
			sl = append(sl, v)
			countIn++
		case ')':
			if i == 0 || len(sl) == 0 {
				return false
			} else if v == sl[len(sl)-1]+1 {
				sl = sl[:len(sl)-1]
				countOu++
			} else {
				return false
			}
		case ']':
			if i == 0 || len(sl) == 0 {
				return false
			} else if v == sl[len(sl)-1]+2 {
				sl = sl[:len(sl)-1]
				countOu++
			} else {
				return false
			}
		case '}':
			if i == 0 || len(sl) == 0 {
				return false
			} else if v == sl[len(sl)-1]+2 {
				sl = sl[:len(sl)-1]
				countOu++
			} else {
				return false
			}
		default:
			return false

		}
	}
	if countIn != countOu {
		return false
	} else {
		return true
	}
}

func main() {
	s := "(("
	s1 := "[}{]"
	s2 := "()({{}})"
	s3 := "()))"
	s4 := "{]"
	s5 := "({})[({})]"

	fmt.Println(ValidBraces(s))
	fmt.Println(ValidBraces(s1))
	fmt.Println(ValidBraces(s2))
	fmt.Println(ValidBraces(s3))
	fmt.Println(ValidBraces(s4))
	fmt.Println(ValidBraces(s5))
}
