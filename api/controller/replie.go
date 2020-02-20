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
	var replie model.Replie

	json.NewDecoder(r.Body).Decode(&replie)

	replie.CreateReplie()

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if replie.ID != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(replie)
	}else{
		m := "Fail in replie creation"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	}
}

//Update replie
func UpdateReplie(w http.ResponseWriter, r *http.Request) {
	var replie model.Replie

	json.NewDecoder(r.Body).Decode(&replie)

	params := mux.Vars(r)

	replieId := params["id"]

	replie.ID, _ = strconv.Atoi(replieId)

	replie.UpdateReplie()

	w.Header().Set(ACCESS_CONTROL_ORIGIN, "*")
    w.Header().Set(ACCESS_CONTROL_HEADER, CONTENT_TYPE)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if replie.ID != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(replie)
	}else{
		m := "Fail in replie update"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	}
}

//Delete replie
func DeleteReplie(w http.ResponseWriter, r *http.Request) {
	//Missing implementation
}

