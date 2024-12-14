package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supratim531/gomongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

const dbName = "gonetflix"
const collectionName = "watchlist"
const connectionString = "mongodb+srv://supratim531:sayan2002@cluster0.okgdt49.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// connect with MongoDB
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// MongoDB connected successfully
	fmt.Println("MongoDB connected successfully:", client)
	collection = client.Database(dbName).Collection(collectionName)

	// Collection is ready
	fmt.Println("Collection instance is ready:", collection)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
	}

	insertedId := insertOneMovie(movie)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      insertedId,
		"message": "Movie inserted successfully",
	})
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	modifiedCount := updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":            params["id"],
		"modifiedCount": modifiedCount,
		"message":       "Marked as watched successfully",
	})
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deletedCount := deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":           params["id"],
		"deletedCount": deletedCount,
		"message":      "Deleted record successfully",
	})
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deletedCount := deleteAllMovies()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"deletedCount": deletedCount,
		"message":      "Deleted all records successfully",
	})
}

// MongoDB helper functions - file
// insert 1 record
func insertOneMovie(movie model.Netflix) interface{} {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted movie:", inserted.InsertedID)
	return inserted.InsertedID
}

// update 1 record
func updateOneMovie(movieId string) int64 {
	objectId, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectId}
	fmt.Printf("%v <-> %v\n", movieId, objectId)
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateByID(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
	return result.ModifiedCount
}

// delete 1 record
func deleteOneMovie(movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)
	return result.DeletedCount
}

// delete all records
func deleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)
	return result.DeletedCount
}

// get all records
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}
