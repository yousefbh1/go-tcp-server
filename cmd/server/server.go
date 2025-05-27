package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to bind to port: %v", err)
	}
	defer ln.Close()

	log.Printf("echo server listening on: %v", ln.Addr())
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Printf("error from %s: %v", conn.RemoteAddr(), err)
			return
		}

		_, err = conn.Write([]byte(fmt.Sprintf("Echo: %s", line)))
		if err != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
