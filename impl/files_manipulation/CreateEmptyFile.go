package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {

	/*
	   Boş Dosya Oluşturma (Create Empty File)
	*/

	newFile, err = os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
