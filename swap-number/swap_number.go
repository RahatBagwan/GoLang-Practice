package main

import "fmt"

func main() {
	//swaping two numbers
	var a int
	var b int
	fmt.Println("enter two numbers")
	fmt.Scan(&a, &b)
	fmt.Printf("numbers before swap\n")
	fmt.Printf("a=%d  b=%d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("numbers after swap\n")
	fmt.Printf("a=%d  b=%d", a, b)

}
