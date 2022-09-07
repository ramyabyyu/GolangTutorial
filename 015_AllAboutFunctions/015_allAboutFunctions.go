package main

import "fmt"

// Simple Function
func sayHello(name string) {
	fmt.Println("Hello", name)
}

// Function with Return Value
func getHello(name string) string {
	return "Hello " + name
}

// Function Returning Multiple Values
func getPersonInfo () (string, int) {
	return "Ramy Abyyu", 19
}

func getFullDates () (int, int, int) {
	return 27, 8, 2022
}

// Named Return Values
func getPersonInfo2() (fullName string, age int, isGraduated bool) {
	fullName = "Ramy Abyyu"
	age = 19
	isGraduated = true

	return
}

// Variadic Function
// Turn parameter into Slice
func sumAll (numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Function as Value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// Function as Parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFilterd := filter(name)

	fmt.Println("Hello " + nameFilterd)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

// Function Type Declaration
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

// Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Factorial With Looping
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Factorial With Recursive Function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}

	return value * factorialRecursive(value - 1)
}

func main() {
	sayHello("Ramy Abyyu")
	fmt.Println(getHello("Ramy Abyyu"))

	fullName, age := getPersonInfo()

	fmt.Println(fullName, age)

	// This function returning 3 values, but write this if you only need one
	day, _, _ := getFullDates()

	fmt.Println(day)

	// Named Return Values
	fullName, age, isGraduated := getPersonInfo2()
	fmt.Println(fullName, age, isGraduated)

	// Variadic Function
	total := sumAll(10, 23, 44, 32, 11)
	fmt.Println(total)

	// Slice as parameter in variadic function
	nums := []int {10, 20, 30, 40, 50, 60}
	total = sumAll(nums...)
	fmt.Println(total)

	// Function as Value
	goodBye := getGoodBye

	fmt.Println(goodBye("Ramy Abyyu"))

	// Function as Parameter
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter2("Guys", spamFilter)

	// Anonymous Function 1
	blacklist := func (name string) bool {
		return name == "anjing"
	}
	registerUser("Ramy", blacklist)

	// Anonymous Function 2
	registerUser("Anjing", func(s string) bool {
		return s == "Anjing"
	})

	// Factorial With Loop
	fmt.Println(factorialLoop(5))

	// Factorial With Recursive Function
	fmt.Println(factorialRecursive(5))
}

