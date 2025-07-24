package Repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoTaskRepository implements Domain.TaskRepository using MongoDB.
type MongoTaskRepository struct {
	collection *mongo.Collection
}

// CreateTask implements Domain.TaskRepository.
func (m *MongoTaskRepository) CreateTask(ctx context.Context, task Domain.Task) (Domain.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, task)
	if err != nil {
		return Domain.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

// DeleteTask implements Domain.TaskRepository.
func (m *MongoTaskRepository) DeleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := m.collection.DeleteOne(ctx, bson.M{"_id":objID})
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("task not found: %s", id)
	}
	return nil
}

// GetAllTasks implements Domain.TaskRepository.
func (m *MongoTaskRepository) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil{
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	defer cursor.Close(ctx)
	var tasks []Domain.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks: %w", err)
	}
	return tasks, nil
}

// GetTaskByID implements Domain.TaskRepository.
func (m *MongoTaskRepository) GetTaskByID(ctx context.Context, id string) (Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return Domain.Task{}, fmt.Errorf("invalid ID format: %w", err)
	}

	var task Domain.Task
	err = m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Domain.Task{}, fmt.Errorf("task not found: %s", id)
		}
		return  Domain.Task{}, fmt.Errorf("failed to retrieve task:%w", err)
	}

	return task, nil
}

// UpdateTask implements Domain.TaskRepository.
func (m *MongoTaskRepository) UpdateTask(ctx context.Context, id string, task Domain.Task) (Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.Task{}, fmt.Errorf("invalid ID format: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"due_date":    task.DueDate,
		},
	}

	result, err := m.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil{
		return Domain.Task{}, fmt.Errorf("failed to update task: %w", err)
	}

	task.ID = objID
	if result.MatchedCount == 0 {
		return Domain.Task{}, fmt.Errorf("task not found: %s", id)
	}

	return task, nil
}

// NewMongoTaskRepository creates a new MongoTaskRepository
func NewMongoTaskRepository(client *mongo.Client, dbName, collName string) Domain.TaskRepository {
	collection := client.Database(dbName).Collection(collName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"_id": 1},
		Options: options.Index(),
	})
	if err != nil {
		panic(fmt.Errorf("failed to create task index: %w", err))
	}
	return &MongoTaskRepository{collection: collection}
}
