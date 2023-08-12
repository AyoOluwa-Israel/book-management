package main

import (
	"log"
	"net/http"

	"github.com/AyoOluwa-Israel/book-management/config"
	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db.NewConnection(&config)
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Welcome to Pedigree API",
	})

	app.Use(logger.New())

	app.Use(cors.New(
		cors.Config{
			AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		}))

	router := app.Group("/v1/api")

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.BookRoutes(router)

	app.Get("/v1/api", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Welcome to Pedigree Api version 1",
			"status":  http.StatusOK,
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Welcome to Pedigree Api",
			"status":  http.StatusOK,
		})
	})

	log.Fatal(app.Listen(":" + config.ServerPort))
}
