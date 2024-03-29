package main

import "fmt"
import . "strings"

func Solve(s string) rune {
	rez := make([]rune, len(s))

	for _, v := range s {
		lI := lastIndex(v, s)
		fI := fistIndex(v, s)
		if lI != fI {
			if rez[lI-fI] == 0 {
				rez[lI-fI] = v
			} else if rez[lI-fI] > v {
				rez[lI-fI] = v
			}

		}

	}
	return rezOut(rez, s)
}

func lastIndex(v rune, s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if v == rune(s[i]) {
			return i
		}
	}
	return -1
}

func fistIndex(v rune, s string) int {
	for i := 0; i < len(s); i++ {
		if v == rune(s[i]) {
			return i
		}
	}
	return -1
}

func rezOut(sl []rune, s string) rune {
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] != 0 {
			return sl[i]
		}
	}
	return rezNorep(s)
}

func rezNorep(s string) rune {
	min := rune(s[0])
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func SolveKata(s string) rune {
    max,r := 0,rune(123)
    for _,c := range s {
        d := LastIndex(s, string(c)) - Index(s, string(c))
        if max < d || max == d && r > c {
            max = d; r = c
        }
    }
    return r
}

func main() {

	s := "aabccc"

	fmt.Println(Solve(s))
	fmt.Println(SolveKata(s))
	fmt.Println("a - z", 'a', 'z')
	fmt.Println(string(rune(123)))

}
