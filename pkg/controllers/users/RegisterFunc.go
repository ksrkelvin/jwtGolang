package users

import (
	"diino/pkg/database"
	"diino/pkg/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//RegisterFunc Funcao de registro de cliente.
func RegisterFunc(c *fiber.Ctx) (err error) {
	var data map[string]string

	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
