package main

import (
	"errors"
	"fmt"
)

func dividedBy(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Pembagian Dengan Nol")
	}

	return value / divider, nil
}

func main() {
	result, err := dividedBy(10, 5)

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println("Result :", result)
	}
}