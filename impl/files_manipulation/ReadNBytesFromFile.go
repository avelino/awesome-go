package main

import (
	"io"
	"log"
	"os"
)

func main() {

	/*
	   En az n byte oku (Read At Least n Bytes)
	*/

	// Okumak için dosyayı aç
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 512)
	minBytes := 8
	// io.ReadAtLeast() okumak için en az minBytes bulamıyorsa hata döndürecektir.
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
