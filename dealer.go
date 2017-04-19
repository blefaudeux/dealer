// Package dealer :
// Implements a TCP client with read/write calls, based on Json requests.
// Read calls can be blocking, in which case the call will return when a message with the given field
// and value is received.
package dealer

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

// The Socket client handling all further requests
type Socket struct {
	addr string
	port int
	conn net.Conn

	// archive chan map[string]interface{}
	decoder *json.Decoder
	dooze   chan bool
}

func (s *Socket) getConnectionParameters() string {
	return s.addr + ":" + strconv.Itoa(s.port)
}

// (private) just a quick output formatting
func (s *Socket) printout(message string) {
	fmt.Println("Socket " + s.getConnectionParameters() + " : " + message)
}

// (private) reconnect given known settings
func (s *Socket) autoConnect() error {
	return s.Connect(s.addr, s.port)
}

// Connect to a TCP Socket
func (s *Socket) Connect(addr string, port int) error {
	s.addr = addr
	s.port = port
	conn, err := net.Dial("tcp", s.getConnectionParameters())

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

// SendBytes : send a binary asap
func (s *Socket) SendBytes(message []byte) error {
	if s.conn == nil {
		if err := s.autoConnect(); err != nil {
			return err
		}
	}

	s.conn.Write(message)
	return nil
}

// ReadLine : until new line symbol
func (s *Socket) ReadLine() (string, error) {
	if s.conn == nil {
		if err := s.autoConnect(); err != nil {
			return "", err
		}
	}

	msg, err := bufio.NewReader(s.conn).ReadString('\n')

	if err != nil {
		s.printout("Error reading Socket : " + err.Error())
		return "", err
	}

	return msg, nil
}

// ReadJSON : read one Json object and return.
// Blocking until the object appears on the socket
func (s *Socket) ReadJSON() (map[string]interface{}, error) {
	if s.conn == nil {
		if err := s.autoConnect(); err != nil {
			return nil, err
		}
	}

	var msg map[string]interface{}

	if err := s.decoder.Decode(&msg); err != nil {
		return nil, err
	}

	return msg, nil
}

// Populates the mailbox channel. Used as a goroutine
func (s *Socket) read(stop <-chan bool, mailbox chan<- map[string]interface{}) {
	for {
		if newMessage, _ := s.ReadJSON(); len(newMessage) > 0 {
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
func (s *Socket) ReadBlock(field string, value string, timeout int) (map[string]interface{}, error) {

	mailbox := make(chan map[string]interface{})
	dooze := make(chan bool)

	go s.read(dooze, mailbox)

	for {
		select {
		case testMessage := <-mailbox:
			if testMessage[field] == value {
				dooze <- true
				return testMessage, nil
			}
		case <-time.After(time.Second * time.Duration(timeout)):
			dooze <- true // Not needed strictly speaking, stop the goroutine
			return make(map[string]interface{}), errors.New("Timeout")
		}
	}
}
