package main

import (
	"Ariadne_Management/routers"
	servicies "Ariadne_Management/services"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetch the individual environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Check if any required environment variable is missing
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("One or more database environment variables are not set")
	}

	// Construct the database URL using the components
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not open the database: ", err)
	}
	defer db.Close()

	// Ping the database to check if the connection is working
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Initialize the Gin router
	r := gin.Default()

	// Enable CORS for all routes and allow the 'Authorization' header
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * 3600, // 24 hours
	}))

	r.POST("/signup", routers.RegisterUser(db))
	r.POST("/login", routers.LoginUser(db))

	r.POST("/createTeam", servicies.AuthenticateJWT(), routers.CreateTeam(db))
	r.GET("/getTeams", servicies.AuthenticateJWT(), routers.GetUserTeams(db))

	// Run the server on port 8080
	r.Run(":8080")
}
