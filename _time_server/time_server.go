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

func serveTime(conn net.Conn, id int) {
	for {
		message := map[string]string{"id": strconv.Itoa(id), "time": strconv.Itoa(time.Now().Second())}
		fmt.Println("Time is : " + message["time"] + " (client " + strconv.Itoa(id) + ")")
		mByte, _ := json.Marshal(message)
		conn.Write(mByte)
		time.Sleep(5 * time.Second)
	}
}

func main() {

	// listen on all interfaces
	ln, _ := net.Listen("tcp", "localhost:8081")

	id := 1

	for {
		fmt.Println("Server ready, awaiting connection")
		conn, _ := ln.Accept() // accept connection on port

		// Serve time and go back waiting for a new connection
		fmt.Println("Server got a new client : " + strconv.Itoa(id))
		go serveTime(conn, id)
		id = id + 1
	}
}
