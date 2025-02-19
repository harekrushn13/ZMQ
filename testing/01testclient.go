package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	cnt, _ := zmq.NewContext()
	defer cnt.Term()

	skt, _ := zmq.NewSocket(zmq.PULL)
	defer skt.Close()
	skt.Connect("tcp://localhost:5555")

	fmt.Println("Client started, receiving file...")

	file, err := os.Create("rtest.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		chunk, err := skt.RecvBytes(0)
		fmt.Printf("Chunk received: %d\n", len(chunk))
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		//fmt.Printf("Chunk: %s\n", string(chunk))
		_, err = file.Write(chunk)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			break
		}
	}

	fmt.Println("File received successfully!")

	//data, err := skt.RecvBytes(0)
	//fmt.Printf("Received %d bytes\n", len(data))
	//if err != nil {
	//	fmt.Printf("Failed to receive file: %v", err)
	//}
	//
	//err = os.WriteFile("rtest.txt", data, 0644)
	//if err != nil {
	//	fmt.Printf("Failed to write file: %v", err)
	//}
	//
	//fmt.Println("File received")
}
