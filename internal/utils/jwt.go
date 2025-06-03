package utils

import (
	"fmt"
	"time"

	"github.com/ProtoSG/app-salud-back/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

func CreateJWT(userID int, roleName string) (string, error) {
	claims := Claims{
		UserID:   userID,
		RoleName: roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	config := config.NewConfig()
	tokenString, err := token.SignedString([]byte(config.TOKEN_SECRET))
	if err != nil {
		return "", fmt.Errorf("Error al crear token")
	}

	return tokenString, nil
}
