package model

import (
	"app/config"
	"log"
)

//Replie Struct (Model) 

type Replie struct {
	ID            	 int         `json:"replie_id"`
	ThreadParent_id  int      `json:"threadParent_id"`
	Content   	  	 string      `json:"content"`
}

func (replie *Replie) GetAllReplies(thread_id int) ([]*Replie) {
	var m []*Replie

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return m
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM replies WHERE threadParent_id = ?", thread_id)
	if err != nil {
		log.Println(err)
		return m
	}

	for result.Next() {
		var replie Replie

		scanErr := result.Scan(&replie.ID, &replie.ThreadParent_id, &replie.Content)
		if scanErr != nil {
			return m
		}

		m = append(m, &replie)
	}

	return m
}

func (replie *Replie) GetReplie(id string) (Replie) {
	var m Replie

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return m
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM replies WHERE replie_id = ?", id)
	if err != nil {
		log.Println(err)
		return m
	}

	for result.Next() {
		var replie Replie

		scanErr := result.Scan(&replie.ID, &replie.ThreadParent_id, &replie.Content)
		if scanErr != nil {
			return m
		}
		m = replie
	}

	return m
}

func (replie *Replie) CreateReplie() {

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
	
	result, err := db.Exec("INSERT INTO replies(threadParent_id, content) VALUES (?,?)", replie.ThreadParent_id, replie.Content)
    if err != nil {
		log.Println(err)
		return
    } else {
        id, err := result.LastInsertId()
        if err != nil {
            log.Println(err)
			return
        } else {
            replie.ID = int(id)
        }
    }
}

func (replie *Replie) UpdateReplie() {

	db, err := config.Open()

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
	_, execErr := db.Exec("UPDATE replies SET content = ? WHERE replie_id = ?", replie.Content, replie.ID)
    if execErr != nil {
		log.Println(execErr)
    }
}