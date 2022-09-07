package main

import "fmt"

func main() {
	// Pertambahan
	var a = 10
	var b = 15
	var c = a + b
	fmt.Println(c)

	// Augmented Assignment
	var i = 10
	i += 1
	fmt.Println(i)

	// Unary Operator
	a++ // a = a + 1
	b++
	fmt.Println(a)
}