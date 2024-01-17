package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	/*
	   Dosyanın var olup olmadığını kontrol etmek (Check if File Exists)
	*/

	// Dosya bilgisini döndürür.
	// Eğer dosya yoksa hata döndürür.
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	log.Println("File does exist. File information : ")
	log.Println(fileInfo)
}
