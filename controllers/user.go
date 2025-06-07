package controllers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/v2/bson"
// 	"go.mongodb.org/mongo-driver/v2/bson/primitive"
// 	"go.mongodb.org/mongo-driver/v2/mongo"
// 	"go.mongodb.org/mongo-driver/v2/mongo/options"

// 	// "gopkg.in/mgo.v2"
// 	// "gopkg.in/mgo.v2/bson"

// 	"github.com/BhandariG29/mongo-golang/models"
// 	// "go.mongodb.org/mongo-driver/bson"
// 	// "go.mongodb.org/mongo-driver/bson/primitive"
// 	// "go.mongodb.org/mongo-driver/mongo"
// 	// "go.mongodb.org/mongo-driver/mongo/options"
// )

// const connectionString = "mongodb+srv://bhandarilaksh:<db_password>@cluster0.svgah.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
// const dbName = "database001"
// const collName = "crudop"

// var collection *mongo.Collection

// func init(){
// 	clientOption := options.Client().ApplyURI(connectionString)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, clientOption)
// 	if err != nil {
// 		log.Fatal("MongoDB connection error:", err)
// 	}
// 	fmt.Println("MongoDB coneection successful")

// 	collection = client.Database(dbName).Collection(collName)
// 	fmt.Println("collection instance is ready")
// }

// type UserController struct{
// 	// session *mgo.Session
// }

// func NewUserController() *UserController{
// 	return &UserController{}
// }

// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request){
// 	id := mux.Vars(r)["id"]

// 	// if !bson.IsObjectIdHex(id){
// 	// 	w.WriteHeader(http.StatusNotFound)
// 	// 	return
// 	// }
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	// oid := bson.ObjectIdHex(id)

// 	user := models.User{}

// 	// err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusNotFound)
// 	// 	return
// 	// }
// 	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	// uj, err := json.Marshal(u)
// 	// if err != nil {
// 	// 	log.Println("JSON Marshal error:", err)
// 	// }

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// 	// w.WriteHeader(http.StatusOK)
// 	// fmt.Fprintf(w, "%s\n", uj)
// }

// func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request){
// 	user := models.User{}

// 	// err := json.NewDecoder(r.Body).Decode(&u)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid user data", http.StatusBadRequest)
// 		return
// 	}

// 	// u.Id = bson.NewObjectId()
// 	user.Id = primitive.NewObjectID()

// 	// err = uc.session.DB("mongo-golang").C("users").Insert(u)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	_, err = collection.InsertOne(context.TODO(), user)
// 	if err != nil {
// 		http.Error(w, "Failed to create user", http.StatusInternalServerError)
// 		return
// 	}

// 	// uj, err := json.Marshal(u)
// 	// if err != nil {
// 	// 	log.Println("JSON Marshal error:", err)
// 	// }

// 	w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(http.StatusCreated)
// 	// fmt.Fprintf(w, "%s\n", uj)
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(user)
// }

// func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request){
// 	id := mux.Vars(r)["id"]

// 	// if !bson.IsObjectIdHex(id){
// 	// 	w.WriteHeader(http.StatusNotFound)
// 	// 	return
// 	// }
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	// oid := bson.ObjectIdHex(id)

// 	// err := uc.session.DB("mongo-golang").C("users").RemoveId(oid)
// 	// if err != nil {
// 	// 	w.WriteHeader(404)
// 	// }
// 	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
// 	if err != nil || result.DeletedCount == 0 {
// 		http.Error(w, "User not found or deletion failed", http.StatusNotFound)
// 		return
// 	}

// 	// w.WriteHeader(http.StatusOK)
// 	// fmt.Fprintf(w, "Deleted user %s\n", oid.Hex()) //oid or oid.Hex()?
// 	fmt.Fprintf(w, "Deleted user %s\n", id)
// }