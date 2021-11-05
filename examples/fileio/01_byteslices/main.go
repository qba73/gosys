package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "test"
	// Create a file

	data := []byte("This is data!\n")
	if err := ioutil.WriteFile(fileName, data, 0777); err != nil {
		log.Println(err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	moreData := make([]byte, 200)
	n, err := f.Read(moreData)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, moreData)
}
