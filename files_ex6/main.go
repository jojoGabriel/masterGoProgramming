package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"info.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	textToWrite := []byte("The Go gopher is an iconic mascot!")
	_, err = file.Write(textToWrite)

	if err != nil {
		log.Fatal(err)
	}

}
