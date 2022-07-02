package main

import (
	"fmt"
	"strconv"
	// "strings"
)

// func Solve(st string, k int) int {
// 	var max int

// 	for i := 0; i < len(st); i++ {
// 		sl := strings.SplitN(st[i:], "", k+1)
// 		fmt.Println("sl 1", sl)
// 		if i < len(st)-k && k >= i{
// 			sl = append(sl, st[i:len(st)-(k-i)])
// 			fmt.Println("sl 2", sl)
// 		}
// 		for _, v := range sl {
// 			val, _ := strconv.Atoi(string(v))
// 			if max < val {
// 				max = val
// 			}
// 		}
// 	}
// 	return max

// }

func Solve(st string, k int) int {
	var max int
	sl := make([]string, 0)
	for i, _ := range st {
		if len(st)-k+i <= len(st) {
			sl = append(sl, st[i:len(st)-k+i])
		}
	}
	for _, v := range sl {
		val, _ := strconv.Atoi(string(v))
		if max < val {
			max = val
		}
	}
	return max
}

// func Solve(st string, k int) int {
// 	var max int
// 	sl := make([]string, 0, len(st))

// 	for i, _ := range st {
// 		for j := 0; j < k+1; j++ {
// 			if j <= i {
// 				sl = append(sl, st[j:i])
// 				fmt.Println("sl", sl)
// 			}

// 		}
// sl = strings.SplitN(st[i:], "", k+1)
// fmt.Println("sl", sl)
// for _, v := range sl {
// 	val, _ := strconv.Atoi(string(v))
// 	if max < val {
// 		max = val
// 	}
// }

// 	}

// 	return max

// }

// func reverse(st string) (result string) {
// 	for _, v := range st {
// 		result = string(v) + result
// 	}
// 	return
// }

// func Solve(st string, k int) int {
// 	var maxP int
// 	var max string
// 	for i := 0; i < len(st); i++ {
// 		sl := strings.SplitN(st[i:], "", k+1)
// 		for _, v := range sl {
// 			val := v
// 			if len(max) < len(val) {
// 				max = val
// 			}
// 		}
// 	}
// 	maxP, _ = strconv.Atoi(max)
// 	return maxP

// }

func main() {

	s := "2020"
	// s1 := "2020"
	// s2 := "555470000470099"
	// fmt.Println(Solve(s, 1))
	// fmt.Println(Solve(s, 2))
	// fmt.Println(Solve(s1, 1))
	// fmt.Println(Solve(s1, 3))
	fmt.Println(Solve(s, 1))

}
