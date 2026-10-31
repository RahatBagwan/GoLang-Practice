package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number ")
	fmt.Scan(&num)
	rev := 0
	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num = num / 10

	}

	fmt.Printf("reverse number =%d", rev)
}
