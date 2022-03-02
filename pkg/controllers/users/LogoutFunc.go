package users

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

//LogoutFunc Função de logout
func LogoutFunc(c *fiber.Ctx) (err error) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout Success",
	})

}
