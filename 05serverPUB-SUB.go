package main

import (
	"bufio"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"strings"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	skt, _ := zmq.NewSocket(zmq.PUB)
	defer skt.Close()

	skt.Bind("tcp://*:5555")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Publish Topic: ")
		topic, _ := reader.ReadString('\n')
		topic = strings.TrimSpace(topic)

		fmt.Print("Enter Message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		skt.Send(fmt.Sprintf("%s %s", topic, message), 0)
	}
}
