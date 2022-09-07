package main

import (
	"fmt"
)

func main() {
	// Konversi Tipe Data Number
	var nilai32 int32 = 32768

	// Ubah ke yang lebih besar
	var nilai64 int64 = int64(nilai32)

	// Ubah ke yang lebih besar
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi Tipe Data Byte ke String
	var name = "Ramy Abyyu"
	var r byte = name[0]
	var rString = string(r)

	fmt.Println(rString)
}