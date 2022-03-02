package users

import (
	"diino/pkg/database"
	"diino/pkg/models"
	"diino/pkg/oauth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//UserFunc Utilizado para permanecer logado cliente
func UserFunc(c *fiber.Ctx) (err error) {

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(oauth.SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated.",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}
