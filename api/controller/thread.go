package controller

import (
	"net/http"
	"fmt"
	"app/model"
	"encoding/json"
	"github.com/gorilla/mux"
)
const CONTENT_TYPE  = "content-type"
const APPLICATION_JSON = "application/json"

//Get all threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	var thread model.Thread

	var output []*model.Thread
	output = thread.GetAllThreads()

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

//Get all threads
func GetThread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	threadId := params["id"]

	var output model.Thread
	output = output.GetThread(threadId)

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if output.ID == 0 {
		m := "No thread found"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	}
}

//Get all threads
func CreateThread(w http.ResponseWriter, r *http.Request) {
	var thread model.Thread

	json.NewDecoder(r.Body).Decode(&thread)

	thread.CreateThread()

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if thread.ID != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(thread)
	}else{
		m := "Fail in thread creation"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	}
}

//Get all threads
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>") //Missing implementation
}

//Get all threads
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>") //Missing implementation
}

