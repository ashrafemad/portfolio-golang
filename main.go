package main

import (
	"log"
	"net/http"
	"portfolio/controller"
	"portfolio/db"
	"portfolio/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvVars()
	connectToDatabase()
	http.ListenAndServe(":8000", getRouters())
}

func connectToDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&models.Project{})
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getRouters() *mux.Router {
	routers := mux.NewRouter()
	routers.StrictSlash(true)
	routers.HandleFunc("/projects", controller.CreateProject).Methods("POST")
	routers.HandleFunc("/projects", controller.ListProjects).Methods("GET")
	routers.HandleFunc("/projects/{id}", controller.RetrieveProject).Methods("GET")
	routers.HandleFunc("/projects/{id}", controller.DeleteProject).Methods("DELETE")
	return routers
}
