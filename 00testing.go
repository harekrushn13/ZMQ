package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	// Create a PUB socket with a High-Water Mark of 5
	pub, _ := zmq.NewSocket(zmq.PUB)
	defer pub.Close()
	pub.SetSndhwm(5) // Set send queue to a maximum of 5 messages
	pub.Bind("tcp://*:5556")

	// Create a SUB socket (simulating a slow subscriber)
	sub, _ := zmq.NewSocket(zmq.SUB)
	defer sub.Close()
	sub.SetRcvhwm(3)     // Set receive queue to a maximum of 3 messages
	sub.SetSubscribe("") // Subscribe to all messages
	sub.Connect("tcp://localhost:5556")

	go func() {
		// Simulate a slow subscriber
		for {
			msg, err := sub.Recv(0)
			if err == nil {
				fmt.Println("Received:", msg)
				//time.Sleep(2 * time.Second) // Simulate slow processing
			}
		}
	}()

	// Publisher sends messages
	for i := 1; i <= 20; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println("Sending:", msg)
		pub.Send(msg, 0)
		time.Sleep(500 * time.Millisecond) // Simulate message frequency
	}
}
