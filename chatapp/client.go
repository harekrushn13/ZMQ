package main

import (
	"bufio"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter server IP address: ")
	serverIP, _ := reader.ReadString('\n')
	serverIP = strings.TrimSpace(serverIP)

	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	var clientName string
	fmt.Print("Enter client name: ")
	fmt.Scanln(&clientName)

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://" + serverIP + ":5555")

	fmt.Println("Client Started...")

	go func(cnt *zmq.Context, clientName string, serverIP string) {
		receiver, _ := cnt.NewSocket(zmq.SUB)
		defer receiver.Close()
		receiver.Connect("tcp://" + serverIP + ":5556")

		receiver.SetSubscribe("all")
		receiver.SetSubscribe(clientName)

		for {
			msg, _ := receiver.Recv(0)
			fmt.Println(msg)
		}
	}(cnt, clientName, serverIP)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.HasPrefix(text, "@") {
			sender.Send(clientName+" "+text, 0)
		} else {
			fmt.Println("No user found !!")
		}
	}
}
