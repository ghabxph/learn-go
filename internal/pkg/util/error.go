package util

import (
	"log"
	"os"
)

func catch(e error) {
	if e != nil {
		log.Fatalf("Error: %v\n", e)
		os.Exit(-1)
	}
}
