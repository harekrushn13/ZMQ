package main

import (
	"bufio"
	"fmt"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	sender, _ := zctx.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5555")

	receiver, _ := zctx.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5556")

	fmt.Println("Connected to chat server.")

	go func() {
		for {
			msg, _ := receiver.Recv(0)
			fmt.Println("Server:", msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()

		sender.Send(line, 0)
	}
}
