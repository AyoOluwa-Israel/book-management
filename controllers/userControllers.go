package controllers

import (
	"net/http"

	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
	Verified bool      `json:"isVerified" `
}

func CreateUserResponse(user models.User) User {
	return User{
		Id:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
		Verified: user.Verified,
	}
}

func GetAllUsers(c *fiber.Ctx) error {
	db := db.Database.Db
	var users []User

	// Use Select to fetch only the required fields from the database
	db.Select("id, full_name, email").Find(&users)

	return c.Status(http.StatusOK).JSON(utils.ResponsePayload{
		Status:  http.StatusOK,
		Message: "Users Retrieved Successfully!!",
		Payload: users,
	})

}
