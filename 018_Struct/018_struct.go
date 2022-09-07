package main

import "fmt"

/*
	- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
		lainnya dalam satu kesatuan
	- Struct biasanya representasi data dalam program aplikasi yang kita buat
	- Data di struct disimpan dalam field
	- Sederhananya struct adalah kumpulan dari field
*/

type Student struct {
	Name, Address	string
	Age				int
	IsGraduated		bool
}

// Struct Method
func (student Student) sayHello() {
	fmt.Println("Hello my name is", student.Name)
}

func main() {
	var ramy Student
	ramy.Name = "Ramy Abyyu"
	ramy.Address = "Jl Abc No 123"
	ramy.Age = 19
	ramy.IsGraduated = true

	fmt.Println(ramy)

	// Struct Literal
	joko := Student{
		Name: "Joko",
		Address: "Jl Def No 456",
		Age: 20,
		IsGraduated: true,
	}

	fmt.Println(joko)

	budi := Student{"Budi", "Jl Ghi No 789", 16, false}
	fmt.Println(budi)

	tiara := Student{Name: "Tiara"}
	tiara.sayHello()
}