package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	cnt, _ := zmq.NewContext()

	skt, _ := cnt.NewSocket(zmq.REQ)
	//defer skt.Close()

	skt.Connect("tcp://localhost:5555")

	fmt.Println("Connected to chat server.")
	for {
		var snd string
		fmt.Scanln(&snd)
		skt.Send(snd, 0)

		reply, _ := skt.Recv(0)
		fmt.Println("Server : ", reply)

	}
}
