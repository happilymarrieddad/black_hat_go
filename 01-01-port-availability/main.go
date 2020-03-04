package main

import (
	"fmt"
	"net"
)

func main() {
	if _, err := net.Dial("tcp", "scanme.nmap.org:80"); err != nil {
		panic(err)
	}

	fmt.Println("Connection successful!")
}
