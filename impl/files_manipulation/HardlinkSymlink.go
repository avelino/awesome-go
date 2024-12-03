package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*
	   Hard Links and Symlinks
	*/

	// Bir "Hard Link" oluştur
	// Aynı içeriğe işaret eden iki dosya adı olacaktır.
	// Birinin içeriğini değiştirmek diğerini değiştirecek
	// Birini silmek / yeniden adlandırmak diğerini etkilemez
	err := os.Link("demo.txt", "demo_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sym oluşturma")

	// Symlink oluştur
	err = os.Symlink("demo.txt", "demo_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat dosya bilgisini döndürür.
	// Ama aslında o bir symlink ise symlink hakkında bilgi döndürür.
	// Bağlantıyı takip etmeyecek ve gerçek dosya hakkında bilgi vermeyecektir.
	// Symlink'ler Windows'da çalışmaz.
	fileInfo, err := os.Lstat("demo_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bağlantı Bilgisi: %+v", fileInfo)

	// Sadece bir symlink'in sahipliğini değiştir.
	// Ama işaret ettiği dosyanın sahipliğini değiştirmez.
	err = os.Lchown("demo_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
