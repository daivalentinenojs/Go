package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)

func main() {
	var m map[string] string
	m = make(map[string] string)

	fmt.Printf("Input your name: ")	
	read_name := bufio.NewReader(os.Stdin)
	rname, _, _ := read_name.ReadLine()
	name := string(rname)
	m["name"] = name

	fmt.Printf("Input your address: ")
	read_address := bufio.NewReader(os.Stdin)
	raddress, _, _ := read_address.ReadLine()
	address := string(raddress)
	m["address"] = address

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Encoding process failed")
	} else {
		fmt.Println("Encoded data : ")
		fmt.Println(b)
		fmt.Println("Decoded data : ")
		fmt.Println(string(b))
	}
	
}