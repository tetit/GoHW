package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	sL := strings.ToLower(str)
	for i, v := range sL {
		count := strings.Count(sL, string(v))
		if count == 1 {
			return string(str[i])
		}
	}
	return ""
}

func main() {
	s := "sTAretSS"
	fmt.Println(FirstNonRepeating(s))
}
