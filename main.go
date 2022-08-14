package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Sifouo/Blog-BackEnd/database"
	"github.com/Sifouo/Blog-BackEnd/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected successfully")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
