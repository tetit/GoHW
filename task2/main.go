package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	base := 10
	valInt := make([]int32, 0)

	inputSplit := strings.Split(input, " ")

	for _, v := range inputSplit {
		if convert, err := strconv.Atoi(v); err == nil {
			valInt = append(valInt, int32(convert))
		}
	}

	maxVal := float64(valInt[0])
	minVal := float64(valInt[0])

	for _, v := range valInt {
		maxVal = math.Max(maxVal, float64(v))
		minVal = math.Min(minVal, float64(v))
	}

	if maxVal == minVal {
		result = strconv.FormatInt(int64(maxVal), base)
	} else {
		result = strconv.FormatInt(int64(maxVal), base) + " " + strconv.FormatInt(int64(minVal), base)
	}

	fmt.Println(result)
}
