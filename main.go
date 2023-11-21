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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.Id == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movie {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.Id = strconv.Itoa(rand.Intn(10000))
	movie = append(movie, mov)
	json.NewEncoder(w).Encode(mov)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.Id == params["id"] {
			// movie = append(movie[:index], movie[index+1:]...)
			var mov Movie

			_ = json.NewDecoder(r.Body).Decode(&mov)
			mov.Id = params["id"]
			movie[index] = mov
			// movie = append(movie, mov)
			json.NewEncoder(w).Encode(mov)
			return
		}
	}

}
func main() {
	r := mux.NewRouter()

	movie = append(movie, Movie{Id: "1", Isbn: "438", Title: "1", Director: &Director{FirstName: "Sourav", LastName: "Dey"}})
	movie = append(movie, Movie{Id: "2", Isbn: "438", Title: "2", Director: &Director{FirstName: "Sourav", LastName: "Roy"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("starting server at 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
