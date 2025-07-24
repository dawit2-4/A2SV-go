// Package domain defines the core business entities for the Task Manager API.
package Domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status represents the status of a task.
type Status string

const (
	Pending   Status = "pending"
	Completed Status = "completed"
	NotDone   Status = "not-done"
)

//IsValid checks if Status value is valid
func (s Status) IsValid() bool {
	return s == Pending || s == Completed || s == NotDone
}

// Task represents a task entity.
type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string            `json:"title" bson:"title"`
	Description string            `json:"description" bson:"description"`
	DueDate     time.Time         `json:"due_date" bson:"due_date"`
	Status      Status            `json:"status" bson:"status"`
}

// Validate validates the Task data.
func (t Task) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(t.Title) > 100 {
		return errors.New("title cannot exceed 100 characters")
	}
	if len(t.Description) > 1000 {
		return errors.New("description cannot exceed 1000 characters")
	}
	if t.DueDate.Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	if !t.Status.IsValid() {
		return fmt.Errorf("invalid status: %s", t.Status)
	}
	return nil
}

// UserRole represents a user's role.
type UserRole string

const (
	RoleAdmin UserRole = "Admin"
	RoleUser  UserRole = "User"
)
//IsValid checks if a UserRole value is valid

func (r UserRole) IsValid() bool {
	return r == RoleAdmin || r == RoleUser
}

// User represents a user entity.
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string            `json:"username" bson:"username"`
	Password string            `json:"password" bson:"password"`
	Role     UserRole          `json:"role" bson:"role"`
}

// Validate validates the User data.
func (u User) Validate() error {
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

// TaskRepository defines task data access methods.
type TaskRepository interface {
	CreateTask(ctx context.Context, task Task) (Task, error)
	GetTaskByID(ctx context.Context, id string) (Task, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
	UpdateTask(ctx context.Context, id string, task Task) (Task, error)
	DeleteTask(ctx context.Context, id string) error
}

// UserRepository defines user data access methods.
type UserRepository interface{
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}