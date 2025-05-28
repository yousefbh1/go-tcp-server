package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// open a tcp connection and listen on port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to bind to port: %v", err)
	}
	defer ln.Close()

	log.Printf("echo server listening on: %v", ln.Addr())
	// track active connections
	var wg sync.WaitGroup

	// signal handler
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("shutdown signal received - closing listener")
		ln.Close()
	}()

	// continuously listen for incoming connections
acceptLoop:
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			break acceptLoop
		}
		log.Printf("accepted connection: %s", conn.RemoteAddr())

		// handle connection concurrently
		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
			handleConn(c)
		}(conn)
	}

	log.Println("waiting for active connections to finish")
	wg.Wait()
	log.Println("server exiting")
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

		// write back to the connected client
		_, err = conn.Write([]byte(fmt.Sprintf("Echo: %s", line)))
		// log in the server
		log.Printf("received from %s: %s", conn.RemoteAddr(), line)
		if err != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
