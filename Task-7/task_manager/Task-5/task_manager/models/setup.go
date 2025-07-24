package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const db = "tasks"
const collName = "tasks"

var mongoClient *mongo.Client

func ConnectDatabase(){
	ctx , cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx,clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	mongoClient = client
	log.Println("MongoDB connection established")
}

func GetTaskCollection() *mongo.Collection {
	if mongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return mongoClient.Database(db).Collection(collName)
}

func DisconnectDatabase() {
	if mongoClient != nil{
		_ = mongoClient.Disconnect(context.TODO())
		log.Println("MongoDB connection closed")
	}
}