package controller

import (
	"github.com/Sifouo/Blog-BackEnd/database"
	"github.com/Sifouo/Blog-BackEnd/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := database.DB.Create(&blog).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Post created successfully",
	})
}



func AllPost