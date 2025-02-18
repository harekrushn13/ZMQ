package main

import (
	"bufio"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	zctx, _ := zmq.NewContext()

	receiver, _ := zctx.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5555")

	sender, _ := zctx.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5556")

	fmt.Println("Chat server started...")

	go func() {
		for {
			msg, _ := receiver.Recv(0)
			fmt.Println("Client:", msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()

		sender.Send(line, 0)
	}
}
