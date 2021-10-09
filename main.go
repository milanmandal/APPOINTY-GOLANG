package main
import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


//USERS MODEL
type Person struct {
	ID        	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string             `json:"name,omitempty" bson:"name,omitempty"`
	Email		string			   `json:"email,omitempty" bson:"email,omitempty"`
	Password  	string			   `json:"password,omitempty" bson:"password,omitempty"`
}


//POST MODEL
type Post struct {
	ID        	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id			string			   `json:"id,omitempty" bson:"id,omitempty"`
	Caption		string             `json:"name,omitempty" bson:"name,omitempty"`
	ImageURL	string			   `json:"email,omitempty" bson:"email,omitempty"`
	Timestamp  	string			   `json:"password  " bson:"password,omitempty"`
}


//CREATE USER
func CreateUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("instagram").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

//GET USER BY ID
func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("instagram").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

//CREATE POSTS
func CreatePosts (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post Post
	json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("instagram").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}


//GET POSTS
func GetPost (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("instagram").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}


//GET ALL POSTS OF THAT USER
func GetUserPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var userpost []Post
	params := mux.Vars(request)
	userid, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("instagram").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.D{{"id", userid}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		userpost = append(userpost, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(userpost)
}


func main(){
	fmt.Println("Starting the application...")
	clientOptions := options.Client().ApplyURI("mongodb+srv://appointy:appointy@instagram.tgxv5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _= mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	router.HandleFunc("/users", CreateUsers).Methods("POST")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	router.HandleFunc("/posts", CreatePosts).Methods("POST")
	router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	router.HandleFunc("/posts/users/{id}", GetUserPosts).Methods("GET")

	http.ListenAndServe(":8000", router)

}