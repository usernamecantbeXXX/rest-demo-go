package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-demo-go/controllers"
)

func main() {
	//initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initHandlers(router *mux.Router) {
	//router.PathPrefix("/tasks")
	router.HandleFunc("/tasks/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/get", controllers.GetAllTask).Methods("GET`")
	router.HandleFunc("/tasks/get/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/update/{id}", controllers.UpdateTaskByID).Methods("PUT")
	router.HandleFunc("/tasks/delete/{id}", controllers.DeleteTaskByID).Methods("DELETE")
}

//func initDB() {
//	config :=
//		database.Config{
//			ServerName: "localhost:3306",
//			User:       "root",
//			Password:   "root",
//			DB:         "learning_demo",
//		}
//
//	connectionString := database.GetConnectionString(config)
//	err := database.Connect(connectionString)
//	if err != nil {
//		panic(err.Error())
//	}
//	database.Migrate(&entity.Task{})
//}
