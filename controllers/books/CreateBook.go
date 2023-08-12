package books

import (
	"fmt"
	"net/http"

	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateBook(c *fiber.Ctx) error {
	var book models.Books

	err := c.BodyParser(&book)

	user := c.Locals("user").(models.User)

	fmt.Print(user.Id, "user")

	if book.Name == "" || book.Isbn == "" || book.Edition <= 0 {
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

	var newBook models.Books
	db.Database.Db.Where("name = ? AND edition = ? ", book.Name, book.Edition).First(&newBook)

	// fmt.Print(exists, "EXISTS", newBook)
	if book.Edition == newBook.Edition {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.Response{
				Status:  fiber.StatusBadRequest,
				Message: "This book edition already exists",
			},
		)
	} else {
		book.Id = uuid.New()

		book.UserId = user.Id

		db.Database.Db.Create(&book)
	}

	return c.Status(http.StatusCreated).JSON(utils.Response{
		Status:  fiber.StatusCreated,
		Message: "Book Created Successfully!",
	})

}
