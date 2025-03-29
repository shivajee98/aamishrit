package utils

import (
	"fmt"
	"log"
	"os"
)

func CheckError(messege string, err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "Error %s: %v\n", messege, err)
		os.Exit(1)
	}
}
