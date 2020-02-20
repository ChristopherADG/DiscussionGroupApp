package controller

import (
	"net/http"
	"app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

//Get all replies of a thread
func GetRepliesByThread(w http.ResponseWriter, r *http.Request) {
	var replie model.Replie

	params := mux.Vars(r)

	threadId := params["thread_id"]

	threadIdInt, _ := strconv.Atoi(threadId)

	var output []*model.Replie
	output = replie.GetAllReplies(threadIdInt)

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

//Get single thread
func GetReplie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	replieId := params["id"]

	var output model.Replie
	output = output.GetReplie(replieId)

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if output.ID == 0 {
		m := "No replie found"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	}
}

//Create replie
func CreateReplie(w http.ResponseWriter, r *http.Request) {

	parentId := r.FormValue("threadParent_id")
	content := r.FormValue("content")

	if content != "" {
		var replie model.Replie

		replie.ThreadParent_id,_ = strconv.Atoi(parentId)

		replie.Content = content
	
		w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    	w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

		replie.CreateReplie()

		if replie.ID != 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(replie)
			return
		}
	}
	m := "Fail in replie create"
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(m)
}

//Update replie
func UpdateReplie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	replieId := params["id"]

	parentId := r.FormValue("threadParent_id")
	content := r.FormValue("content")

	if content != "" {
		var replie model.Replie

		replie.ID, _ = strconv.Atoi(replieId)

		replie.ThreadParent_id,_ = strconv.Atoi(parentId)

		replie.Content = content
	
		w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    	w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

		replie.UpdateReplie()

		if replie.ID != 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(replie)
			return
		}
	}
	m := "Fail in replie update"
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(m)
}

//Delete replie
func DeleteReplie(w http.ResponseWriter, r *http.Request) {
	//Missing implementation
}

