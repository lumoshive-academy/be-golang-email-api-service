package middleware

import (
	"net/http"

	"github.com/lumoshive-academy/be-golang-email-api-service/utils"
)

// Middleware for check API Key
func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		isValid := false

		for _, k := range utils.ApiKeyList {
			if k == apiKey {
				isValid = true
				break
			}
		}

		if !isValid {
			utils.ResponseBadRequest(w, http.StatusUnauthorized, "Unauthorized access, invalid API key", nil)
			return
		}

		next.ServeHTTP(w, r)

	})

}
