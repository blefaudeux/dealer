package dealer

import (
	"bufio"
	"dealer"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

func serveEcho(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('}')

		if len(message) > 0 {
			conn.Write([]byte(message))
		}
	}
}

func echo(host, port string) {
	ln, _ := net.Listen("tcp", host+":"+port)
	id := 0

	for {
		// Blocking call, awaiting client
		fmt.Println("Server ready, awaiting connection")
		conn, _ := ln.Accept()

		// Start the echo routine, and go back waiting for a new connection
		fmt.Println("Server got a new client : " + strconv.Itoa(id))
		go serveEcho(conn)
		id = id + 1
	}
}

// BenchmarkRPS : Test the number of requests per second handled by our client
func BenchmarkRPS(b *testing.B) {
	host := "localhost"
	port := "8082"

	// Start our echo test server
	go echo(host, port)
	time.Sleep(time.Second)

	// Connect our client
	test := dealer.Socket{}
	test.Connect("localhost", "8082")

	// Warmup
	mess := map[string]string{"id": "5", "content": "This is a test"}
	messBytes, _ := json.Marshal(mess)

	test.SendBytes(messBytes)
	_ = test.ReadBlock("id", "5")

	// Benchmark
	fmt.Printf("\n----------------\n")
	fmt.Println("Starting benchmark")
	numberOfRuns := b.N
	numberOfBenchs := 10

	var runtime = make([]float64, numberOfBenchs)

	for j := 0; j < numberOfBenchs; j++ {
		start := time.Now()
		for i := 0; i < numberOfRuns; i++ {
			test.SendBytes(messBytes)
			_ = test.ReadBlock("id", "5")
		}

		elapsed := time.Since(start)
		runtime[j] = float64(numberOfRuns) / elapsed.Seconds()
		fmt.Printf("Processed %.0f requests per second\n", runtime[j])
	}

	fmt.Printf("\n----------------\n")
	fmt.Print("Runs completed.\n")
	avg := 0.
	for _, val := range runtime {
		avg += val
	}

	avg /= float64(numberOfBenchs)
	fmt.Printf("Processed on average %.0f requests per second\n", avg)

	delay := 1000. / avg
	fmt.Printf("corresponding to %.3f ms delay\n", delay)

	test.Close()
}
