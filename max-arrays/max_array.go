package main

import "fmt"

func main() {
	var arr [3]int
	max := arr[0]
	fmt.Println("enter 3 elements in array : ")

	for i := 0; i < 3; i++ {
		fmt.Scan(&arr[i])
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Printf("maximum of array %d", max)

}
