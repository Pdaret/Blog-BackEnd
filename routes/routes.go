package routes

import (
	"github.com/Sifouo/Blog-BackEnd/controller"
	"github.com/Sifouo/Blog-BackEnd/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/createpost", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
}
