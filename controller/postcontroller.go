package controller

import (
	"math"
	"strconv"

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

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var blogs []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&blogs)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": blogs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}
