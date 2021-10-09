package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Posts struct {
	Id               string `json:"Id"`
	Caption          string `json:"Caption"`
	ImageUrl         string `json:"ImageUrl"`
	Posted_Timestamp string `json:"Posted_Timestamp"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) { //create user
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	log.Println(u.Id)
	log.Println(u.Name)
	log.Println(u.Email)
	log.Println(u.Password)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Access a MongoDB collection through a database
	col := client.Database("some_database").Collection("Some Collection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	// Declare a MongoDB struct instance for the document's fields and data
	oneDoc := User{
		FieldStr: u.Id,
		FieldStr: u.Name,
		FieldStr: u.Email,
		FieldStr: u.Password
	}
	fmt.Println("oneDoc TYPE:", reflect.TypeOf(oneDoc), "\n")

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr !=

		nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1) // safely exit script on error
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)

		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) { //create user
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var p Posts
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	log.Println(p.Id)
	log.Println(p.Caption)
	log.Println(p.ImageUrl)
	log.Println(p.Posted_Timestamp)
}

func GetUser(w http.ResponseWriter, r *http.Request) { // get user
store:
	user
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/post", CreatePost)

	//http.HandleFunc("/users/", GetUser)
	//http.HandleFunc("/post/", GetPost)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
