package routes

import (
	"github.com/Sifouo/Blog-BackEnd/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controller.Register)

}
