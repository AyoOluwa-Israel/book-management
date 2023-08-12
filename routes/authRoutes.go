package routes

import (
	"github.com/AyoOluwa-Israel/book-management/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	app.Post("/login/", auth.Login)
	app.Post("/register/", auth.Register)
}
