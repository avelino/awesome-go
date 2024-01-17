package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*
	   Seek Positions in File
	*/

	file, _ := os.Open("demo.txt")
	defer file.Close()

	// Offset kaç byte taşımalıdır?
	// Offset negatif(-) ya da pozitif(+) olabilir
	var offset int64 = 5

	// Offset için referans noktası nerededir?
	// 0 = Dosyanın başlangıcı
	// 1 = Şu anki pozisyon
	// 2 = Dosyanın sonu
	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5: ", newPosition)

	// Şu anki pozisyondan 2 byte geri git
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just momved back two: ", newPosition)

	// 0 byte taşındıktan sonra Seek'den dönüş değerini elde ederek geçerli konumu bul.
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position: ", currentPosition)

	// Dosyanın başlangıcına git
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0: ", newPosition)
}
