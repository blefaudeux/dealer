// Package dealer :
// Implements a TCP client with read/write calls, based on Json requests.
// Read calls can be blocking, in which case the call will return when a message with the given field
// and value is received.
package dealer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

// The Socket client handling all further requests
type Socket struct {
	addr, port string
	conn       net.Conn

	archive chan map[string]interface{}
	decoder *json.Decoder
	dooze   chan bool
}

// (private) just a quick output formatting
func (s *Socket) printout(message string) {
	fmt.Println("Socket " + s.addr + ":" + s.port + " : " + message)
}

// Connect to a TCP Socket
func (s *Socket) Connect(addr string, port string) error {
	s.addr = addr
	s.port = port
	conn, err := net.Dial("tcp", s.addr+":"+s.port)
	s.archive = make(chan map[string]interface{})

	if err != nil {
		s.printout("Error connecting")
		s.printout(err.Error())
		return err
	}
	s.conn = conn

	s.decoder = json.NewDecoder(s.conn)
	s.printout("Connection accepted")
	return nil
}

// Close : Clean up the stage and leave
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
}

// SendBytes :  a binary asap
func (s *Socket) SendBytes(message []byte) {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	s.conn.Write(message)
}

// ReadLine : until new line symbol
func (s *Socket) ReadLine() string {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	msg, err := bufio.NewReader(s.conn).ReadString('\n')

	if err != nil {
		s.printout("Error reading Socket : " + err.Error())
		return ""
	}

	return msg
}

// ReadJSON : read one Json object and return.
// Blocking until the object appears on the socket
func (s *Socket) ReadJSON() map[string]interface{} {
	if s.conn == nil {
		s.Connect(s.addr, s.port)
	}

	var msg map[string]interface{}

	if err := s.decoder.Decode(&msg); err != nil {
		s.printout("Error decoding just received json object")
	}

	return msg
}

// Populates the mailbox channel. Used as a goroutine
func (s *Socket) read(stop <-chan bool, mailbox chan<- map[string]interface{}) {
	for {
		newMessage := s.ReadJSON()

		if len(newMessage) > 0 {
			mailbox <- newMessage
		}

		select {
		// If the parent functions calls for a sleep
		case <-stop:
			return
		}
	}
}

// ReadBlock : Return the message corresponding to the given ID, when it arrives
func (s *Socket) ReadBlock(field string, value string) map[string]interface{} {

	mailbox := make(chan map[string]interface{})
	dooze := make(chan bool)

	go s.read(dooze, mailbox)

	// TODO: Go through the archive channel first ?

	for {
		testMessage := <-mailbox

		if testMessage[field] == value {
			dooze <- true
			return testMessage
		}

		s.archive <- testMessage
	}
}
