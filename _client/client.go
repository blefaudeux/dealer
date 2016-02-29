package main

import "net"
import "fmt"
import "bufio"

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

//package main

//import (
//	"encoding/json"
//	"fmt"
//	"jsonpipe"
//)

//func main() {
//	jsonpipe.Handle("jsonPrint", jsonPrint)
//	jsonpipe.ListenAndServe("localhost:8081")
//}

//func jsonPrint(data *json.RawMessage) (map[string]interface{}, error) {

//	var output map[string]interface{}

//	if err := json.Unmarshal(*data, output); err != nil {
//		panic(err)
//	}
//	fmt.Println(output)
//	return output, nil
//}
