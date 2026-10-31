package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var num int
	var found = false
	fmt.Println("enter the element to search in array :  ")
	fmt.Scan(&num)
	for i := 0; i < 5; i++ {
		if arr[i] == num {
			found = true
			break
		}
	}
	if found {
		fmt.Println("element exist ")
	} else {
		fmt.Println("not exist")
	}
}
