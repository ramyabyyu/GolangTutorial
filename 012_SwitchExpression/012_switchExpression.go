package main

import "fmt"

func main() {
	name := "Ramy"

	fmt.Print("\n")
	fmt.Println("Switch Expression")

	switch name {
	case "Ramy":
		fmt.Println("Hello Ramy")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Your name is not Ramy or Joko")
	}

	/*
		Switch Expression dengan Short Statement

		- Sama dengan If Expression, Switch Expression juga mendukung Short Statement
			sebelum variable yang akan dicek kondisinya
	*/

	fmt.Print("\n")
	fmt.Println("Switch Expression with Short Statement")

	switch nameLength := len(name); nameLength > 5 {
	case true:
		fmt.Println("Your name is too long")
	case false:
		fmt.Println("Your name is good")
	}

	/*
		Switch tanpa kondisi

		- Kondisi di switch expression tidak wajib
		- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan di setiap case nya
	*/

	fmt.Print("\n")
	fmt.Println("Switch Expression without Codition")

	age := 25
	minAge := 18
	maxAge := 24

	switch {
	case age < minAge:
		fmt.Println("You are still underage")
	case age <= maxAge:
		fmt.Println("Your age is qualified")
	case age > maxAge:
		fmt.Println("You are too old")
	}
}