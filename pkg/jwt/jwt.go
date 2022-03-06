package jwt

import (
	"fmt"
	"strings"
	"time"

	config "go-just-portfolio/pkg/config"

	"github.com/golang-jwt/jwt"
)

func GetFieldFromJWT(token string, field string) (string, error) {
	words := strings.Fields(token)
	claims, err := ParseJWT(words[1])

	if err != nil {
		return "", err
	}

	idString := fmt.Sprint(claims[field])
	return idString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	conf := config.GetConfig()
	hmacSampleSecret := []byte(conf.JWT_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func MakeJWT(Shortname string, Mail string, Id string) (string, error) {
	conf := config.GetConfig()
	hmacSampleSecret := []byte(conf.JWT_SECRET)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"shortname": Shortname,
		"mail":      Mail,
		"id":        Id,
		"nbf":       time.Date(2021, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
