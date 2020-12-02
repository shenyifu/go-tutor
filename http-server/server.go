package main

import (
	"log"
	"net/http"
)

func main() {
	handler := &PlayServer{}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen to port %v", err)
	}
}
