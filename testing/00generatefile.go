package main

import (
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunk := strings.Repeat("abcdefghj", 1024*1024) // 10MB
	chunkSize := len(chunk)
	var totalSize int64 = 2 * 1024 * 1024 * 1024 // 2 GB

	for i := int64(0); i < totalSize; i += int64(chunkSize) {
		_, err := file.WriteString(chunk)
		if err != nil {
			panic(err)
		}
	}
}
