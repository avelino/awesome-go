package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {

	/*
	   Compress a File
	*/

	// Create .gz file to write to
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compressed data written to file.")
}
