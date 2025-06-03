package middleware

import (
	"context"
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/config"
	"github.com/ProtoSG/app-salud-back/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"message": "No token, autorización deniegada",
			})
			return
		}

		tokenString := cookie.Value
		secret := config.NewConfig().TOKEN_SECRET

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"message": "Token no es válido",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"message": "Token claims no estan el formato esperado.",
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value("user").(jwt.MapClaims)
	return claims, ok
}
