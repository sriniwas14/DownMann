package configloader

import (
	"log"
	"os"
)

var DestFolder = "./"
var TempDir = "./"
var MaxThreads int64 = 4

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	TempDir = homeDir + "/Temp/"
	DestFolder = homeDir + "/Downloads/"
	log.Println(homeDir)
}

// Create config if not present

// Load Config
