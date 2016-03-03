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
	letterbox  chan map[string]interface{}
	decoder    *json.Decoder
	dooze      chan bool
}

// (private) just a quick output formatting
func (s *Socket) printout(message string) {
	fmt.Println("Socket " + s.addr + ":" + s.port + " : " + message)
}

// Connect to a TCP socket
func (s *Socket) Connect(addr string, port string) error {
	s.addr = addr
	s.port = port
	conn, err := net.Dial("tcp", s.addr+":"+s.port)

	if err != nil {
		s.printout("Error connecting")
		s.printout(err.Error())
		return err
	}

	s.letterbox = make(chan map[string]interface{})
	s.dooze = make(chan bool)

	s.conn = conn
	s.decoder = json.NewDecoder(s.conn)
	s.printout("Connection accepted")
	return nil
}

func (s *Socket) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

// Send a string asap
func (s *Socket) Send(message string) {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	fmt.Fprintf(s.conn, message)
	s.printout("Message sent : " + message)
}

// Send a binary asap
func (s *Socket) SendBytes(message []byte) {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	s.conn.Write(message)
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

	return message
}

// Read one Json object only
func (s *Socket) ReadJson() map[string]interface{} {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	var msg map[string]interface{}

	if err := s.decoder.Decode(&msg); err != nil {
		s.printout("Error decoding just received json object")
	}

	return msg
}

// Populates the letterbox channel. Used as a goroutine
func (s *Socket) read() {
	newMessage := s.ReadJson()

	if len(newMessage) > 0 {
		s.letterbox <- newMessage
	}

	select {
	case <-s.dooze:
		return
	}
}

// Return the message corresponding to the given ID, when it arrives
func (s *Socket) ReadBlock(reqId string) map[string]interface{} {
	go s.read()

	for {
		testMessage := <-s.letterbox

		if testMessage["id"] == reqId {
			s.dooze <- true
			return testMessage
		}
	}
}
