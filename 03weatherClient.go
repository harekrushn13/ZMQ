package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer socket.Close()
	defer context.Term()

	socket.Connect("tcp://localhost:5555")
	socket.SetSubscribe("")

	fmt.Println("Waithing for weather update")

	for {
		msg, _ := socket.Recv(0)
		fmt.Println("Received Weather update: ", msg)
	}
}
