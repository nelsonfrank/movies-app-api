package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	movie "github.com/nelsonfrank/movies-app-api/movie"
)



func main(){
	r := mux.NewRouter()

	movie.SeedingData()

	r.HandleFunc("/movies", movie.GetMoviesContoller).Methods("GET")
	r.HandleFunc("/movies/{id}", movie.GetMovieByIdContoller).Methods("GET")
	r.HandleFunc("/movies", movie.CreateMovieContoller).Methods("POST")
	r.HandleFunc("/movies/{id}", movie.UpdateMovieContoller).Methods("PUT")
	r.HandleFunc("/movies/{id}", movie.DeleteMovieContoller).Methods("DELETE")


	fmt.Printf("Starting Server at port 8000 \n")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
 