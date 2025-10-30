package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Println("before", list)
	handle(list[:1])
	//[1, 5, 3, 4]
	fmt.Println("after", list)
}

func handle(list []int) {
	//[1, 2}, 3, 4
	if true {
		list = append(list, 5)
		// [1, 5], 3, 4
		fmt.Println("if", list)
	}
	// list = append(list, 10)
	//[1], 5, 3, 4
	fmt.Println("append", list)
}
