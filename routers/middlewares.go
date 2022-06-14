package routers

import (
	"BI.ZONE_test/handlers"
	"BI.ZONE_test/models"
	"context"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Request to endpoint %s. Method: %s\n", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

var authCheck = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/login", "/api/user/register"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = handlers.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = handlers.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.JWT{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			response = handlers.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}

		if !token.Valid {
			response = handlers.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}

		log.Infof("User %d with role %d", tk.UserID, tk.Role)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		ctx = context.WithValue(ctx, "role", tk.Role)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

var adminCheck = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role")
		if role != uint(models.ADMIN_ROLE) {
			response := handlers.Message(false, "You are not admin")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}
		next.ServeHTTP(w, r)
	}
}

var analystCheck = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role")
		if role != uint(models.ANALYST_ROLE) {
			response := handlers.Message(false, "You are not analyst")
			w.WriteHeader(http.StatusForbidden)
			handlers.Respond(w, response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
