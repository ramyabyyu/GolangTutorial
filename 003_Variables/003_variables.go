package main

/*
	Integer, tipe data angka biasa dengan negatif :
		int8, int16, int32, int64

	Unsigned Integer, tipe data integer tanpa negatif (sign) :
		uint8, uint16, uint32, uint64

	Floating Point, tipe data angka dengan koma :
		float32, float64, complex64, complex128

	Alias, nama lain dari sebuah tipe data :
		byte	alias untuk		uint8
		rune	alias untuk		int32
		int		alias untuk		Minimal int32
		uint	alias untuk		Minimal uint32

*/

import "fmt"

func main() {
	// Number
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)

	// Boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// String
	fmt.Println("Ramy Abyyu")
	fmt.Println(len("Rmay Abyyu"))
	fmt.Println("Ramy Abyyu"[0])
}