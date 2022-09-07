package main

import (
	"fmt"
)

func main() {
	/*
		-	Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
		-	Type Declaratiions biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
				dengan tujuan agar lebih mudah dimengerti
	*/

	// Contohnya jika kita ingin membuat tipe data baru noKtp yang menggantikan tipe data string
	type noKtp string

	var ktpRamy noKtp = "1111111"

	fmt.Println(ktpRamy)
	fmt.Println(noKtp("2222"))

}