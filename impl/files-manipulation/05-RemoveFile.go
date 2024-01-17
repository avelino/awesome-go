package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Dosya Silme (Delete a file)
	*/

	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
