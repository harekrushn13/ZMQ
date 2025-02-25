package chat

import (
	"fmt"
	"github.com/harekrushn13/ZMQ/chatapp/config"
	"github.com/pebbe/zmq4"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	context, err := zmq4.NewContext()

	if err != nil {

		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fmt.Println("Server started")

	go func(ctx *zmq4.Context) {

		pull, err := context.NewSocket(zmq4.PULL)

		if err != nil {

			log.Fatal(err)
		}

		pull.Bind("tcp://*:6000")

		pub, err := context.NewSocket(zmq4.PUB)

		if err != nil {

			log.Fatal(err)
		}

		pub.Bind("tcp://*:6001")

		for {

			if config.GlobalShutdownServer {

				pull.Close()

				pub.Close()

				return
			}

			msg, err := pull.Recv(0)

			if err != nil {
				continue
			}

			fmt.Println("Message Received:", msg)

			pub.Send(msg, 0)
		}

	}(context)

	<-sigChan
	fmt.Println("\nShutting down server...")

	config.GlobalShutdownServer = true

	context.Term()

	//time.Sleep(1 * time.Second)

	fmt.Println("Server stopped")
}
