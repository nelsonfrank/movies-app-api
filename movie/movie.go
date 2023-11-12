package movie

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func SeedingData(){
	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "Movie 0ne", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4321", Title: "Movie Two", Director: &Director{FirstName: "Joel", LastName: "Munich"}})
}

func setRequestHeader(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
}

func GetMoviesContoller(w http.ResponseWriter, r *http.Request){
	setRequestHeader(w)
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByIdContoller(w http.ResponseWriter, r *http.Request){
	setRequestHeader(w)

	// get params from request 
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"]{
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func CreateMovieContoller(w http.ResponseWriter, r *http.Request){
	setRequestHeader(w)

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID =  strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovieContoller(w http.ResponseWriter, r *http.Request){
	setRequestHeader(w)

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]... )
			
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
		}
	}

}

func DeleteMovieContoller(w http.ResponseWriter, r *http.Request){
	setRequestHeader(w)

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]... )
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}
