package routes

import (
	"github.com/AyoOluwa-Israel/book-management/controllers/books"
	"github.com/AyoOluwa-Israel/book-management/middleWares"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app fiber.Router) {
	app.Post("/book/", middleWares.IsAuthenticated, books.CreateBook)
}
