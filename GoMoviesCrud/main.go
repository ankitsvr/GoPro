package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"Director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//global variable
var movies []Movie

//Start--- function to handle the routes
// function to handle request to fetch all the movies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return	
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn((100000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request){

	//set the content type
	w.Header().Set("Content-Type", "application/json")

	//fetch the parameters
	params := mux.Vars(r)

	//loops overs the movies range

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return

		}

	//delete the movie with the ID that you've sent
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
//END----function to handle the routes




//starting of main function
func main() {
	r := mux.NewRouter()
 	movies = append(movies, Movie{ID: "1", Isbn:"43564", Title:"Van Helsingh", Director: &Director{Firstname:"john", Lastname:"wick"}})
	movies = append(movies, Movie{ID: "2", Isbn:"32787", Title:"Mission impossible", Director: &Director{Firstname:"ankit", Lastname:"thakur"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("UPDATE")
	r.HandleFunc("/movies/{id)", deleteMovie).Methods("DELETE")

	fmt.Printf("starting the server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
	
}
