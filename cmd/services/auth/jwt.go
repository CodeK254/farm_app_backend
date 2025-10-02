package auth

import (
	// "log"
	"strconv"
	"time"

	"github.com/CodeK254/farm_app_backend/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}