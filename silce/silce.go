package main

import (
	"fmt"
)

func main() {
	//var num []int
	var num = make([]int, 1, 5)

	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6)
	num = append(num, 7)
	var num2 = make([]int, len(num))
	fmt.Println(num)
	fmt.Println("langht of silce", len(num))
	fmt.Println("capacity of silce", cap(num))

	// copy  silce in other
	copy(num2, num)
	fmt.Println("after copy silce ")
	fmt.Printf("num=%d \n num2=%d", num, num2)

	// marge two silces
	marge := append(num, num2...)
	fmt.Println("\nmarge silce \n", marge)

}
