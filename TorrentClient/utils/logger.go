package utils

import (
	"io"
	"log"
	"os"
)

// InitLogger initializes the logger to write to a file and standard output.
func InitLogger() {
	logFile, err := os.OpenFile("torrent-client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
