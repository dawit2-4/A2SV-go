package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie"`
	Actors  []string 					 `json:"actors"`
}

func InsertOne(movie Movie) {
	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted a record with id:", inserted.InsertedID)
}

func InsertMany(movies []Movie) error{
	collection := mongoClient.Database(db).Collection(collName)

	newMovies := make([]interface{}, len(movies)); {
		for i, movie := range movies {
			newMovies[i] = movie
		}
	}
	result, err := collection.InsertMany(context.TODO(),newMovies)

	if err != nil{
		panic(err)
	}

	log.Println(result)
	return  err
}

func UpdateMovie (movieId string, movie Movie) error {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil{
		return err
	}

	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"movie": movie.Movie, "actors": movie.Actors}}

	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil{
		return err
	}

	fmt.Println("New record: ", result)
	return nil
}

func DeleteMovie (movieId string) error {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil{
		return err
	}
	filter := bson.M{"_id":id}
	collection := mongoClient.Database(db).Collection(collName)

	result, err := collection.DeleteOne(context.TODO(),filter)
	if err != nil {
		return err
	}

	fmt.Println("Deleted result: ", result)
	return nil
}

func Find(movieName string) Movie {
	var result Movie

	filter := bson.D{{Key: "movie", Value: movieName}}
	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		log.Fatal(err)
	}
	return result
}

func FindAll(movieName string) []Movie {
	var result []Movie

	filter := bson.D{{Key: "movie", Value: movieName}}
	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &result)
	if err != nil{
		log.Fatal(err)
	}

	return result
}

func ListAll(movieName string) []Movie {
	var result []Movie


	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), result)
	if err != nil{
		log.Fatal(err)
	}

	return result
}

func DeleteAll() error {
	collection := mongoClient.Database(db).Collection(collName)
	deltedRecordsCount, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted records:", deltedRecordsCount)
	return err
}