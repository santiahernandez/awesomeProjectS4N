package controller

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()
var uri = os.Getenv("DB_CONN")

func apiStatus(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri) // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	var users User

	collection := client.Database("Test").Collection("user")

	err = json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		fmt.Print(err)
	}

	insertResult, err := collection.InsertOne(ctx, users)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(insertResult.InsertedID)
}
