package routes

import (
	"github.com/AyoOluwa-Israel/book-management/controllers"
	"github.com/AyoOluwa-Israel/book-management/controllers/user"
	"github.com/AyoOluwa-Israel/book-management/middleWares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	app.Get("/users/", controllers.GetAllUsers)
	app.Get("/user/books/", middleWares.IsAuthenticated, user.GetAllBooksOfUser)

}
