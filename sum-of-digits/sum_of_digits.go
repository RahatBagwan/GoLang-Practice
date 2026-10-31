package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number ")
	fmt.Scan(&num)

	sum := 0
	for num > 0 {
		rem := num % 10
		sum = sum + rem
		num = num / 10

	}

	fmt.Printf("sum of the digits = %d", sum)
}
