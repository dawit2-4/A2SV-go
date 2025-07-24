package data

import (
	"errors"
	"fmt"
	"sync"
	"task_manager/models"
)

// TaskService defines the interface for task data operations
type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
}

type InMemoryTaskService struct {
	tasks map[string]models.Task
	mu sync.RWMutex  // For thread-safe access
}

// NewTaskService creates a new InMemoryTaskService
func NewTaskService() TaskService {
	return &InMemoryTaskService{
		tasks: make(map[string]models.Task),
	}
}

func (s *InMemoryTaskService) CreateTask(task models.Task) (models.Task, error) {
	if task.ID == "" {
		return models.Task{}, errors.New("ID can not be empty")
	}
	if !task.Status.IsValid() {
		return models.Task{}, errors.New("Not a valid task status")
	}
	if task.Title == "" {
		return models.Task{}, errors.New("Title can not be empty")
	}


	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.tasks[task.ID]; exists {
		return models.Task{}, fmt.Errorf("Task with ID %s already exists.", task.ID)
	}

	s.tasks[task.ID] = task
	return task, nil

}

func (s *InMemoryTaskService) GetTaskByID(id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("ID can not be empty")
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id] 
	if !exists {
		return  models.Task{}, fmt.Errorf("No task with ID %s exists", id)
	}

	return task, nil
}

func (s *InMemoryTaskService) GetAllTasks() ([]models.Task, error) {
	s.mu.Unlock()
	defer s.mu.Unlock()

	tasks := make([]models.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *InMemoryTaskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("ID cannot be empty.")
	}

	if task.Title == "" {
		return models.Task{}, errors.New("Title cannot be empty.")
	}

	if !task.Status.IsValid() {
		return models.Task{}, fmt.Errorf("%v is not a valid status.", task.Status)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, fmt.Errorf("task with ID %s does not exist", id)
	}
	task.ID = id
	s.tasks[id] = task
	return  task, nil
}

func (s *InMemoryTaskService) DeleteTask(id string) error {
	if id == "" {
		return errors.New("ID can not be empty.")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("task with ID %s does not exist", id)
	}

	delete(s.tasks, id)
	return nil
}
