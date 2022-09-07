package main

import "fmt"

func main() {
	/*
		- Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
		- Break digunakan untuk menghentikan seluruh perulangan
		- Continue digunakan untuk menghentikan perulangan yang sedang berjalan, dan melanjutkan
			ke perulangan selanjutnya
	*/

	fmt.Println("Break")

	// Break
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)

		if (i == 5) {
			break
		}
	}

	fmt.Print("\n")
	fmt.Println("Continue")

	// Continue
	for i := 0; i < 10; i++ {
		if i % 2 == 1 {continue}
		fmt.Println("Perulangan ke", i)
	}
}