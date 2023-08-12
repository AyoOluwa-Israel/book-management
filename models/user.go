package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type User struct {
	gorm.Model           // Adds some metadata fields to the table
	Id         uuid.UUID `gorm:"type:uuid primaryKey"` 
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	Verified  bool      `json:"isVerified" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SignUpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" `
}


type NewUser struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Verified bool   `json:"isVerified" `
}