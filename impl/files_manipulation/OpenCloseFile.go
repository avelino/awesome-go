package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Dosyaları Açma ve Kapama (Open and Close Files)
	*/

	// Dosyayı salt okunur olarak açtık
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// OpenFile çok seçenekli dosya açma yöntemidir.
	// İkinci parametre dosya açılış amacını ayarlarken, üçüncü parametre dosya izinlerini belirler.
	file, err = os.OpenFile("demo.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	/*

		   OpenFile() ikinci parametrenin tipleri;

		   os.O_RDONLY   : Sadece okuma
		   os.O_WRONLY   : Sadece yazma
		   os.O_RDWR     : Okuma ve yazma yapılabilir
		   os.O_APPEND   : Dosyanın sonuna ekle
		   os.O_CREATE   : Dosya yoksa oluştur
		   os.O_TRUNC    : Açılırken dosyayı kes

		   Bu ayarlar birden fazla olarak da kullanılabilir

		   	-> os.O_CREATE|os.O_APPEND
			-> os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	*/
}
