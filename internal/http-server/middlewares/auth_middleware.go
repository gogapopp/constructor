package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"constructor/internal/config"
	"constructor/internal/lib/jwt"
)

type (
	CtxKeyUserRole string
	CtxKeyUserID   int
)

const (
	UserRoleKey CtxKeyUserRole = ""
	UserIDKey   CtxKeyUserID   = 0
)

// AuthMiddleware checks the user cookie for the JWT token.
func AuthMiddleware(next http.Handler, config *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Extract the remote IP address
		remoteIP := r.RemoteAddr

		// Split the IP from the port (if present)
		remoteHost, _, err := net.SplitHostPort(remoteIP)
		if err != nil {
			log.Printf("Failed to split remote IP: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		// Extract the main server IP from the config
		mainHost, _, err := net.SplitHostPort(config.HTTPConfig.Addr)
		if err != nil {
			log.Printf("Failed to split main server IP: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		// log.Println("remoteIP ", remoteHost, "mainIP ", mainHost)

		// Check if the request is from the main server and is a download request
		if remoteHost == mainHost && r.Header.Get("X-Download-Request") == "true" {
			next.ServeHTTP(w, r)
			return
		}

		// For other requests, check the JWT token
		cookie, err := r.Cookie("ssid")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Redirect(w, r, "/signin?redirected=true", http.StatusTemporaryRedirect)
				return
			default:
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
		}

		jwtToken := cookie.Value

		userID, role, err := jwt.ParseJWTToken(jwtToken)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserRoleKey, fmt.Sprint(role))
		ctx = context.WithValue(ctx, UserIDKey, fmt.Sprint(userID))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
