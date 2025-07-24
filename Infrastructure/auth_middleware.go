// Package Infrastructure provides external services for the Task Manager API.
package Infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware (jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return 
		}

		claims, err := jwtService.ValidateToken(parts[1])
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return 
		}

		c.Set("userID", claims["id"].(string))
		c.Set("role", claims["role"])
		c.Next()
	}
}

//AdminOnlyMiddleware restricts access to admin users.
func AdminOnlyMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found in token"})
			c.Abort()
			return 
		}

		if role != "Admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "admin role required"})
			c.Abort()
			return 
		}

		c.Next()
	}
}