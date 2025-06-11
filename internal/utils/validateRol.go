package utils

import (
	"fmt"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateRol(roles []string, claims jwt.MapClaims) error {
	roleName := claims["role_name"].(string)

	if slices.Contains(roles, roleName) {
		return nil
	}

	return fmt.Errorf("no estás autorizado")
}
