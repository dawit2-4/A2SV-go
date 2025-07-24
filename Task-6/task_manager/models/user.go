package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	RoleAdmin UserRole = "Admin"
	RoleUser  UserRole = "User"
)

func (role UserRole) IsValid() bool {
	switch role {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Role     UserRole           `json:"role" bson:"role"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if len(u.Username) > 50 {
		return errors.New("username cannot exceed 50 characters")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if !u.Role.IsValid() {
		return fmt.Errorf("invalid role: %s", u.Role)
	}
	return nil
}