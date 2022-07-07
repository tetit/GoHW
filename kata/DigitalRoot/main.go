package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {
	sum := 0
	nStr := strconv.Itoa(n)
	if len(nStr) == 1 {
		return n
	}
	for _, v := range nStr {
		val, _ := strconv.Atoi(string(v))
		sum += int(val)
	}
	return DigitalRoot(sum)

}

func main() {

	fmt.Println(DigitalRoot(167346))
	fmt.Println(DigitalRoot(7))
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(195))
	fmt.Println(DigitalRoot(992))
	fmt.Println(DigitalRoot(0))

}

//   var _ = Describe("Test Example", func() {
// 	It("fixed tests", func() {
// 	  Expect(DigitalRoot(16)).To(Equal(7))
// 	  Expect(DigitalRoot(195)).To(Equal(6))
// 	  Expect(DigitalRoot(992)).To(Equal(2))
// 	  Expect(DigitalRoot(167346)).To(Equal(9))
// 	  Expect(DigitalRoot(0)).To(Equal(0))
// 	})
//   })
