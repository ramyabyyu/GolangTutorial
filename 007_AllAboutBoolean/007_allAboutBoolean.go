package main

import "fmt"

func main() {
	// Komparasi / Perbandingan
	var name1 = "Ramy"	
	var name2 = "Ramy"

	var result bool = name1 == name2
	fmt.Println(result)

	// Operasi Boolean && || !
	var nilai1 = true
	var nilai2 = true

	var res1 = nilai1 && nilai2
	var res2 = nilai1 || nilai2
	var res3 = !nilai1
	var res4 = !nilai2

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
}