// dealer.go
package dealer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Socket struct {
	addr, port string
	conn       net.Conn
}

// (private) just a quick output formatting
func (s *Socket) printout(message string) {
	fmt.Println("Socket " + s.addr + ":" + s.port + " : " + message)
}

func (s *Socket) Connect(addr string, port string) {
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

func (s *Socket) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *Socket) Send(message string) {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	fmt.Fprintf(s.conn, message)
	s.printout("Message sent : " + message)
}

// Read until new line
func (s *Socket) ReadLine() string {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	message, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		s.printout("Error reading socket : " + err.Error())
		return ""
	}

	s.printout("Message received : " + message)
	return message
}

func (s *Socket) ReadJson() interface{} {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	// we create a decoder that reads directly from the socket
	d := json.NewDecoder(s.conn)

	var msg interface{}

	err := d.Decode(&msg)
	fmt.Println(msg, err)

	pretty_print, _ := json.MarshalIndent(msg, "", "\t")
	s.printout("Message received : " + string(pretty_print))
	return msg
}
