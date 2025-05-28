package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// open a tcp connection and listen on port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to bind to port: %v", err)
	}
	defer ln.Close()

	log.Printf("echo server listening on: %v", ln.Addr())
	// continuously listen for incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		} else {
			log.Printf("accepted connection: %s", conn.RemoteAddr())
		}
		// handle connection concurrently
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// create buffered reader object
	r := bufio.NewReader(conn)
	// listen for input lines using the reader object
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Printf("error from %s: %v", conn.RemoteAddr(), err)
			return
		}

		_, err = conn.Write([]byte(fmt.Sprintf("Echo: %s", line)))
		log.Printf("received from %s: %s", conn.RemoteAddr(), line)
		if err != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
