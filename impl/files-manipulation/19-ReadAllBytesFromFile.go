package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	/*
	   Dosyanın Tüm Byte'ları Oku (Read All Bytes of File)
	*/

	// Okumak için dosya aç
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read(), io.ReadFull(), ioutil.ReadAll() ve io.ReadAtLeast() da kullanılabilir
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: &s\n", data)
	fmt.Printf("Number of bytes read:", len(data))
}
