package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

type Name struct {
	first_name string
	last_name string
}

func cutLongName(buffer string) string {
    if len(buffer)>20 {
       return string(buffer[0:20])
    } else {
       return buffer 
    }
}

func main() {
	path := ""
	fmt.Printf("Input your file full path: ")
	fmt.Scanln(&path)

	var name[] Name
	file_data, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file_data.Close()

	read_data := bufio.NewReader(file_data)
	for {
		line, _, err := read_data.ReadLine()

		if err != nil || io.EOF == err {
			break
		}

		temp_name := strings.Split(string(line), " ")
		struct_name := Name {
			cutLongName(temp_name[0]),
			cutLongName(temp_name[1]),
		}
		name = append(name, struct_name)
	}

	for i:= 0; i < len(name); i++ {
		fmt.Println(i+1)
		fmt.Println("First name: "+ name[i].first_name)
		fmt.Println("Last name: "+ name[i].last_name)
	}

}