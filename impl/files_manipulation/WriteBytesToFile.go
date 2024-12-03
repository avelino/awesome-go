package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Byte'ları Bir Dosyaya Yazın (Write Bytes to a File)
	*/

	// Demo.txt dosyasını sadece yazılabilir bir dosya olarak aç
	file, err := os.OpenFile(
		"demo.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
