package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = randStringRunes(10) + file.Filename
		err = c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + filename,
	})
}
