package model

import (
)

//Tread Struct (Model) 

type Thread struct {
	ID            int         `json:"thread_id"`
	Title          string      `json:"name"`
	Description   string      `json:"description"`
} 
