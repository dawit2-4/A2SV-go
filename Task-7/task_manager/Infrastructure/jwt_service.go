// Package Infrastructure provides external services for the Task Manager API.
package Infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService defines methods fro JWT operations
type JWTService interface {
	GenerateToken(id, username, role string) (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

//jwtService implements JWTService
type jwtService struct {
	secret string
}

// GenerateToken implements JWTService.
func (j *jwtService) GenerateToken(id string, username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"username": username,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(j.secret))
}

// ValidateToken implements JWTService.
func (j *jwtService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing methos")
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// NewJWTService creates a new JWTService
func NewJWTService(secret string) JWTService {
	return &jwtService{secret: secret}
}
