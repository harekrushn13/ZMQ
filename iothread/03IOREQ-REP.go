package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	reqAddress = "tcp://localhost:5555"
	repAddress = "tcp://*:5555"
	totalMsgs  = 10 // Number of messages to send
)

// Server (REP): Receives requests and replies
func server() {
	rep, _ := zmq.NewSocket(zmq.REP)
	defer rep.Close()
	rep.Bind(repAddress)

	for i := 1; i <= totalMsgs; i++ {
		req, err := rep.Recv(0)
		if err != nil {
			log.Println("Receive Error:", err)
			continue
		}

		fmt.Println("Server received:", req)

		// Simulate processing time
		time.Sleep(200 * time.Millisecond)

		// Send response back
		_, err = rep.Send("Ack: "+req, 0)
		if err != nil {
			log.Println("Reply Error:", err)
		}
	}
}

// Client (REQ): Sends sequential messages
func client() {
	req, _ := zmq.NewSocket(zmq.REQ)
	defer req.Close()
	req.Connect(reqAddress)

	for i := 1; i <= totalMsgs; i++ {
		msg := strconv.Itoa(i)
		fmt.Println("Client sent:", msg)

		_, err := req.Send(msg, 0)
		if err != nil {
			log.Println("Send Error:", err)
			continue
		}

		// Wait for reply
		reply, err := req.Recv(0)
		if err != nil {
			log.Println("Receive Error:", err)
			continue
		}

		fmt.Println("Client received:", reply)

		time.Sleep(100 * time.Millisecond) // Simulate client delay
	}
}

func main() {
	// Start the server
	go server()
	time.Sleep(time.Second) // Give server time to start

	// Start the client
	client()
}
