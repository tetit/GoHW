package main

import "fmt"

func ArrayDiff(a, b []int) []int {
	rez := make([]int, 0)
	var equal bool
	for i := 0; i < len(a); i++ {
		equal = false
		for _, vb := range b {
			if vb == a[i] {
				equal = true
				continue
			}
		}
		if !equal {
		rez = append(rez, a[i])
		}
	}
	return rez
}

func main() {
	var emptyArr []int

	fmt.Println(ArrayDiff([]int{1, 2}, []int{1}))
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{1}))
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{2}))
	fmt.Println(ArrayDiff([]int{1, 2, 2}, emptyArr))
	fmt.Println(ArrayDiff(emptyArr, []int{1, 2}))
	fmt.Println(ArrayDiff([]int{1, 2, 3}, []int{1, 2}))

}

//   var _ = Describe("Sample tests", func() {
// 	It("should handle basic cases", func() {
// 	  var emptyArr []int
// 	  Expect(ArrayDiff([]int{1,2},[]int{1})).To(Equal([]int{2}))
// 	  Expect(ArrayDiff([]int{1,2,2},[]int{1})).To(Equal([]int{2,2}))
// 	  Expect(ArrayDiff([]int{1,2,2},[]int{2})).To(Equal([]int{1}))
// 	  Expect(ArrayDiff([]int{1,2,2},emptyArr)).To(Equal([]int{1,2,2}))
// 	  Expect(ArrayDiff(emptyArr,[]int{1,2})).To(BeEmpty())
// 	  Expect(ArrayDiff([]int{1,2,3},[]int{1,2})).To(Equal([]int{3}))
// 	})
// })
