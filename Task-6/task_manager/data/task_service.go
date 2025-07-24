package data

import (
	"context"
	"errors"
	"fmt"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService defines the interface for task data operations
type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
}

type MongoTaskService struct {
	collection *mongo.Collection
}

// NewTaskService creates a new MongoTaskService
func NewTaskService(collection *mongo.Collection) *MongoTaskService {
	return &MongoTaskService{collection: collection}
}

func (s *MongoTaskService) CreateTask(task models.Task) (models.Task, error) {
	if task.ID == "" {
		return models.Task{}, errors.New("ID can not be empty")
	}
	if !task.Status.IsValid() {
		return models.Task{}, errors.New("not a valid task status")
	}
	if task.Title == "" {
		return models.Task{}, errors.New("title can not be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//check if task already exists
	filter := bson.M{"id": task.ID}
	count, err := s.collection.CountDocuments(ctx, filter)

	if err != nil{
		return models.Task{}, err
	}

	if count > 0 {
		return models.Task{}, fmt.Errorf("task with ID %s already exists", task.ID)
	}

	_, err = s.collection.InsertOne(ctx, task)
	if err != nil{
		return models.Task{}, err
	}

	return task, nil
}

func (s *MongoTaskService) GetTaskByID(id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("ID can not be empty")
	}
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	var task models.Task
	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err != nil{
		if err == mongo.ErrNoDocuments {
			return  models.Task{}, fmt.Errorf("no task with ID %s exists", id)
		}
		return models.Task{}, err
	}

	return task, nil
}

func (s *MongoTaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil{
		return []models.Task{}, err
	}
	if err = cursor.All(ctx, &tasks); err != nil{
		return []models.Task{}, err
	}

	return tasks, nil
}

func (s *MongoTaskService) UpdateTask(id string, updated models.Task) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("ID cannot be empty")
	}

	if updated.Title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}

	if !updated.Status.IsValid() {
		return models.Task{}, fmt.Errorf("%v is not a valid status", updated.Status)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{	
			"$set": bson.M{
				"title":       updated.Title,
				"description": updated.Description,
				"status":      updated.Status,
				"duedate":     updated.DueDate,
		},
	}

	result, err := s.collection.UpdateOne(ctx,filter,update)
	if err != nil{
		return models.Task{}, err
	}
	if result.MatchedCount == 0{
		return models.Task{}, fmt.Errorf("no task with ID %s found", id)
	}
	
	updated.ID = id
	
	return  updated, nil
}

func (s *MongoTaskService) DeleteTask(id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no task with ID %s found", id)
	}

	return nil
}

