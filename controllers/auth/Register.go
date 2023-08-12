package auth

import (
	"net/http"
	"strings"

	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/asaskevich/govalidator.v9"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)

	if user.FullName == "" || user.Email == "" || user.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Please provide all fields",
		})
	}

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.ResponsePayload{
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

	exists := db.Database.Db.Where("email = ?", user.Email).First(&user)

	if exists.RowsAffected > 0 {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Email already exists",
		})
	} else {
		if strings.TrimSpace(user.Password) == "" {
			return c.Status(http.StatusBadRequest).JSON(utils.Response{
				Status:  http.StatusBadRequest,
				Message: "Kindly enter a password",
			})
		}

		user.Password, err = utils.EncryptPassword(user.Password)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		user.Id = uuid.New()
		db.Database.Db.Create(&user)
	}

	res := models.SignUpResponse{
		Status:  http.StatusCreated,
		Message: "User Created Successfully",
	}

	return c.Status(http.StatusCreated).JSON(res)

}
