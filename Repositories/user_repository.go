package Repositories

import (
	"context"
	"errors"
	"fmt"
	"task_manger/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoUserRepository implements Domain.UserRepository using MongoDB.
type MongoUserRepository struct {
	collection *mongo.Collection
}

// CreateUser implements Domain.UserRepository.
func (m *MongoUserRepository) CreateUser(ctx context.Context, user Domain.User) (Domain.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err){
			return Domain.User{}, fmt.Errorf("username already taken")
		}
		return Domain.User{}, fmt.Errorf("failed to create user")
	}
	user.Password = ""
	return user, nil
}

// GetUserByUsername implements Domain.UserRepository.
func (m *MongoUserRepository) GetUserByUsername(ctx context.Context, username string) (Domain.User, error) {
	var user Domain.User
	err := m.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil{
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Domain.User{}, fmt.Errorf("user not found")
		}
		return Domain.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

// NewMongoUserRepository creates a new MongoUserRepository
func NewMongoUserRepository(client *mongo.Client, dbName, collName string) Domain.UserRepository {
	collection := client.Database(dbName).Collection(collName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"_id": 1},
		Options: options.Index(),
	})

	if err != nil {
		panic(fmt.Errorf("failed to create user index: %w", err))
	}

	return &MongoUserRepository{collection: collection}
}
