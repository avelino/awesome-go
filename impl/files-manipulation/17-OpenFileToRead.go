package main

import (
	"io"
	"log"
	"os"
)

func main() {

	/*
	   Tam olarak n byte'ı oku (Read Exactly n Bytes)
	*/

	// Okumak için dosyayı aç
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
