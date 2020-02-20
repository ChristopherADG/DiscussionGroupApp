package model

import (
	"app/config"
	"log"
)

//Tread Struct (Model) 

type Thread struct {
	ID            int         `json:"thread_id"`
	Title          string      `json:"title"`
	Description   string      `json:"description"`
} 


func (thread *Thread) GetAllThreads() ([]*Thread) {
	var m []*Thread

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return m
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM threads")
	if err != nil {
		log.Println(err)
		return m
	}

	for result.Next() {
		var thread Thread

		scanErr := result.Scan(&thread.ID, &thread.Title, &thread.Description)
		if scanErr != nil {
			return m
		}

		m = append(m, &thread)
	}

	return m
}

func (thread *Thread) GetThread(id string) (Thread) {
	var m Thread

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return m
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM threads WHERE thread_id = ?", id)
	if err != nil {
		log.Println(err)
		return m
	}

	for result.Next() {
		var thread Thread

		scanErr := result.Scan(&thread.ID, &thread.Title, &thread.Description)
		if scanErr != nil {
			return m
		}
		m = thread
	}

	return m
}

func (thread *Thread) CreateThread() {

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
	
	result, err := db.Exec("INSERT INTO threads(title, description) VALUES (?,?)", thread.Title, thread.Description)
    if err != nil {
		log.Println(err)
		return
    } else {
        id, err := result.LastInsertId()
        if err != nil {
            log.Println(err)
			return
        } else {
            thread.ID = int(id)
        }
    }
}

func (thread *Thread) UpdateThread() {

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
	_, execErr := db.Exec("UPDATE threads SET title = ?, description = ? WHERE thread_id = ?", thread.Title, thread.Description, thread.ID)
    if execErr != nil {
		log.Println(execErr)
    }
}