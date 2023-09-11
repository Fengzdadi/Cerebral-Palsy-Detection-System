package Algorithm

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JWT_SECRET = "YourSecretKey"

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return t, nil
}
