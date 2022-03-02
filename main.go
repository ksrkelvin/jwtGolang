package main

import (
	"diino/pkg/database"
	"diino/pkg/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routers.Setup(app)

	app.Listen(":5000")
}
