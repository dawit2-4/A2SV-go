package Usecase

import (
	"context"
	"task_manger/Domain"
)

// TaskUsecase defines task-related business logic.
type TaskUsecase interface {
	CreateTask(ctx context.Context, task Domain.Task) (Domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (Domain.Task, error)
	GetAllTasks(ctx context.Context) ([]Domain.Task, error)
	UpdateTask(ctx context.Context, id string, task Domain.Task) (Domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

// taskUsecase implements TaskUsecase.
type taskUsecase struct {
	taskRepo Domain.TaskRepository
}

// CreateTask implements TaskUsecase.
func (t *taskUsecase) CreateTask(ctx context.Context, task Domain.Task) (Domain.Task, error) {
	if err := task.Validate(); err != nil {
		return Domain.Task{}, err
	}
	return t.taskRepo.CreateTask(ctx, task)
}

// DeleteTask implements TaskUsecase.
func (t *taskUsecase) DeleteTask(ctx context.Context, id string) error {
	return t.taskRepo.DeleteTask(ctx, id)
}

// GetAllTasks implements TaskUsecase.
func (t *taskUsecase) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	return t.taskRepo.GetAllTasks(ctx)
}

// GetTaskByID implements TaskUsecase.
func (t *taskUsecase) GetTaskByID(ctx context.Context, id string) (Domain.Task, error) {
	return t.taskRepo.GetTaskByID(ctx, id)
}

// UpdateTask implements TaskUsecase.
func (t *taskUsecase) UpdateTask(ctx context.Context, id string, task Domain.Task) (Domain.Task, error) {
	if err := task.Validate(); err != nil{
		return Domain.Task{}, err
	}

	return t.taskRepo.UpdateTask(ctx, id, task)
}

// NewTaskUsecase creates a new task with validation.
func NewTaskUsecase(taskRepo Domain.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}
