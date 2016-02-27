// dealer.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

type Socket struct {
	addr, port string
	conn       net.Conn
}

func (s *Socket) printout(message string) {
	fmt.Println("Socket " + s.addr + ":" + s.port + " : " + message)
}

func (s *Socket) connect(addr string, port string) {
	s.addr = addr
	s.port = port
	conn, err := net.Dial("tcp", s.addr+":"+s.port)

	if err != nil {
		s.printout("Error connecting")
		s.printout(err.Error())
		return
	}
	s.conn = conn
	s.printout("Connection accepted")
}

func (s *Socket) close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *Socket) write(message string) {
	if s.conn == nil {
		s.connect(s.addr, s.port)
	}

	fmt.Fprintf(s.conn, message)
	s.printout("Message sent : " + message)
}

func (s *Socket) read() string {
	if s.conn == nil {
		s.connect(s.addr, s.port)
	}

	// TODO: Make this generic, read whatever comes
	message, _ := bufio.NewReader(s.conn).ReadString('}')

	s.printout("Message received : " + message)
	return message
}

func main() {
	test := Socket{}
	test.connect("localhost", "8081")

	test.write("Is there anyone here ?")
	test.read()

	fmt.Println("Program closes")
	test.close()
}
