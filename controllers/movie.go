package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FadyGamilM/gomongo/database"
	"github.com/FadyGamilM/gomongo/models"
	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// specify the allowed method
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	movies := database.GetAll()
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		fmt.Println("Error while parsing the result to json")
		log.Fatal(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// specify the allowed method
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// decode the json data to send it to the repository to create a new instance
	newMovie := models.Movie{}
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		fmt.Println("Error while decoding json data from user")
		log.Fatal(err)
	}

	// now call the repo
	database.InsertOne(newMovie)

	// now return the response
	json.NewEncoder(w).Encode(newMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// specify the allowed method
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// get the id
	params := mux.Vars(r)

	// call the repo
	database.UpdateOneById(params["id"])

	// get the updated one
	// database.Get

	// send the response back
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// specify what method you want to handle for this handler
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// get the id
	params := mux.Vars(r)

	// call the repo
	database.DeleteById(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}
