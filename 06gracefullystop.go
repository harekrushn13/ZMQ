package main

import (
	"context"
	"fmt"
	"github.com/pebbe/zmq4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cnt, _ := zmq4.NewContext()
	defer cnt.Term()

	skt, _ := cnt.NewSocket(zmq4.REP)
	defer skt.Close()
	skt.Bind("tcp://*:5555")
	fmt.Println("Server started...")

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, skt)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	<-sign
	fmt.Println("Server exiting...")

	cancel()

	time.Sleep(time.Second)
	fmt.Println("Server exited cleanly...")
}
func worker(ctx context.Context, skt *zmq4.Socket) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped...")
			return
		default:
			msg, err := skt.Recv(0)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Message received: %s\n", msg)
			skt.Send("Reply", 0)
		}
	}
}
