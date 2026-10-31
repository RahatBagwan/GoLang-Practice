package main

import "fmt"

func main() {
	var arr [3]int
	sum := 0
	fmt.Println("enter 3 elements in array : ")

	for i := 0; i < 3; i++ {
		fmt.Scan(&arr[i])
		sum = sum + arr[i]
	}

	fmt.Printf("sum of array elemnts = %d ", sum)
}
