package main

import (
	taskTests "GoHW/taskTests/funcDo"
	"fmt"
)

func main() {
	fmt.Println(taskTests.Do("d", 3, true))
	fmt.Println(taskTests.Do("d", 21, true))
	fmt.Println(taskTests.Do("e", 3, true))
}
