package chat

import (
	"bufio"
	"fmt"
	"github.com/harekrushn13/ZMQ/chatapp/config"
	zmq "github.com/pebbe/zmq4"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	serverIP := "10.20.41.40"

	fmt.Print("Enter client name: ")

	clientName, _ := reader.ReadString('\n')

	clientName = customTrim(clientName)

	context, err := zmq.NewContext()

	if err != nil {

		log.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fmt.Println("Client started")

	go func(context *zmq.Context, clienName string, serverIP string) {

		push, err := context.NewSocket(zmq.PUSH)

		if err != nil {

			log.Fatalln(err)
		}

		err = push.Connect("tcp://" + serverIP + ":6000")

		if err != nil {

			log.Fatalln(err)
		}

		for {

			if config.GlobalShutdownClient {

				push.Close()

				return
			}

			fmt.Print("> ")

			text, _ := reader.ReadString('\n')

			text = customTrim(text)

			if len(text) == 0 || text[0] != '@' {

				fmt.Println("Enter valid format")

				continue
			}

			recipient, msgBody := customSplit(text)

			if recipient == "" || msgBody == "" {

				fmt.Println("Invalid message format.")

				continue
			}

			fullMsg := recipient + " " + clientName + " " + msgBody

			push.Send(fullMsg, 0)
		}

	}(context, clientName, serverIP)

	go func(context *zmq.Context, clienName string, serverIP string) {

		sub, err := context.NewSocket(zmq.SUB)

		if err != nil {

			log.Fatalln(err)
		}

		err = sub.Connect("tcp://" + serverIP + ":6001")

		if err != nil {

			log.Fatalln(err)
		}

		sub.SetSubscribe("all")

		sub.SetSubscribe(clientName)

		for {
			if config.GlobalShutdownClient {

				sub.Close()

				return
			}

			msg, err := sub.Recv(0)

			if err != nil {

				continue
			}

			fmt.Println("Received:", msg)
		}
	}(context, clientName, serverIP)

	<-sigChan
	fmt.Println("\nShutting down client...")

	config.GlobalShutdownClient = true

	context.Term()

	//time.Sleep(1 * time.Second)

	fmt.Println("Client stopped")
}

func customSplit(text string) (string, string) {

	ind := strings.IndexByte(text, ' ')

	if ind == -1 {

		return "", ""
	}

	return text[1:ind], text[ind+1:]
}

func customTrim(s string) string {

	n := len(s)

	if n > 0 && s[n-1] == '\n' {

		return s[:n-1]
	}

	return s
}
