package utils

import (
	"log"
	"os"
) 

func CreateShutdownLogFile() *os.File {
	file, err := os.OpenFile("./shutdown.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error creating shutdown log file: %v", err)
	}
	return file
}