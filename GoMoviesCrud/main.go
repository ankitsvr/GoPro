package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`


}

//global variable
var movies []Movie





//Start--- function to handle the routes
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request){

}
func createMovie(w http.ResponseWriter, r *http.Request){

}

func updateMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}
//END----function to handle the routes




//starting of main function
func main() {
	r := mux.NewRouter()
 	movies = append(movies , Movie{ID: "1", Isbn:"43564", Title:"Van Helsingh", Director: &Director{firstname:"john", lastname:"wick"}})
	movies = append(movies , Movie{ID: "2", Isbn:"32787", Title:"Mission impossible", Director: &Director{"firstname:"ankit",lastname:"thakur"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("UPDATE")
	r.HandleFunc("/movies/{id)", deleteMovie).Methods("DELETE")

	fmt.Printf("starting the server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
	
}
