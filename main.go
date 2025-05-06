package main

import (
	"downmann/internal/download"
	"log"
)

func main() {
	log.Println("Started Downmann")

	download, err := download.NewDownload("https://dl.google.com/go/go1.24.2.linux-amd64.tar.gz")
	if err != nil {
		panic(err)
	}

	err = download.Start()
	// download.Debug()
	if err != nil {
		panic(err)
	}
}
