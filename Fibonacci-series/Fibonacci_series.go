package main

import "fmt"

func main() {
	a := 0
	b := 1
	var c int
	var num int
	fmt.Println("enter the number : ")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fmt.Println(a)
		c = a + b
		a = b
		b = c

	}
}
