package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	/*
	   Downloading a File Over HTTP
	*/

	// Create output file
	newFile, err := os.Create("courses.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request dijibil.com
	url := "http://www.dijibil.com/courses"
	response, err := http.Get(url)
	defer response.Body.Close()

	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}
