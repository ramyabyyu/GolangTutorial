package main

import "fmt"

func main() {

	name := "Rahmad"

	if name == "Ramy" {
		fmt.Println("Hello Ramy")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Your name is not Ramy or Joko")
	}


	/*
		If Short Statement

		- If mendukung short statement sebelum kondisi.
		- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan
			pengecekan terhadap kondisi

		- if <shortStatement>; <condition> {
				<action>
			}
	*/

	if nameLength := len(name); nameLength > 5 {
		fmt.Println("Your name is too long")
	} else {
		fmt.Println("Your name is good")
	}
}