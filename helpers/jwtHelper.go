package helpers

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	configs "github.com/tarunrana0222/user_project_go/config"
)

type SignedDetails struct {
	Client string `json:"client"`
	jwt.StandardClaims
}

func GenerateJwtToken(clientId string) (string, error) {
	signed := &SignedDetails{
		Client: clientId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, signed).SignedString([]byte(configs.Jwt_Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(xToken string) (string, error) {
	token, err := jwt.ParseWithClaims(xToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Jwt_Secret), nil
	})

	if claims, ok := token.Claims.(*SignedDetails); ok && token.Valid {
		clientExists, err := ClientExists(claims.Client)
		if err != nil {
			return "", err
		}
		return clientExists.ClientID, nil
	} else {
		return "", err
	}
}
