package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input_number float64

	fmt.Printf("Input your number: ")
	num, _ :=	fmt.Scan(&input_number) 	// num is the pointer
	s := fmt.Sprintf("%f", input_number)	// Convert float64 to String

	fmt.Println(s)							// Print String value
	fmt.Println(int(input_number))			// Convert float64 to Integer, Print
	fmt.Println(strconv.Itoa(num))			// Convert Integer to String
}