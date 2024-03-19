package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"Title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firdtname"`
	LastName  string `json:"lastname"`
}

var Movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for MovieIndex, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:MovieIndex], Movies[MovieIndex+1:]...)
		}
	}
}

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{ID: "1", Isbn: "43573", Title: "Star Wars", Director: &Director{FirstName: "Jhon", LastName: "Doe"}})
	Movies = append(Movies, Movie{ID: "2", Isbn: "3214", Title: "Dune", Director: &Director{FirstName: "marcus", LastName: "Alonso"}})

	r.HandleFunc("movies", GetMovies).Methods("GET")
	r.HandleFunc("movie/{id}", GetMovieById).Methods("GET")
	r.HandleFunc("CreateMovie", CreateMovie).Methods("POST")
	r.HandleFunc("movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
