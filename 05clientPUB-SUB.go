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

	skt, _ := zmq.NewSocket(zmq.SUB)
	defer skt.Close()

	skt.Connect("tcp://localhost:5555")
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter listening topic : ")
	topic, _ := reader.ReadString('\n')
	topic = strings.TrimSpace(topic)

	skt.SetSubscribe(topic)
	skt.SetSubscribe("broadcast")

	for {
		data, _ := skt.Recv(0)
		fmt.Println("Message Received: ", string(data))
	}
}
