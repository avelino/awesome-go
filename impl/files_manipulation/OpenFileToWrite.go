package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Yazmak için dosyayı aç
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Dosyadan bir "buffered writer" oluştur
	bufferedWriter := bufio.NewWriter(file)

	// Arabelleğe byte yaz
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// String'i belleğe(buffer) yaz
	// Ayrıca bunlarda mevcut: WriteRune() ve WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// Buffer beklenirken ne kadar depolandığını kontrol etme
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// Kullanılabilir arabellek miktarını görme
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Belleği diske yaz
	bufferedWriter.Flush()

	// Henüz Flush() ile dosyaya yazılmamış arabellekte yapılan değişiklikleri geri alın.
	bufferedWriter.Reset(bufferedWriter)

	// Kullanılabilir arabellek miktarını görme
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Tamponu yeniden boyutlandırır.
	// Bu kullanınmda aynı buffer'ı yeniden boyutlandırıp kullanıyoruz.
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// Yeniden boyutlandırma sonrasında tekrar kullanılabilir arabellek miktarını kontrol etme
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
}
