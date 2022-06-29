package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	slices := make([][]int, size)

	for i := 0; i < size; i++ {
		slice := make([]int, size)
		for j := 0; j < size; j++ {
			slice[j] = (j + 1) * (i + 1)
		}
		slices[i] = slice
	}
	return slices
}

func MultiplicationTableKata(size int) [][]int {
	res := make([][]int, size)
	for i := 0 ; i < size ; i ++ {
	  for x := 1 ; x < size + 1 ; x ++ {
		res[i] = append(res[i], (i + 1) * x)
		}}
	return res
		
  }

func main() {

	fmt.Println(MultiplicationTable(2))
}
