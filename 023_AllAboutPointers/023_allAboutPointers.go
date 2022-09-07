package main

import "fmt"

/*
	Pass by Value
		- Secara default di Go-Lang semua variable itu di passing by value, bukan by refference
		- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim
			adalah duplikasi value nya
*/

type Address struct {
	City, Province, Country string
}

/*
	Pointer
		- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa
			menduplikasi data yang sudah ada
		- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

	Operator &
		- Untuk membuat sebuah variable dengan nilai pointer ke variable lain, kita bisa menggunakan operator "&"
			diikuti dengan nama variablenya

	Operator *
		- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
		- Semua variable yang mengacu ke data yang sama tidak akan berubah
		- Jika kita ingin merubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

	Function new
		- Sebelumnya untuk membuat pointer dengan menggunakan operator &
		- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
		- Namun function new hanya mengembalikan pointer ke data yang kosong, artinya tidak ada
			data awal

	Pointer di Function
		- Saat kita membuat parameter di function, secara default adalah pass by value, artinya
			data akan di copy lalu dikirim ke function tersebut
		- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah
			berubah
		- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
		- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
		- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
		- Untuk menjadikan sebuah parameter sebagai function, kita bisa menggunakan operator "*"
			di parameternya
*/

// Pointer di Function
func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}


func main() {

	// Pass by value
	address1 := Address{City: "Sukabumi", Province: "West Java", Country: "Indonesia"}
	address2 := address1

	address2.City = "Depok"

	fmt.Println(address1) // tidak berubah
	fmt.Println(address2)

	// Pass by reference
	var address3 Address = Address{City: "South Tangerang", Province: "Banten", Country: "Indonesia"}
	var address4 *Address = &address3
	// address4 := &address3

	address4.City = "Serpong"

	// Operator *
	*address4 = Address{City: "Padang", Province: "West Sumatera", Country: "Indonesia"}

	fmt.Println(address3) // Ikut berubah
	fmt.Println(address4)

	// Function New
	address5 := new(Address)
	address6 := address5

	address6.Country = "Indonesia"
	fmt.Println(address5) // address5 berubah
	fmt.Println(address6)


	// Pointer di Function
	address7 := Address{City: "Surabaya", Province: "East Java", Country: ""}
	changeAddressToIndonesia(&address7)

	fmt.Println(address7)
}