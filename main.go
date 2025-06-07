package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/BhandariG29/mongo-golang/controllers"
)

var mongoClient *mongo.Client

func init(){
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	log.Println("env file loaded")

	//create mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("mongo connected")
}

func main(){
	// close the mongo connection
	defer mongoClient.Disconnect(context.Background())
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COL_NAME"))
	uc := controllers.UserController{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.DeleteUser).Methods("DELETE")

	fmt.Println("Starting server at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}