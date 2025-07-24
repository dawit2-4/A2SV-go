package models

import (
	"errors"
	"fmt"
	"time"
)

// Status represents the task status
type Status string

// Define status constants
const (
	Pending   Status = "pending"
	Completed Status = "completed"
	NotDone   Status = "not-done"
)

// IsValid checks if a Status value is valid
func (s Status) IsValid() bool {
	return s == Pending || s == Completed || s == NotDone
}

// String implements the Stringer interface for Status
func (s Status) String() string {
	return string(s)
}

// Task represents a task entity
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
}

// NewTask creates a new Task with validation
func NewTask(id, title, description string, dueDate time.Time, status Status) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if !status.IsValid() {
		return nil, fmt.Errorf("invalid status: %s", status)
	}
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}, nil
}