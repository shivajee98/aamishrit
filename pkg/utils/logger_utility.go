package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[LOG] ", log.Ldate|log.Ltime|log.Lshortfile)

func Info(msg string) {
	logger.Println("[INFO] " + msg)
}

func Error(err error) {
	if err != nil {
		logger.Println("[ERROR] ", err)
	}
}
