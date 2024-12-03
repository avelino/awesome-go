package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Dosyadan n byte kadarlık akışı okuma (Read up to n Bytes from File)
	*/

	// Okumak için dosyayı aç
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// len(b) kadar byte okuma yapılabilir.
	// Dosya sonu io.EOF hata türünü döndürür.
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
