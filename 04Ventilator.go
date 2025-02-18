package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

const totalNumbers = 100000
const batchSize = 1000

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	sender, _ := cnt.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5555")

	fmt.Println("Ventilator: Sending 100,000 numbers in batches...")

	rand.Seed(time.Now().UnixNano())

	// Send 100,000 numbers in batches
	for i := 0; i < totalNumbers; i += batchSize {
		batch := ""
		for j := 0; j < batchSize; j++ {
			batch += fmt.Sprintf("%d ", j+1)
		}
		sender.Send(batch, 0)
	}
	fmt.Println("Ventilator: Done sending numbers!")
}
