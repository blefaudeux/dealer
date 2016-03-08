package main

import (
	"dealer"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	test := dealer.Socket{}
	test.Connect("localhost", "8082")

	mess := map[string]string{"id": "5", "content": "This is a test"}
	messBytes, _ := json.Marshal(mess)

	// Warmup
	test.SendBytes(messBytes)
	_ = test.ReadBlock("5")

	// Benchmark
	fmt.Println("Starting benchmark")
	numberOfRuns := 500
	numberOfBenchs := 10

	var runtime = make([]float64, numberOfBenchs)

	for j := 0; j < numberOfBenchs; j++ {
		start := time.Now()
		for i := 0; i < numberOfRuns; i++ {
			test.SendBytes(messBytes)
			_ = test.ReadBlock("5")
		}

		elapsed := time.Since(start)
		runtime[j] = float64(numberOfRuns) / elapsed.Seconds()
		fmt.Printf("Processed %.2f requests per second\n", runtime[j])
	}

	fmt.Println("----------------\n")
	fmt.Println("Runs completed.\n")
	avg := 0.
	for _, val := range runtime {
		avg += val
	}

	avg /= float64(numberOfBenchs)
	fmt.Printf("Processed %.2f requests per second on average\n", avg)

	test.Close()
}
