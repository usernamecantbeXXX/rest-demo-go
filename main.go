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

	tasks = initJsonFile(tasks)

	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initJsonFile(tasks []Task) []Task {
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened tasks.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal([]byte(byteValue), &tasks)
	return tasks
}

func initHandlers(router *mux.Router) {
	router.HandleFunc("/tasks/{id}", retrieveTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", retrieveTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", addTaskHandler).Methods("POST")
	router.HandleFunc("/tasks", updateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")
}

// Init JSON File as database
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

//Restful API GET，Retrieve one / all / expiring tasks
func retrieveTaskHandler(w http.ResponseWriter, r *http.Request) {

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
		log.Println("ID could not be converted to integer")
		return
	}

	// error checking
	if id >= len(tasks) {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		log.Println("No task found with specified ID")
		return
	}

	task := tasks[id]
	json.NewEncoder(w).Encode(task)
	log.Println("getTask", task)
}

//Restful API POST，create one task
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
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

//Restful API PUT，update one task
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// get the ID of the task from JSON body
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)
	id := updatedTask.Id
	if 0 >= id {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		log.Println("ID could not be converted to integer")
		return
	}

	if id >= len(tasks) {
		w.WriteHeader(404)
		w.Write([]byte("No task found with specified ID"))
		log.Println("No task found with specified ID")
		return
	}

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

//Restful API DELETE，delete one task
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// get the ID of the task from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		log.Println("ID could not be converted to integer")
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
