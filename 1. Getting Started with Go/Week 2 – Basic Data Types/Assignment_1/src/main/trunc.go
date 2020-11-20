package main

import (
	"fmt"
)

func main() {
	var input_number float64

	fmt.Printf("Input your number: ")
	fmt.Scan(&input_number) 				
	fmt.Println(int(input_number))
}