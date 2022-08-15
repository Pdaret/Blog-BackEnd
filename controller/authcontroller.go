package controller

import (
	"strconv"
	"strings"
	"time"

	"github.com/Sifouo/Blog-BackEnd/database"
	"github.com/Sifouo/Blog-BackEnd/models"
	"github.com/Sifouo/Blog-BackEnd/utils"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.User

	//Check if the request is a json request
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	//Check password is less than 6 characters
	if len(data["password"].(string)) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password must be at least 6 characters",
		})
	}

	//Check if email address is valid
	if !models.IsEmail(data["email"].(string)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email address",
		})
	}

	//Check if phone number is valid
	if !models.IsPhone(data["phone"].(string)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid phone number",
		})
	}

	//Check if the email address is already in use
	if err := database.DB.Where("email = ?", data["email"].(string)).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email address already in use",
		})
	}

	//Check if the phone number is already in use
	if err := database.DB.Where("phone = ?", data["phone"].(string)).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone number already in use",
		})
	}

	//Create new user
	user.FirstName = data["first_name"].(string)
	user.LastName = data["last_name"].(string)
	user.Email = strings.TrimSpace(data["email"].(string))
	user.Phone = data["phone"].(string)
	user.HashPassword(data["password"].(string))

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "User created successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]interface{}
	var user models.User

	//Check if the request is a json request
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	//Check if the email address is already in use
	if err := database.DB.Where("email = ?", data["email"].(string)).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email address",
		})
	}
	//Check if the password is correct
	if !user.CompareHashAndPassword(data["password"].(string)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "User logged in successfully",
	})
}

type Claims struct {
	jwt.StandardClaims
}
