package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strings"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5555")

	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Bind("tcp://*:5556")

	fmt.Println("Server started...")

	for {
		msg, _ := receiver.Recv(0)
		fmt.Println("Message Received: ", string(msg))

		parts := strings.SplitN(msg, " ", 2)
		if len(parts) < 2 {
			continue
		}

		clientName := parts[0]
		msgContent := parts[1]

		msgparts := strings.SplitN(msgContent, " ", 2)
		if len(parts) < 2 {
			continue
		}
		recipient := strings.TrimPrefix(msgparts[0], "@")
		msgBody := msgparts[1]

		if recipient == "all" {
			sender.Send(fmt.Sprintf("all %s %s", clientName, msgBody), 0)
		} else {
			sender.Send(fmt.Sprintf("%s %s %s", recipient, clientName, msgBody), 0)
		}
	}
}
