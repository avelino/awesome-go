package main

import (
	"os"
	"log"
)

func main() {

	/*
	   Hızlıca Dosya Yazma (Quick Write to File)
	*/

	err := os.WriteFile("demo.txt", []byte("Hi!\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
