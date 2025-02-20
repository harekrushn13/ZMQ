package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf16"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("File Discriptor: %v\n", file.Fd())

	//write file in utf-16
	bom := []byte{0xFF, 0xFE}
	_, err = file.Write(bom)
	if err != nil {
		panic(err)
	}

	longData := strings.Repeat("abcdefghij", 2*1024*1024*1024/10) // 2 GB
	dataSize := int64(len(longData))
	chunkSize := int64(100 * 1024 * 1024) // 100 MB
	//var totalSize int64 = 2 * 1024 * 1024 * 1024                  // 2 GB

	//fmt.Printf("Chunk Type: %T, length:%d %d\n", chunk, len(chunk), utf8.RuneCountInString(chunk))

	//for i := int64(0); i < totalSize/2; i += int64(chunkSize) {
	for i := int64(0); i < dataSize/chunkSize; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > dataSize {
			end = dataSize
		}
		chunk := longData[start:end]

		utf16Data := utf16.Encode([]rune(chunk))
		//fmt.Printf("Chunk received: %d %T\n", len(utf16Data), utf16Data)
		//fmt.Println(utf16Data[:20])
		//fmt.Println(string(utf16Data[:21]))

		utf16Bytes := make([]byte, len(utf16Data)*2)
		//fmt.Printf("Buffer sized: %d\n", len(utf16Bytes))
		//for j, val := range utf16Data {
		//	//fmt.Printf("j: %d val:%d\n", j, val)
		//	binary.LittleEndian.PutUint16(utf16Bytes[j*2:], val) // Little-endian encoding
		//}
		for j, val := range utf16Data {
			utf16Bytes[j*2] = byte(val) // Lower byte
			utf16Bytes[j*2+1] = byte(0) // Upper byte
			//fmt.Printf("%v %v:%v\n", i, byte(val), byte(val<<8))
			//if i == 10 {
			//	return
			//}
		}
		//fmt.Println(utf16Bytes[:20])
		//return

		_, err := file.Write(utf16Bytes)
		if err != nil {
			panic(err)
		}

		//_, err := file.WriteString(chunk)
		//if err != nil {
		//	panic(err)
		//}
	}
}
