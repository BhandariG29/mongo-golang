package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BhandariG29/mongo-golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{
	MongoCollection *mongo.Collection
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	user.Id = primitive.NewObjectID()

	result, err := uc.MongoCollection.InsertOne(context.Background(), user);
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	log.Println("results", result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := models.User{}

	err = uc.MongoCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	result, err := uc.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: objID}})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "User not found or deletion failed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Deleted user %s\n", id)
}