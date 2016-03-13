// See https://blog.filippo.io/building-python-modules-with-go-1-5/#thecompletedemosource
// and https://gopy.qur.me/extensions/examples.html


package main

import (
	"bufio"
	"dealer"
	"encoding/json"
	"fmt"
  "gopy"
	"net"
	"strconv"
	"time"
)

struct PySocket {
  socket  dealer.Socket
  port    string
  address string
}

func Connect(s *PySocket, args *py.Tuple)(err Error) {
  // Parse python arguments
  var o py.Object
  if err := py.ParseTuple(args, "O", &o); err != nil {
    fmt.Println("Error parsing arguments")
  }

  // Connect dealer
  s.socket = dealer.Socket{}
  if err = s.socket.Connect( o[0], o[1]); err != nil {
    fmt.Println("Error creating the socket")
  }
}

func Read(s * PySocket, list *py.List) (error Error){
  // TODO: Ben
  // We need to grab the ID here, and do a ReadJSON on the dealer socket
}

func Send(s * PySocket, list *py.List) (error Error){
  // TODO: Ben
}

func Close(s *PySocket) {
  s.socket.Close()
}
