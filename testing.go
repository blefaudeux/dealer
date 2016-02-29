package main

import (
	"dealer"
	"fmt"
)

func main() {
	test := dealer.Socket{}
	test.Connect("localhost", "8081")

	test.Send("Is there anyone here ?")
	jResponse := test.ReadJson()

	fmt.Println(jResponse)
	fmt.Println("Program closes")
	test.Close()
}
