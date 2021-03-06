package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input_number string
	var arr_slice []int = make([]int, 0, 3)
	var i int

	fmt.Println("Enter an interger or X to exit: ")
	for true {
		fmt.Scanln(&input_number)
		if input_number == "X"{
            break
		}

		input_number, err := strconv.Atoi(input_number)
		if err != nil{
			fmt.Println("Your input must be a number")
			i -= 1
			continue
		} else {
			arr_slice = append(arr_slice, input_number)
		}	
	}
	sort.Ints(arr_slice[:])
	fmt.Println(arr_slice)
}