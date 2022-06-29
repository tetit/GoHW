package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
    	var rezult string
	sl := strings.Fields(s)
	max := rune(0)

	for _, v := range sl {
		h := countVal(v)
		if max < h {
			max = h
			rezult = v
		}
	}
	return rezult
}

func countVal(s string) rune {
	rez := rune(0)
	for _, v := range s {
		rez += v - 96
	}
	return rez
}

func main() {
	s := "man i need a taxi up to ubud"

	fmt.Println(High(s))

}
