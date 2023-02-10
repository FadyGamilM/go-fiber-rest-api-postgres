package routers

import (
	"github.com/FadyGamilM/gomongo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAll).Methods("GET")
	router.HandleFunc("/api/movies", controllers.Create).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies", controllers.DeleteAll).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOne).Methods("DELETE")

	return router
}
