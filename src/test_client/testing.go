package main

import (
	"dealer"
	"encoding/json"
	"fmt"
)

func main() {
	test := dealer.Socket{}
	test.Connect("localhost", "8081")

	mess := map[string]string{"id": "5", "content": "This is a test"}
	mess_bytes, _ := json.Marshal(mess)
	test.SendBytes(mess_bytes)

	_ = test.ReadBlock("4")

	fmt.Println("Program closes")
	test.Close()
}
