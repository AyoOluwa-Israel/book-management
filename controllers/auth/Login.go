package auth

import (
	"fmt"
	"net/http"

	"github.com/AyoOluwa-Israel/book-management/config"
	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/asaskevich/govalidator.v9"
)

func Login(c *fiber.Ctx) error {
	var user *models.UserLogin
	err := c.BodyParser(&user)

	if user.Email == "" || user.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Please provide all fields",
		})
	}

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "Error parsing data",
		})
	}

	if !govalidator.IsEmail(user.Email) {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Email format is wrong!",
		})
	}

	user.Email = utils.ConvertEmail(user.Email)
	fmt.Printf(user.Email, "\tEmail is here")

	var newUser models.User
	exists := db.Database.Db.Where("email = ?", user.Email).First(&newUser)
	fmt.Print(exists)
	if exists.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Email does not exists",
		})
	}

	if err := utils.VerifyPassword(newUser.Password, user.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Password!",
		})
	}

	config, _ := config.LoadConfig(".")

	token, err := utils.GenerateToken(config.TokenExpiresIn, newUser.Id, config.TokenSecret)

	res := models.LoginResponse{
		Status:  http.StatusCreated,
		Message: newUser.FullName,
		Token:   token,
	}

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "Couldn't generate token",
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}
