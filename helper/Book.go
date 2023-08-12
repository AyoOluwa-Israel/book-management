package helper

import "github.com/google/uuid"

type Book struct{
	Id        uuid.UUID  `json:"id"`
	Name      string    `json:"name"`
	NoOfPages int       `json:"noOfPages"`
}