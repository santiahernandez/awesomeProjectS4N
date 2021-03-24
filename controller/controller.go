package controller

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()
var uri = os.Getenv("DB_CONN")

func createUser(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var users User
	collection := client.Database("UsersDB").Collection("user")
	reqBody, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	insertResult, err := collection.InsertOne(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single user: ", insertResult.InsertedID)
	fmt.Fprint(w, "Inserted a single user: ", insertResult.InsertedID)
}
