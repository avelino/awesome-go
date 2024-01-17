package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	/*
	   Dosya Bilgisi Almak (Get File Info)
	*/

	// Dosya bilgisini döndürür.
	// Eğer dosya yoksa hata döndürür.
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("System interface type: %T\n", fileInfo.Sys())
	fmt.Println("System info: %+v\n\n ", fileInfo.Sys())
}
