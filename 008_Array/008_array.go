package main

import "fmt"

func main() {

	// Array 1
	var names [2]string

	names[0] = "Ramy"
	names[1] = "Abyyu"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	// Array 2
	var numbers = [3]int {
		90,
		95,
		100,
	}

	fmt.Println(numbers)

	// Array 3 (Jika kita tidak set kapasitas array nya, maka gunakan ...)
	var cars = [...]string {
		"Mercedes Benz",
		"BMW",
		"Volkswagen",
		"Audi",
	}

	fmt.Println(cars)

	// Operasi Array

	//len(array) = untuk mengetahui panjang Array
	fmt.Println(len(names))
	fmt.Println(len(numbers))

}