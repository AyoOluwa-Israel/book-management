package middleWares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AyoOluwa-Israel/book-management/config"
	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	var tokenString string

	authHeader := c.Get("Authorization")

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  fiber.StatusUnauthorized,
			Message: "You are not logged in",
		})
	}

	config, _ := config.LoadConfig(".")

	userId, err := utils.ValidateToken(tokenString, config.TokenSecret)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  fiber.StatusUnauthorized,
			Message: fmt.Sprintf("Invalid token: %v", err),
		})
	}

	var user models.User
	// db.Database.Db.First(&user, "id = ?", userId)

	exists := db.Database.Db.Where("id = ?", userId).First(&user)

	// fmt.Print(exists)
	// fmt.Print(user.FullName, user.Id, "=======================", userId)
	if exists.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{
			Status:  http.StatusBadRequest,
			Message: "User does not exists",
		})
	}

	c.Locals("user", user)
	return c.Next()

}
