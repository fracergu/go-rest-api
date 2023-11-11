package routes

import (
	"go-rest-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
