package main

import (
	"log"
	"os"
)

func Ls() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		log.Println(e.Name())
	}
}
