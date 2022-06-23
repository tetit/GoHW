package main

import "fmt"

func SumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func main() {
	fmt.Println(SumCubes(10))
	
}