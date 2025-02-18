package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	push, _ := zmq.NewSocket(zmq.PUSH)
	defer push.Close()
	push.Bind("tcp://*:5555")

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Task #%d", i)
		push.Send(msg, 0)
		fmt.Println("Sent:", msg)
		time.Sleep(time.Second)
	}

}
