package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"time"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	skt, _ := zmq.NewSocket(zmq.PUSH)
	defer skt.Close()
	skt.Bind("tcp://*:5555")

	fmt.Println("Server started, sending file...")

	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 500*1024*1024) // 1 GB
	for {
		n, err := file.Read(buffer)
		fmt.Printf("Buffer length : %d\n", len(buffer))
		fmt.Printf("Read %d bytes\n", n)
		if err != nil {
			break
		}

		_, err = skt.SendBytes(buffer[:n], 0)
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("File sent successfully!")

	//file, err := os.ReadFile("test.txt")
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//}
	//fmt.Printf("Sent %d bytes\n", len(file))
	//_, err = skt.SendBytes(file, 0)
	//if err != nil {
	//	fmt.Println("Error sending file:", err)
	//}
	//time.Sleep(1 * time.Second)
	//
	//fmt.Println("File sent")
}
