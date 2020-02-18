package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"../app/controller"
)

func main() {
	//init Router
	r := mux.NewRouter()

	//Router handlers
	r.HandleFunc("/api/threads", controller.GetThreads).Methods("GET")
	r.HandleFunc("/api/threads/{id}", controller.GetThread).Methods("GET")
	r.HandleFunc("/api/threads", controller.CreateThread).Methods("POST")
	r.HandleFunc("/api/threads/{id}", controller.UpdateThread).Methods("PUT")
	r.HandleFunc("/api/threads/{id}", controller.DeleteThread).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Starting up on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}