package main

import (
	"fmt"
	"strconv"
)

func countSheep(num int) string {
	var rez string
	if num == 0 {
		return ""
	} else {
		for i := 1; i <= num; i++ {
			rez += strconv.FormatInt(int64(i), 10) + " sheep..."
		}
	}
	return rez
}

func main() {
	fmt.Print(countSheep(2))
	fmt.Printf("")
	fmt.Println(countSheep(0))
	fmt.Printf("")
	fmt.Println(countSheep(1))

}
