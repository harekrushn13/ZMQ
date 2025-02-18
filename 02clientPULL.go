package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	pull, _ := zmq.NewSocket(zmq.PULL)
	defer pull.Close()
	pull.Connect("tcp://localhost:5555")

	for {
		msg, _ := pull.Recv(0)
		fmt.Println(msg)
	}
}
