package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup
	for i, v := range n {
		wg.Add(1)
		go func(v []int, i int) {
			defer wg.Done()
			fmt.Printf("slice %v: ", i+1)
			sum(v)

		}(v, i)
	}

	wg.Wait()

}

func sum(sum []int) {
	var sumSl int
	for _, v := range sum {
		sumSl += v
	}
	fmt.Printf("%d\n", sumSl)
}
