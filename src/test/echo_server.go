package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
)

func serveEcho(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('}')

		if len(message) > 0 {
			conn.Write([]byte(message))
		}
	}
}

func echo(host, port string) {
	ln, _ := net.Listen("tcp", host+":"+port)
	id := 0

    fmt.Println("Server starting")
	for {
		// Blocking call, awaiting client
		fmt.Println("Server ready, awaiting connection")
		conn, _ := ln.Accept()

		// Start the echo routine, and go back waiting for a new connection
		fmt.Println("Server got a new client : " + strconv.Itoa(id))
		go serveEcho(conn)
		id = id + 1
	}
}

func main() {
	var host = flag.String("host", "localhost", "Server host")
	var port = flag.String("port", "1234", "Server port")

    flag.Parse()
    
    fmt.Println("Server host :", *host)
    fmt.Println("Server port :", *port)
 
	echo(*host, *port)
}
