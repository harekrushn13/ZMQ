package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	address      = "tcp://*:5555"
	totalNumbers = 20 // Number of messages to send
	ioThreads    = 4  // Number of I/O threads for receiver
)

// Sender: Sends sequential numbers
func sender() {
	push, _ := zmq.NewSocket(zmq.PUSH)
	defer push.Close()
	push.Connect("tcp://localhost:5555")

	for i := 1; i <= totalNumbers; i++ {
		msg := strconv.Itoa(i)
		fmt.Println("Sent:", msg)
		push.Send(msg, 0)
		time.Sleep(100 * time.Millisecond) // Small delay to observe order
	}
}

// Worker: Each worker thread processes received messages
func worker(id int, pull *zmq.Socket) {
	for {
		msg, err := pull.Recv(0)
		if err != nil {
			log.Println("Worker", id, "Error:", err)
			continue
		}
		fmt.Printf("Worker %d processed: %s\n", id, msg)
	}
}

// Receiver: Uses multiple I/O threads to handle incoming messages
func receiver() {
	// Set up the context with multiple I/O threads
	context, _ := zmq.NewContext()
	context.SetIoThreads(ioThreads)

	// Create a PULL socket
	pull, _ := context.NewSocket(zmq.PULL)
	defer pull.Close()
	pull.Bind(address)

	// Start worker threads
	for i := 1; i <= ioThreads; i++ {
		go worker(i, pull)
	}

	// Keep the receiver running
	select {}
}

func main() {
	// Start receiver with multiple I/O threads
	go receiver()
	time.Sleep(time.Second) // Give receiver time to initialize

	// Start sender
	sender()
}
