package main

import (
	"log"
	"os"
)

func main() {

	/*
			   Dosyayı Kes (Truncate a File)

		       Bir dosya 100 byte'a kesilmelidir.
		       Eğer dosya 100 byte'tan az ise içerik kalır, geri kalan kısım boş byte ile dolacaktır.
		       Eğer dosya 100 byte'ın üzerinde ise 100 byte'tan sonraki herşey kaybolur.
		       Her iki durumda da kesme işlemi 100 byte üzerinden gerçekleştirilmelidir.
		       Eğer kesilecek dosya boş ise sıfır değeri kullanılır.
	*/

	err := os.Truncate("demo.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
