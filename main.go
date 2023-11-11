package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
	"log"
	"net/http"

	_ "go-rest-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example API
// @description This is an example API for Go REST API with Swagger and GORM
// @version 1
// @host localhost:8000
// @BasePath /api/v1
func main() {
	database.ConnectDatabase()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	log.Println("Server running on port 8000")
	log.Println("Documentation running on http://localhost:8000/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8000", router))
}
