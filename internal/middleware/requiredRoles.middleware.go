package middleware

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func RequireRoles(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := FromContext(r.Context())
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "No hay claims en el contexto")
				return
			}

			err := utils.ValidateRol(allowedRoles, jwt.MapClaims(claims))
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, err.Error())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
