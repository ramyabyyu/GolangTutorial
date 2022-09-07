package main

import "fmt"

func main() {
	/*
		Cara membuat slice

		1. array[low:high]		= Membuat slice dari array dimulai index low sampai index sebelum high
		2. array[low:]			= Membuat slice dari array dimulai index low sampai index terakhir di array
		3. array[:high]			= Membuat slice dari array dimulai index 0 sampai index sebelum index high
		4. array[:]				= Membuat slice dari array dimulai index 0 sampai index terakhir di array
	*/

	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:9]
	fmt.Println(slice1)

	// Function Slice

	// Untuk mendapatkan panjang slice
	fmt.Println(len(slice1))

	// Untuk mendapatkan kapasitas
	fmt.Println(cap(slice1))

	// Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitasnya sudah penuh, maka akan membuat array baru
	fmt.Println(append(slice1, "NewMonth"))

	// Membuat slice dari 0
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ramy"
	newSlice[1] = "Abyyu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Meng-copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Perbedaan Array vs Slice
	iniArray := [...]int {2, 3, 4, 55} // jika di dalam kurung siku ada value nya, maka itu adalah array
	iniSlice := []int {33, 44, 22, 66, 4} // jika di dalam kurung siku tidak ada value nya, maka itu adalah slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}