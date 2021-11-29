package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_addTaskHandler(t *testing.T) {
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
	//expected, _ := json.Marshal(tasks)
	expected := `[{"id":1,"title":"write some code","dueDate":"18/11/2021","status":"done"},{"id":2,"title":"write some code 2","dueDate":"29/11/2021","status":"todo"},{"id":3,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":4,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":5,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":6,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":7,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":8,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":9,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":10,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":11,"title":"content","dueDate":"17/11/2021","status":"done"}]`
	//println("sssssssssssssssssssssssssssss",string(expected))
	//if err != nil {
	//	return
	//}
	println("rr.Body.String()", rr.Body.String())
	println("string(expected)", string(expected))
	println(rr.Body.String() == string(expected))
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}

func Test_deleteTaskHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_retrieveTaskHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_updateTaskHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
