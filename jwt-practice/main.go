package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type User struct{
	ID       uint    `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
}


var users = make(map[string]*User)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error":"Authorization header required"})
			c.Abort()
			return 
		}

		authParts := strings.Split(authHeader, " ")
		if len(authHeader) != 2 || strings.ToLower(authParts[0]) != "bearer"{
			c.JSON(401, gin.H{"error":"Invalid authorization error"})
			c.Abort()
			return 
		}
		c.Next()
	}
}

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"Welcome to jwt auth GO."})
	})

	router.POST("/register", func(c *gin.Context){
		var user User

		err := c.ShouldBindJSON(&user)
		if err != nil{
			c.JSON(400, gin.H{"message":"Invalid json body"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil{
			c.JSON(500, gin.H{"error":"Internal server error"})
			return
		}

		user.Password = string(hashedPassword)
		users[user.Email] = &user

		c.JSON(201, gin.H{"message":"user registered successfully"})
	})

	router.POST("/login", func(c *gin.Context){
		var user User

		if err := c.ShouldBindJSON(&user); err != nil{
			c.JSON(400, gin.H{"error":"Invalid request payload"})
			return
		}

		var jwtSecret = []byte(my_secrete_key)
		existingUser, ok := users[user.Email]
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(user.Password)) != nil {
			c.JSON(401, gin.H{"error":"Invalid email or password"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": existingUser.ID,
			"email": existingUser.Email,
		})

		jwtToken ,err := token.SignedString(jwtSecret)
		if err != nil{
			c.JSON(500, gin.H{"error":"Internal server error"})
			return
		}

		c.JSON(200, gin.H{"message":"user registered successfully"})
	})

	router.Run()
}