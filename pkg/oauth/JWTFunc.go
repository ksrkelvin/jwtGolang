package oauth

import (
	"diino/pkg/models"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//SecretKey Oauth secret
var SecretKey = os.Getenv("OAUTHSECRETS")

//JWTFunc gerar o token jwt
func JWTFunc(user models.User) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1day
	})
	token, err = claims.SignedString([]byte(SecretKey))

	return token, err
}
