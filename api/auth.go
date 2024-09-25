package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("TEST_SECRET_KEY_FOR_JWT")

// claims represents the JWT claims
type claims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // For expiration time
}

// GenerateJWT generates a new JWT token with the given username.
// The token has an expiration time of 24 hours.
func generateJWT(username string) (string, error) {

	// Create the JWT claims, which include the username and expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Use Hash and Sign paradigm to create token with the claims using HMAC SHA256 (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}
	return tokenString, nil
}

// ValidateJWT validates a JWT token
func verifyJWT(tokenString string) (*claims, error) {
	claims := &claims{}

	// Inner function to return secret key used to sign the JWT token.
	// ParseWithClaims uses a function that can return the key dynamically for more complex use cases.
	// In this case, I am using a fixed key, so I return it directly with the keyFunc.
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// ExtractToken extracts the JWT token from the Authorization header and removes the "Bearer " prefix.
func extractToken(request *http.Request) (string, error) {

	// Get the Authorization header
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization token required")
	}

	// Split the token from Bearer in format: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return parts[1], nil
}

// Middleware that checks if the request is authorized by verifying the JWT token in the Authorization header.
// All calculations endpoints require this token authorization.
func (api *API) authMiddleware(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Extract token from the request header and remove the "Bearer " prefix
		token, err := extractToken(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validate JWT
		claims, err := verifyJWT(token)
		if err != nil {
			http.Error(writer, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid. Set user information in request context.
		// The type is used to avoid key collisions in the context (for instance, if another middleware uses the same key).
		type usernameKeyType struct{}
		usernameKey := usernameKeyType{}
		ctx := context.WithValue(request.Context(), usernameKey, claims.Username)
		request = request.WithContext(ctx)

		// Call the next handler, which is now authorized
		nextHandler.ServeHTTP(writer, request)
	}
}
