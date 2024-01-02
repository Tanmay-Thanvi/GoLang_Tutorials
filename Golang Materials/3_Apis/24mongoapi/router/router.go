package router

import (
	"github.com/Tanmay-Thanvi/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	router.HandleFunc("/api/movie/", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
