package main

import (
	"fmt"
	"strings"
)

func Anagrams(word string, words []string) []string {
	var s []string
	for _, v := range words {
		count := 0
		if len(word) == len(v) {
			for _, val := range word {
				bo := strings.Count(word, string(val)) == strings.Count(v, string(val))
				if strings.Contains(v, string(val)) && bo {
					count++
				}
			}
		}
		if count == len(v) {
			s = append(s, v)
		}
	}
	return s
}

func main() {
	s := []string{"lazing", "lazy", "lacer"}
	fmt.Println(Anagrams("laser", s))
}
