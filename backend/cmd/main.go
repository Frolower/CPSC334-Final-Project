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

	// User routes
	r.POST("/signup", routers.RegisterUserHandler(db))
	r.POST("/login", routers.LoginUserHandler(db))

	// Protected routes (must have JWT)
	auth := r.Group("/")
	auth.Use(servicies.AuthenticateJWT())

	// Teams
	auth.POST("/createTeam", routers.CreateTeamHandler(db))
	auth.GET("/getTeams", routers.GetUserTeamsHandler(db))
	auth.POST("/assignCarToTeam/:team_id", routers.AssignCarToTeamHandler(db))

	// Parts
	auth.POST("/assignPartToCar/:chassis_number", routers.AssignPartToCarHandler(db))

	// Delete endpoints
	auth.DELETE("/deleteTeam/:team_id", routers.DeleteTeamHandler(db))
	auth.DELETE("/deleteCar/:chassis_number", routers.DeleteCarHandler(db))
	auth.DELETE("/deletePart/:part_id", routers.DeletePartHandler(db))
	auth.DELETE("/deleteTire/:tire_id", routers.DeleteTireHandler(db))
	auth.DELETE("/deleteChampionship/:championship_id", routers.DeleteChampionshipHandler(db))
	auth.DELETE("/deleteStage/:stage_id", routers.DeleteStageHandler(db))
	auth.DELETE("/deleteSession/:session_id", routers.DeleteSessionHandler(db))
	auth.DELETE("/sessions/:session_id/laps/:lap_number", routers.DeleteLapHandler(db))

	// Get data
	auth.GET("/getCarsByUser", routers.GetCarsByUserHandler(db))
	auth.GET("/getPartsByUser", routers.GetPartsByUserHandler(db))
	auth.GET("/getTiresByUser", routers.GetTiresByUserHandler(db))
	auth.GET("/getChampionshipsByUser", routers.GetChampionshipsByUserHandler(db))
	auth.GET("/getStagesByUser", routers.GetStagesByUserHandler(db))
	auth.GET("/getSessionsByUser", routers.GetSessionsByUserHandler(db))
	auth.GET("/sessions/:session_id/laps", routers.GetLapsBySessionHandler(db))
	auth.GET("/sessions/:session_id/laps/:lap_number", routers.GetLapByKeyHandler(db))

	auth.GET("/getPilotsByUser", routers.GetPilotsByUserHandler(db))
	auth.GET("/getStaffByUser", routers.GetStaffByUserHandler(db))
	auth.GET("/getSetupsByUser", routers.GetSetupsByUserHandler(db))
	auth.GET("/getParametersByUser", routers.GetParametersByUserHandler(db))

	// Analytics
	auth.GET("/avgPilotFinishPosition/:document_number/:first_name/:last_name", routers.GetAvgPilotFinishPositionHandler(db))
	auth.GET("/avgPilotFinishPositionPerStage/:document_number/:first_name/:last_name", routers.GetAvgPilotFinishPositionPerStageHandler(db))
	auth.GET("/avgCarFinishPosition", routers.GetAvgCarFinishPositionHandler(db))
	auth.GET("/avgCarFinishPositionPerChampionship", routers.GetAvgCarFinishPositionPerChampionshipHandler(db))
	auth.GET("/analyzeFastestLap/:session_id", routers.GetFastestLapHandler(db))
	auth.GET("/analyzeAverageLap/:session_id", routers.GetAverageLapHandler(db))
	auth.GET("/analyzePartsCount/:chassis_number", routers.GetPartsCountForCarHandler(db))

	// Additional Lap Endpoints:
	auth.POST("/sessions/:session_id/laps", routers.CreateLapHandler(db))
	auth.PUT("/sessions/:session_id/laps/:lap_number", routers.UpdateLapHandler(db))

	// Run the server on port 8080
	r.Run(":8080")
}
