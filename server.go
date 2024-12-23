package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Name      string    `json:"name" bson:"name"`
    Email     string    `json:"email" bson:"email"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func main() {
	// MongoDB connection string
	uri := "mongodb://localhost:27017"

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Access a database and collection
	database := client.Database("mycollection")
	collection := database.Collection("users")

	// API endpoints
	http.HandleFunc("/api", handleRequests)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handleGetUsers(w, r, collection)
	})

	// Serve the HTML file for browser access
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		handlePostRequest(w, r)
	} else if r.Method == http.MethodGet {
		handleGetRequest(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON message",
		})
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}

	// Decode the JSON body into a generic map
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON message",
		})
		return
	}

	// Check if the "message" key exists
	messageValue, exists := requestData["message"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON message",
		})
		return
	}

	// Check if the "message" value is a string
	message, ok := messageValue.(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON message",
		})
		return
	}

	// Check if the "message" value is an empty string
	if message == "" {
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "success",
			Message: "Data successfully received",
		})
		return
	}

	// For a valid non-empty message
	json.NewEncoder(w).Encode(ResponseData{
		Status:  "success",
		Message: "Data successfully received",
	})
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{
		Status:  "success",
		Message: "Data successfully received",
	}
	json.NewEncoder(w).Encode(response)
}

func handleGetUsers(w http.ResponseWriter, r *http.Request, collection *mongo.Collection) {
    var users []User

    // Perform the query to get all users
    cursor, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        log.Printf("Error fetching users: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(ResponseData{
            Status:  "fail",
            Message: "Error fetching users",
        })
        return
    }
    defer cursor.Close(context.Background())

    // Decode each document into a User struct
    for cursor.Next(context.Background()) {
        var user User
        if err := cursor.Decode(&user); err != nil {
            log.Printf("Error decoding user: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(ResponseData{
                Status:  "fail",
                Message: "Error decoding user",
            })
            return
        }
        users = append(users, user)
    }

    // Check if there were any errors during iteration
    if err := cursor.Err(); err != nil {
        log.Printf("Cursor error: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(ResponseData{
            Status:  "fail",
            Message: "Cursor error",
        })
        return
    }

    // Return the list of users
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Serve the HTML file for browser requests
func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
