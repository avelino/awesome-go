package main

import (
	"io"
	"log"
	"os"
)

func main() {

	/*
	   Dosya Kopyalama (Copy a file)
	*/

	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Yeni bir dosya oluştur
	newFile, err := os.Create("demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Byte'ları kaynaktan hedefe Kopyalama
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Dosya içeriğini işle(commit)
	// Belleği diske boşalt(flush)
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
