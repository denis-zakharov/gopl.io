package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd() // cwd, err := os.Getwd() initializes new cwd variable in a local scope.
	if err != nil {
		log.Fatalf("Error os.Getwd failed: %v", err)
	}
	log.Printf("Working dir = %s", cwd)
}

func main() {
	log.Printf("cwd after init: %v", cwd)
}
