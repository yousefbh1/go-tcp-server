package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		// error handling
	}
	fmt.Fprintf(conn, "GET / HTTP")
}
