package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Get All Tasks
func Test_retrieveTaskHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(retrieveTaskHandler)
	tasks = initJsonFile(tasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected, _ := json.Marshal(tasks)

	if strings.Replace(rr.Body.String(), "\n", "", -1) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}

// Get Tasks Expired List
func Test_retrieveExpiredTaskHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)

	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("expiredToday", "--expiring-today")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(retrieveTaskHandler)
	tasks = initJsonFile(tasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":2,"title":"write some code 2","dueDate":"29/11/2021","status":"todo"}`

	if strings.Replace(rr.Body.String(), "\n", "", -1) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}

func Test_addTaskHandler(t *testing.T) {

}

func Test_deleteTaskHandler(t *testing.T) {
	//id := 11

	url := "http://localhost:8080/tasks/1"

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "d02d64ae-fb37-8604-5170-b1ddbf870a64")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func Test_updateTaskHandler(t *testing.T) {

}
