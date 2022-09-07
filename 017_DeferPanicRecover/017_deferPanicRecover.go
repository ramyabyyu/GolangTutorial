package main

import "fmt"

/*
	1. Defer
	 - Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function
		selesai di eksekusi
	 - Defer function akan selalu di eksekusi walaupun terjadi error di function yang dieksekusi
*/
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // akan mengeksekusi function logging setelah function runApplication tereksekusi
	fmt.Println("Run Application")
}

/*
	2. Panic
	 - Panic function adalah function yang bisa kita gunakan untuk menghentikan program
	 - Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
	 - Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

// func endApp() {
// 	fmt.Println("End App")
// }

// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Error")
// 	}

// 	fmt.Println("App Running")
// }

/*
	3. Recover
	 - Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	 - Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/
func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error occurred", message)
	}

	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Error!")
	}
	fmt.Println("App Running")
}

func main() {
	// defer
	// runApplication()

	// panic
	// runApp(true)

	// recover
	runApp(true)
}