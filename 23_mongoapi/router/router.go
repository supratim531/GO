package router

import (
	"github.com/gorilla/mux"
	"github.com/supratim531/gomongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie", controller.InsertOneMovie).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
