package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Task Entity
type Task struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	DueDate string `json:"dueDate"`
	Status  string `json:"status"`
}

var tasks []Task = []Task{}

func main() {
	initLogFile()

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

func initLogFile() {
	file := "./" + "log" + ".out"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func retrieveTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var idParam string = mux.Vars(r)["id"]

	//If ${id} is null, then return tasks List
	if idParam == "" {
		var isExp = r.URL.Query().Get("expiredToday")
		//return exp lists
		if "--expiring-today" == isExp {
			var expTasks []Task = []Task{}
			for _, t := range tasks {
				if time.Now().Format("02/01/2006") == t.DueDate {
					expTasks = append(expTasks, t)
				}
			}
			json.NewEncoder(w).Encode(expTasks)
			log.Println("getExpTasks: ", isExp, expTasks)
		} else {
			//retrun all list
			json.NewEncoder(w).Encode(tasks)
			log.Println("getAllTasks: ", tasks)
		}
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
	//w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
	log.Println("getTask", task)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)
	newTask.Id = tasks[len(tasks)-1].Id + 1
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

	for i, t := range tasks {
		if t.Id == updatedTask.Id {
			tasks[i] = updatedTask
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)

	taskUpdates, _ := json.Marshal(tasks)

	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)
	log.Println("update: ", updatedTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// get the ID of the task from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	var delIndex = 0
	var hasId = 0
	for i, t := range tasks {
		if t.Id == id {
			delIndex = i
			hasId = 1
		}
	}

	//can not find task for input id
	if hasId == 0 {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		log.Println("No task found with specified ID")
		return
	}

	tasks = append(tasks[:delIndex], tasks[delIndex+1:]...)

	taskUpdates, _ := json.Marshal(tasks)

	_ = ioutil.WriteFile("tasks.json", taskUpdates, 0644)

	w.WriteHeader(200)
	log.Println("delete: ", id)
}
