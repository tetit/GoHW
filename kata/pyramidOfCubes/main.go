package main

import (
	"fmt"
)

// func FindHeight(cubes int) int {
// 	t := 1
// 	sec := 3
// 	if cubes != 0 {

// 		for i := 1; i <= cubes; {
// 			t += sec * i
// 			if t < cubes {
// 				i++
// 			} else if t == cubes {
// 				return i + 1
// 			} else {
// 				return i
// 			}
// 		}
// 	}
// 	return 0
// }
func FindHeight(cubes int) int {
	t := 0
	sec := 0
	if cubes != 0 {

		for i := 1; i <= cubes; {
			t += i
			sec += t
			if sec < cubes {
				i++
			} else if sec == cubes {
				return i
			} else {
				return i - 1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(FindHeight(0))
	fmt.Println(FindHeight(1))
	fmt.Println(FindHeight(3))
	fmt.Println(FindHeight(4))
	fmt.Println(FindHeight(10))
	fmt.Println(FindHeight(19))
	fmt.Println(FindHeight(60))
	fmt.Println(FindHeight(168000000))
}
