Easy To Test - Platform for Online Testing
Description
{
Easy To Test is a web-based platform that simplifies online testing for educators, organizations, and learners. 
It allows users to easily create, manage, and take tests. 
The platform has a user-friendly interface with customizable test templates and real-time grading. 
It also provides analytics to help track performance.
}

Key Features
{
Test Creation: Easily design tests with multiple question types (MCQs, short answers, true/false, etc.).
User Management: Supports different user roles (Admin, Instructor, Student).
Real-Time Grading: Instant results and detailed feedback for students.
Performance Analytics: Insights into individual and group performance.
Security: Secure access with role-based authentication.
}

Target Audience
{
Teachers and educators seeking an easy way to create and administer tests.
Companies conducting employee evaluations or certifications.
Individuals seeking a reliable platform for practicing tests.
}

Team Members
{
Kazybek Seitkazy
Nubeket Zhunussov
Miras Yerseiit
}

Screenshot of the Main Page

Add this screenshot to the /screenshots directory in your repository.

Getting Started
Prerequisites
GO lang: Make sure you have the latest version installed.
Database: MySQL/PostgreSQL for storing user and test data.
Git: For version control.
A web browser for accessing the platform.
Steps to Run the Project
1. Clone the Repository
bash
Копировать код
git clone https://github.com/Switch-is-case/Platform-for-online-testing.git  
cd testease  
2. Install Dependencies
Run the following command to start server:
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
3. Set Up the Database
Create a database in MySQL/PostgreSQL.
Update the database configuration in the .env file. Example:
env
Копировать код
DB_HOST=localhost  
DB_PORT=3306  
DB_USER=root  
DB_PASS=yourpassword  
DB_NAME=testease  
4. Start the Backend Server
bash
Копировать код
npm run server  
5. Start the Frontend
bash
Копировать код
npm run client  
6. Access the Application
Open your browser and navigate to:

arduino
Копировать код
http://localhost:3000  
Tools and Resources
Frontend: Html, Css, JavaScript,Bootstrap
Backend: Go lang
Database: MySQL/PostgreSQL
Version Control: Git and GitHub
