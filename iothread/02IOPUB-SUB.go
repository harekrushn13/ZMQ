package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	pubAddress = "tcp://*:5556"
	subAddress = "tcp://localhost:5556"
	totalMsgs  = 20 // Number of messages to send
	ioThreads  = 4  // Number of I/O threads for subscriber
)

// Publisher: Sends sequential numbers
func publisher() {
	pub, _ := zmq.NewSocket(zmq.PUB)
	defer pub.Close()
	pub.Bind(pubAddress)

	time.Sleep(time.Second) // Allow subscribers to connect

	for i := 1; i <= totalMsgs; i++ {
		msg := strconv.Itoa(i)
		fmt.Println("Published:", msg)
		_, err := pub.Send(msg, 0)
		if err != nil {
			log.Println("Publish Error:", err)
		}
		time.Sleep(100 * time.Millisecond) // Delay to observe order
	}
}

// Worker: Each worker thread processes received messages
func worker(id int, sub *zmq.Socket) {
	for {
		msg, err := sub.Recv(0)
		if err != nil {
			log.Println("Worker", id, "Error:", err)
			continue
		}
		fmt.Printf("Worker %d received: %s\n", id, msg)
	}
}

// Subscriber: Uses multiple I/O threads to handle incoming messages
func subscriber() {
	// Set up context with multiple I/O threads
	context, _ := zmq.NewContext()
	context.SetIoThreads(ioThreads)

	// Create a SUB socket
	sub, _ := context.NewSocket(zmq.SUB)
	defer sub.Close()
	sub.Connect(subAddress)
	sub.SetSubscribe("") // Subscribe to all messages

	// Start worker threads
	for i := 1; i <= ioThreads; i++ {
		go worker(i, sub)
	}

	// Keep subscriber running
	select {}
}

func main() {
	// Start subscriber with multiple I/O threads
	go subscriber()
	time.Sleep(time.Second) // Give subscriber time to initialize

	// Start publisher
	publisher()
}
