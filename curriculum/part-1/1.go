package main

import (
	"github.com/kettek/goro"
	"log"
)

func main() {
	// Initialize goro!
	if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	// Quit goro!
	goro.Quit()
}
