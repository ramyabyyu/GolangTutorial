package main

import (
	"fmt"
)

/*
	- Type Assertion merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
	- Fitur ini sering digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "OK"
}

/*
	Type Assertions menggunakan Switch Expression

	- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	- Jika panic tidak ter-recover, maka otomatis program kita akan mati
	- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/

func random2() interface{} {
	return true
}

func main() {
	result := random()
	resultStr := result.(string)
	fmt.Println(resultStr)

	// Type Assertion with Switch Statement
	result2 := random2()
	switch value := result2.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Integer")
	default:
		fmt.Println("Unknown Type")
	}
}