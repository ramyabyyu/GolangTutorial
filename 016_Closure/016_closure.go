package main

import "fmt"

/*
	Closure merupakan kemampuan sebuah function untuk berinteraksi dengan data-data di sekitarnya dalam scope yang sama
*/

func main() {

	/*
		counter dan increment berada dalam scope yang sama, sehingga increment bisa mengakses/mengubah data dari counter
	*/
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}