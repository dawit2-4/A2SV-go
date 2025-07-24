package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db        = "tasks"
	collName  = "tasks"
	userColl  = "users"
)

var mongoClient *mongo.Client

func ConnectDatabase(connectionString string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetMaxPoolSize(100).
		SetMinPoolSize(10)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	mongoClient = client
	log.Println("MongoDB connection established")

	// Create indexes
	if err := initIndexes(client.Database(db).Collection(collName)); err != nil {
		log.Fatal("Failed to create task indexes: ", err)
	}
	if err := initIndexes(client.Database(db).Collection(userColl)); err != nil {
		log.Fatal("Failed to create user indexes: ", err)
	}
}

func initIndexes(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"_id": 1}},
		{Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true)},
	})
	return err
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return mongoClient
}

func GetTaskCollection() *mongo.Collection {
	return GetMongoClient().Database(db).Collection(collName)
}

func DisconnectDatabase() {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
		log.Println("MongoDB connection closed")
	}
}