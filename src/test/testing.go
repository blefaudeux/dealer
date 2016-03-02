package main

import (
	"dealer"
	"fmt"
)

func main() {
	test := dealer.Socket{}
	test.Connect("localhost", "8081")

	test.Send("Is there anyone here ?")
	_ = test.ReadJson()
	//	_ = test.ReadBlock("7")

	fmt.Println("Program closes")
	test.Close()
}
