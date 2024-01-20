package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Yeniden İsimlendirme ve Taşıma (Rename and Move a File)
	*/

	originalPath := "demo.txt"
	newPath := "test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// Taşıma işlemini de Rename() ile yapmayı dene!
}
