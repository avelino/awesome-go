package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Okuma ve Yazma İzinlerini Kontrol Etmek ( Check Read and Write Permissions)
	*/

	// Yazma izinleri testi
	// Dosyanyı açmaya çalışmadan önce varlığını kontrol edebilirsiniz;
	// Bunun için; os.IsNotExist
	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Hata : Yazma izni reddedildi")
		}
	}
	file.Close()

	// Okuma izinleri testi
	file, err = os.OpenFile("demo.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Hata : Okuma izni reddedildi")
		}
	}
	file.Close()
}
