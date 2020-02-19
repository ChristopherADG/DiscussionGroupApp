package controller

import (
	"net/http"
	"fmt"
)
const CONTENT_TYPE  = "content-type"
const APPLICATION_JSON = "application/json"

//Get all threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

//Get all threads
func GetThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

//Get all threads
func CreateThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

//Get all threads
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

//Get all threads
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

