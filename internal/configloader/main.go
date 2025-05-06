package configloader

import (
	"log"
	"os"
)

var DestFolder = "./"
var MaxThreads int64 = 4

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	DestFolder = homeDir
	log.Println(homeDir)
}
