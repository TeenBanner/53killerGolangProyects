package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}

	log.Printf("request to: %v", r.URL.Path)

	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST REQUEST\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s", address)

}
func main() {
	// create a fileserver
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Listeningon port http://Localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
