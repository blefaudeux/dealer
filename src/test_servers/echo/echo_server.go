// echo_server.go
// Just issue json formatted time for every client
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func serveEcho(conn net.Conn, id int) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('}')

		if len(message) > 0 {
			conn.Write([]byte(message))
		}
	}
}

func main() {

	ln, _ := net.Listen("tcp", "localhost:8082")
	id := 0

	for {
		fmt.Println("Server ready, awaiting connection")
		conn, _ := ln.Accept() // Blocking call, awaiting client

		// Start the echo routine, and go back waiting for a new connection
		fmt.Println("Server got a new client : " + strconv.Itoa(id))
		go serveEcho(conn, id)
		id = id + 1
	}
}
