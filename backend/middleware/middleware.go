package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-react-golang/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	createDBInstance()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTodos()
	json.NewEncoder(w).Encode(payload)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var todo models.TodoList
	_ = json.NewDecoder(r.Body).Decode(&todo)
	collection.InsertOne(context.TODO(), todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	todo := getTodoByID(params["id"])
	json.NewEncoder(w).Encode(todo)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	todoComplete(params["id"])
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteTodoByID(params["id"])
}

func UndoTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	todoUndo(params["id"])
}

func getAllTodos() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	var todos []bson.M
	if err = cur.All(context.Background(), &todos); err != nil {
		log.Println(err)
	}
	return todos
}

func getTodoByID(s string) bson.M {
	log.Println("get todo id: ", s)
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Println(err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var todo bson.M

	errr := collection.FindOne(context.Background(), filter).Decode(&todo)
	if errr != nil {
		log.Fatal(errr)
	}
	return todo
}

func todoComplete(id string) {
	log.Println("complete todo id: ", id)
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"completed": true,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func todoUndo(id string) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"completed": false,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteTodoByID(s string) {
	filter := bson.M{"id": s}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
