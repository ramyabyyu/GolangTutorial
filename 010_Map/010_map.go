package main

import (
	"fmt"
)

func main() {
	/*
		- Pada Array dan Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
		- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe
			data index yang akan kita gunakan
		- Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat
			unik, tidak boleh sama
		- Berbeda dengan Array dengan Slice, jumlah data yang kita masukkan ke dalam map ke dalam Map boleh sebanyak-banyaknya,
			asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya
			akan diganti dengan data baru
	*/

	person := map[string]string{
		"name":    "Ramy Abyyu",
		"address": "Tangerang Selatan",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Tambah data person
	person["title"] = "Programmer"

	fmt.Println(person)

	/*
		Function Map

		1. len(map)						=> Untuk mendapatkan jumlah data di map
		2, map[key]						=> Untuk mendapatkan data di map dengan key
		3. map[key] = value				=> Mengubah data di map dengan key
		4. make(map[TypeKey]TypeValue)	=> Membuat map baru
		5. delete(map, key)				=> Menghapus data di map dengan key
	*/

	book := make(map[string]string)

	book["title"] = "Go-Lang"
	book["author"] = "Ramy"
	book["ups"] = "Salah"

	fmt.Println(book)
	
	delete(book, "ups")
	
	fmt.Println(book)
}