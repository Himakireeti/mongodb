package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client
var clientOptions *options.ClientOptions = options.Client().
	ApplyURI("mongodb+srv://partyholic_user:partyholic_password@cluster0.xfyht.mongodb.net")

/*
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Starting the application...\n")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, clientOptions)
	// fmt.Fprintln(w, client)

	collection := client.Database("myFirstDatabase").Collection("firstCollection")
	// fmt.Fprintln(w, collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, episodes)

}*/

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("myFirstDatabase").Collection("firstCollection")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println(person)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database("myFirstDatabase").Collection("firstCollection")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `}"`))
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode((&person))
		people = append(people, person)
	}
	json.NewEncoder(response).Encode(people)

}

func main() {
	// http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/printall", GetDetails).Methods("GET")

	http.ListenAndServe(":12345", router)
}
