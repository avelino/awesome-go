package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {

	/*
	   Uncompress a File
	   Uncompress means = never compressed
	   Decompress correct word.
	*/

	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
