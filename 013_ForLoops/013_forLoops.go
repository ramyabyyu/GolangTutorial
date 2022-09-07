package main

import "fmt"

func main() {
	/*
		Perulangan di Golang hanya ada 1, yaitu For Loops
	*/

	fmt.Println("For without Statement")

	counter := 1

	for counter <= 10 {
		fmt.Print(counter)
		counter++
	}

	/*
		For dengan Statement

		- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan di for.
		- Init Statement, yaitu statement sebelum for di eksekusi
		- Post Statement, yaitu statement yang akan selalu di eksekusi di akhir setiap perulangan
	*/

	fmt.Print("\n")
	fmt.Println("For with Statement")

	for nums := 1; nums <= 10; nums++ {
		fmt.Println(nums)
	}

	slice := []string {"Ramy", "Abyyu"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/*
		For Range

		- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
		- Data collection contohnya Array, Slice, Map
	*/

	fmt.Print("\n")

	fmt.Println("For Range")

	names := []string {"Ramy", "Joko", "Budi"}

	for _, name := range names {
		// fmt.Println("index", index, "name =", name)
		fmt.Println(name)
	}
}
