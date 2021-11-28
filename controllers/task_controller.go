package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rest-demo-go/entity"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//GetAllTask get all task data
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Task
	//database.Connector.Find(&persons)

	file, e := os.Create("./temp.txt")
	if e != nil {
		fmt.Println("创建文件失败！", e)
	}
	encoder := json.NewEncoder(file)
	encode := encoder.Encode(wek)
	if encode != nil {
		fmt.Println("wek写入文件失败！")
	}
	file.Close()
	time.Sleep(2 * time.Second)
	open, i := os.Open("./temp.txt")
	if i != nil {
		fmt.Println("文件打开失败！")
	}
	defer open.Close()

	decoder := json.NewDecoder(open)
	decode := decoder.Decode(&wek2)
	if decode != nil {
		fmt.Println("文件反序列化失败！")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

//GetTaskByID returns task with specific ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var task entity.Task
	//database.Connector.First(&task, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

//CreateTask creates task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var task entity.Task
	json.Unmarshal(requestBody, &task)

	//database.Connector.Create(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

//UpdateTaskByID updates task with respective ID
func UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var task entity.Task
	json.Unmarshal(requestBody, &task)
	//database.Connector.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

//DeletTaskByID delete's task with specific ID
func DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var task entity.Task
	id, _ := strconv.ParseInt(key, 10, 64)
	//database.Connector.Where("id = ?", id).Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
