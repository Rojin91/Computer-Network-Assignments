package utils

import (
	"log"
	"os"
)

// InitLogger initializes the logging system, directing logs to standard output and including date, time, and source file information.
func InitLogger() {
	// Set the output destination for the logger to standard output (console)
	log.SetOutput(os.Stdout)
	
	// Set the flags for the logger to include the date, time, and source file information in each log message
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
 