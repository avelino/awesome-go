package main

import (
	"log"
	"os"
	"time"
)

func main() {

	/*
	   İzinleri, Sahipliği ve Zaman Damgalarını(Timestamps) Değiştirmek Change Permissions, Ownership, and Timestamps
	*/

	// İzinleri değiştirme (Linux tarzı)
	err := os.Chmod("demo.txt", 0777)
	if err != nil {
		log.Println(err)

	}

	// Sahipliği değiştirme
	err = os.Chown("demo.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Zaman Damgalarını(Timestamps) Değiştirme
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("demo.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
