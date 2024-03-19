package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	err := http.ListenAndServe(":8000", fs)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Server runninf on http://localhost:8000")
}
