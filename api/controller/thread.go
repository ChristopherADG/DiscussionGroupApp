package controller

import (
	"net/http"
	"app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)
const CONTENT_TYPE  = "content-type"
const APPLICATION_JSON = "application/json"
const ACCESS_CONTROL_ORIGIN  = "Access-Control-Allow-Origin"
const ACCESS_CONTROL_HEADER = "Access-Control-Allow-Headers"

//Get all threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	var thread model.Thread

	var output []*model.Thread
	output = thread.GetAllThreads()

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

//Get single thread
func GetThread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	threadId := params["id"]

	var output model.Thread
	output = output.GetThread(threadId)

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
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

//Create thread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	desc := r.FormValue("description")

	if title != "" {
		var thread model.Thread

		thread.Title = title

		thread.Description = desc
	
		w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    	w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

		thread.CreateThread()

		if thread.ID != 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(thread)
			return
		}
	}
	m := "Fail in thread creation"
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(m)
}

//Update thread
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	threadId := params["id"]

	title := r.FormValue("title")
	desc := r.FormValue("description")	

	if title != "" {
		var thread model.Thread

		thread.Title = title

		thread.Description = desc

		thread.ID, _ = strconv.Atoi(threadId)

		thread.UpdateThread()

		w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
		w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

		if thread.ID != 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(thread)
			return
		}
	}
	m := "Fail in thread update"
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(m)
}

//Delete thread
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	//Missing implementation
}

