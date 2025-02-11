package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	cnt, _ := zmq.NewContext()

	skt, _ := cnt.NewSocket(zmq.REP)
	//defer skt.Close()

	skt.Bind("tcp://*:5555")
	fmt.Println("Chat server started...")

	for {
		msg, _ := skt.Recv(0)
		fmt.Println("Client : ", msg)

		var rpy string
		fmt.Scanln(&rpy)

		skt.Send(rpy, 0)
	}
}
