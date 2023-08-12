package user

import (
	"github.com/AyoOluwa-Israel/book-management/db"
	"github.com/AyoOluwa-Israel/book-management/helper"
	"github.com/AyoOluwa-Israel/book-management/models"
	"github.com/AyoOluwa-Israel/book-management/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooksOfUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	var books []helper.Book

	db.Database.Db.Select("id, name, no_of_pages").Where("user_id = ?", user.Id).Find(&books)

	return c.Status(fiber.StatusOK).JSON(utils.ResponsePayload{
		Status:  fiber.StatusOK,
		Message: "Books Fetched Successfully!",
		Payload: books,
	})

}
