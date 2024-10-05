package jwt_token

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/online-store-app/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GenerateToken(userId int) (string, error) {

	secretKey := config.Env.SecretKeyJWT

	claim := jwt.MapClaims{}
	claim["id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	secretKey := config.Env.SecretKeyJWT
	tokens, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return tokens, nil
}
