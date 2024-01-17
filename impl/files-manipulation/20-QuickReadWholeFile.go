package main

import (
	"io/ioutil"
	"log"
)

func main() {

	/*
	   Dosyayı Hızlı Okuma (Quick Read Whole File to Memory)
	*/

	// Dosyadan byte dilimine okuma yapma
	data, err := ioutil.ReadFile("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}
