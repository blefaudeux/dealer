// bench_server.go

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func serveEcho(conn net.Conn, id int) {
	// TODO:
	// As soon as a client connects, wait for a Json request with an ID,
	// reply with the same message and the current time.

	//	for {
	//		message := map[string]string{"reqId": "0", "time": strconv.Itoa(time.Now().Second())}
	//		fmt.Println("Time is : " + message["time"] + " (client " + strconv.Itoa(id) + ")")
	//		mByte, _ := json.Marshal(message)
	//		conn.Write(mByte)
	//		time.Sleep(5 * time.Second)
	//	}
}

func main() {

	// listen on all interfaces
	ln, _ := net.Listen("tcp", "localhost:8081")

	id := 0

	for {
		fmt.Println("Server ready, awaiting connection")
		conn, _ := ln.Accept() // accept connection on port

		// Serve time and go back waiting for a new connection
		fmt.Println("Server got a new client : " + strconv.Itoa(id))
		go serveEcho(conn, id)
		fmt.Println("Benchmarking instance started")
		id = id + 1
	}
}
