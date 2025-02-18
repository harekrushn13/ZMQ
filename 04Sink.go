package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strconv"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5556")

	fmt.Println("Sink: Collecting worker averages...")

	totalSum := 0
	totalCount := 0

	for {
		// Receive worker average
		avgStr, _ := receiver.Recv(0)

		workerAvg, _ := strconv.Atoi(avgStr)
		totalSum += workerAvg
		totalCount++

		// Stop after all batches are processed
		if totalCount == (100000 / 1000) { // 100 workers sending 1 average each
			break
		}
	}

	// Compute final average
	finalAvg := totalSum / totalCount
	fmt.Printf("Sink: Final average of 100,000 numbers = %d\n", finalAvg)
}
