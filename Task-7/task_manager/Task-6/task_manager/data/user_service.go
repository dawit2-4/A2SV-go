package data

import (
	"context"
	"errors"
	"fmt"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
	JwtSecret  string
}

var (
	ErrInvalidCredential = fmt.Errorf("invalid username or password")
	ErrNotFound          = fmt.Errorf("user not found")
	ErrUsernameTaken     = fmt.Errorf("username already taken")
)

type UserManager interface {
	RegisterUser(c context.Context, user models.User) (models.User, error)
	LogIn(c context.Context, username, password string) (string, error)
	GetUserByUsername(c context.Context, username string) (models.User, error)
}

func NewUserService(client *mongo.Client, secret string) UserManager {
	return &UserService{
		collection: client.Database("tasks").Collection("users"),
		JwtSecret:  secret,
	}
}

func (us *UserService) GetUserByUsername(c context.Context, username string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := us.collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, ErrNotFound
		}
		return models.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

func (us *UserService) RegisterUser(c context.Context, user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return models.User{}, fmt.Errorf("invalid user data: %w", err)
	}

	_, err := us.GetUserByUsername(c, user.Username)
	if err == nil {
		return models.User{}, ErrUsernameTaken
	}
	if !errors.Is(err, ErrNotFound) {
		return models.User{}, fmt.Errorf("failed to check existing user: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	_, err = us.collection.InsertOne(c, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	user.Password = ""
	return user, nil
}

func (us *UserService) LogIn(c context.Context, username, password string) (string, error) {
	user, err := us.GetUserByUsername(c, username)
	if err != nil {
		return "", ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredential
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(us.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}