// time_server.go
// Just issue json formatted time for every client
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func serveEcho(conn net.Conn, id int) {
	for {
  		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))   } 
	}
}

func main() {

	ln, _ := net.Listen("tcp", "localhost:8081") // listen on all interfaces

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
