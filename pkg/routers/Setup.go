package routers

import (
	"diino/pkg/controllers/users"

	"github.com/gofiber/fiber/v2"
)

//Setup roteia as funçoes
func Setup(app *fiber.App) {

	app.Post("/api/register", users.RegisterFunc)

	app.Post("/api/login", users.LoginFunc)

	app.Get("/api/user", users.UserFunc)

	app.Get("/api/logout", users.LogoutFunc)
}
