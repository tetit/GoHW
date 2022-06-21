package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	mapResult := make(map[int]bool)

	for _, key := range arr {
		if _, value := mapResult[key]; !value {
			mapResult[key] = true
			result = append(result, key)
		}
	}

	// other solution without map

	// var condition bool
	// var s []int

	// for !condition {
	// 	result = append(result, arr[:1]...)
	// 	s = nil
	// 	for i, j := 0, 1; j < len(arr); j++ {
	// 		if arr[i] != arr [j] {
	// 			s = append(s, arr[j])
	// 		}
	// 	}
	// 	arr = s
	// 	if len(arr) == 0 {
	// 		condition = true
	// 	}
	// }

	fmt.Println(result)

}
