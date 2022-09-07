package main

import "fmt"

/*
	- Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
	- Sebuah interface berisikan definisi-definisi method
	- Biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Implementasi interface
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

/*
	Interface Kosong, adalah interface yang dapat menampung tipe data apapun
	dan mereturn tipe data apapun juga
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	}

	return "Ups"
}

func main() {

	ramy := Person{Name: "Ramy"}

	sayHello(ramy)

	// Interface Kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
}