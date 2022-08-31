package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
	"github.com/nabiel-syarif/playlist-api/pkg/response"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
)

func AuthOnly(jwt jwt.JwtHelper) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer", "", 1)
			if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer") || token == "" {
				utils.ResponseWithJson(w, http.StatusUnauthorized, response.StandardResponse{
					Status: "UNAUTHENTICATED",
					Error:  "Missing authorization token",
				})
				return
			}

			claim, verified := jwt.VerifyToken(strings.TrimSpace(token))
			if !verified {
				utils.ResponseWithJson(w, http.StatusUnauthorized, response.StandardResponse{
					Status: "FORBIDDEN",
					Error:  "Token not valid",
				})
				return
			}

			ctx := context.WithValue(r.Context(), "userId", claim["user_id"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
