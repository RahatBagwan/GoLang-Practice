package main

import "fmt"

func main() {
	// lagerst of 3 number

	a := 22
	b := 3
	c := 4

	if a > b && a > c {
		fmt.Printf("%d is larger ", a)
	} else if b > a && b > c {
		fmt.Printf("%d is larger", b)
	} else {
		fmt.Printf("%d is larger ", c)
	}
}
