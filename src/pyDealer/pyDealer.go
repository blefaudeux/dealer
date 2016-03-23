// See https://blog.filippo.io/building-python-modules-with-go-1-5/#thecompletedemosource
// and https://gopy.qur.me/extensions/examples.html

package main

import (
	"dealer"
	"fmt"
)

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
import "C"

//export GoSocket
type GoSocket struct {
	socket  dealer.Socket
	port    string
	address string
}

//export Connect
func Connect(s *GoSocket, address *C.PyObject, port *C.PyObject) (err Error) {

	s.address = C.GoString(address)
	s.port = C.GoString(port)

	fmt.Println("Connecting to {} {}".format(s.address, s.port))

	// Connect dealer
	s.socket = dealer.Socket{}
	if err = s.socket.Connect(address, port); err != nil {
		fmt.Println("Error creating the socket")
	}
}

//export Read
func Read(s *GoSocket, list *py.List) (error Error) {
	// TODO: Ben
	// We need to grab the ID here, and do a ReadJSON on the dealer socket
}

//export Send
func Send(s *GoSocket, list *py.List) (error Error) {
	// TODO: Ben
}

//export Close
func Close(s *GoSocket) {
	s.socket.Close()
}
