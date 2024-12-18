package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequests)

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
			Message: "Method not allowed",
		})
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON format",
		})
		return
	}

	// Validate the JSON structure
	if requestData.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseData{
			Status:  "fail",
			Message: "Invalid JSON message",
		})
		return
	}

	// Log the message to the console
	log.Println("Received message:", requestData.Message)

	// Respond with success
	json.NewEncoder(w).Encode(ResponseData{
		Status:  "success",
		Message: "Data successfully received",
	})
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{
		Status:  "success",
		Message: "GET request received",
	}
	json.NewEncoder(w).Encode(response)
}
