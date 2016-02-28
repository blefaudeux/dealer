package main

import (
	"dealer"
	"fmt"
)

func main() {
	test := dealer.Socket{}
	test.connect("localhost", "8081")

	test.write("Is there anyone here ?")
	test.readString()

	fmt.Println("Program closes")
	test.close()
}
