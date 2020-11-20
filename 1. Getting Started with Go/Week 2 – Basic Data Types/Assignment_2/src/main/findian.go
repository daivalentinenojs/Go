package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input_string string

	fmt.Printf("Input your string: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input_string = strings.Replace(text, " ", "", -1)
	input_string = strings.ToLower(input_string)
	len_input_string := len(input_string)
	
	if (input_string[0] == 105) {
		if (input_string[len_input_string-3] == 110) {
			if strings.Contains(input_string, "a") {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}	
		} else {
			fmt.Println("Not Found!")
		}				
	} else {
		fmt.Println("Not Found!")
	}		
}