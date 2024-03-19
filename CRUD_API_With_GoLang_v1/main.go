package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"Title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, v := range Movies {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			break
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprintf(w, "Method Nor Alowed")
	}
	var movie Movie
	w.Header().Set("Content-type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&movie) //decoding request body to json body and insert it into movie
	movie.ID = strconv.Itoa(rand.Intn(100000))
	Movies = append(Movies, movie)

	_ = json.NewEncoder(w).Encode(Movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idparam := mux.Vars(r)
	var movie *Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	for i, m := range Movies {
		if m.ID == idparam["id"] {
			Movies[i] = *movie
			break
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for MovieIndex, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:MovieIndex], Movies[MovieIndex+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{ID: "1", Isbn: "43573", Title: "Star Wars", Director: &Director{FirstName: "Jhon", LastName: "Doe"}})
	Movies = append(Movies, Movie{ID: "2", Isbn: "3214", Title: "Dune", Director: &Director{FirstName: "marcus", LastName: "Alonso"}})

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", GetMovieById).Methods("GET")
	r.HandleFunc("/CreateMovie", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
