package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanmay-Thanvi/mongoapi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://tanmaythanvi15:Nita250502@cluster0.2kazzdn.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect with database

func init() {
	// Client Option
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to Mongodb
	client, err := mongo.Connect(context.TODO(), clientOption) // context is imp and if we dont know then use todo
	checkError(err)
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance reference
	fmt.Println("Collection reference is ready!")
}

// MongoDB Helpers - file

// Insert 1 record
// method name is small bcoz helper method and never get exported

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkError(err)
	fmt.Println("Inserted 1 movie in db with id : ", inserted.InsertedID, " & ", inserted)
}

// Update 1 record
func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkError(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkError(err)

	fmt.Println("Modified count : ", result.ModifiedCount, " : ", result)
}

// delete One record
func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkError(err)

	filter := bson.M{"_id": id}
	deleteCoumt, err := collection.DeleteOne(context.Background(), filter)
	checkError(err)

	fmt.Println("No. of Movies deleted : ", deleteCoumt.DeletedCount)
}

// delete All records
func deleteAllrecords() int64 {
	filter := bson.D{{}}
	deleteCoumt, err := collection.DeleteMany(context.Background(), filter)
	checkError(err)

	fmt.Println("No. of Movies deleted : ", deleteCoumt.DeletedCount)
	return deleteCoumt.DeletedCount
}

// Get all movies from database
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkError(err)
	var movies []primitive.M // here not using primitive many times give error

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkError(err)
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// Actual Controllers - in this file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params)
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllrecords()
	json.NewEncoder(w).Encode(count)
}
