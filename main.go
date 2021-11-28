package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Task Entity
type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DueDate string `json:"dueDate"`
	Status  int    `json:"status"`
}

var tasks []Task = []Task{}

func main() {

	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened tasks.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal([]byte(byteValue), &tasks)

	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initHandlers(router *mux.Router) {
	router.HandleFunc("/tasks/{id}", retrieveTask).Methods("GET")
	router.HandleFunc("/tasks", retrieveTask).Methods("GET")
	router.HandleFunc("/tasks", addTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
}

func retrieveTask(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]

	if idParam == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
		log.Println("getAllTasks: " + r.Form.Get("expiredToday"))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(tasks) {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		return
	}

	task := tasks[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
	log.Println("getAllTasks")
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)

	taskUpdates, _ := json.Marshal(tasks)

	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)
	log.Println("create: ", newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// get the ID of the task from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(tasks) {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		return
	}

	// get the value from JSON body
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	tasks[id] = updatedTask

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)

	taskUpdates, _ := json.Marshal(tasks)

	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)
	log.Println("update: ", updatedTask)
}

//func patchTask(w http.ResponseWriter, r *http.Request) {
//	// get the ID of the task from the route parameters
//	var idParam string = mux.Vars(r)["id"]
//	id, err := strconv.Atoi(idParam)
//	if err != nil {
//		w.WriteHeader(400)
//		w.Write([]byte("ID could not be converted to integer"))
//		return
//	}
//
//	// error checking
//	if id >= len(tasks) {
//		w.WriteHeader(404)
//		w.Write([]byte("No task found with specified ID"))
//		return
//	}
//
//	// get the current value
//	task := &tasks[id]
//	json.NewDecoder(r.Body).Decode(task)
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(task)
//
//	taskUpdates, _ := json.Marshal(tasks)
//
//	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)
//	log.Println("patch: ", task)
//}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// get the ID of the task from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(tasks) {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		return
	}

	tasks = append(tasks[:id], tasks[id+1:]...)
	taskUpdates, _ := json.Marshal(tasks)

	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)

	w.WriteHeader(200)
	log.Println("delete: ", id)
}
