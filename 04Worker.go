package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strconv"
	"strings"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5555")

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5556")

	fmt.Println("Worker: Ready to process batches...")

	for {
		// Receive batch of numbers
		batch, _ := receiver.Recv(0)

		// Convert batch to numbers and calculate average
		nums := strings.Fields(batch) // Split batch into numbers
		sum := 0
		count := len(nums)

		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			sum += num
		}

		workerAvg := sum / count

		// Send the batch average to the sink
		sender.Send(fmt.Sprintf("%d", workerAvg), 0)
	}
}
