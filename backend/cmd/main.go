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
	auth.DELETE("/deleteTeam/:team_id", routers.DeleteTeamHandler(db))

	// Cars
	auth.POST("/assignCarToTeam/:team_id", routers.AssignCarToTeamHandler(db))
	auth.GET("/getCars/:team_id", routers.GetCarsByTeamHandler(db))
	auth.PUT("/updateCar/:chassis_number", routers.UpdateCarHandler(db))
	auth.DELETE("/deleteCar/:chassis_number", routers.DeleteCarHandler(db))

	// Parts
	auth.POST("/assignPartToCar/:chassis_number", routers.AssignPartToCarHandler(db))
	auth.GET("/getPartsByCar/:chassis_number", routers.GetPartsByCarHandler(db))
	auth.PUT("/updatePart/:part_id", routers.UpdatePartHandler(db))
	auth.DELETE("/deletePart/:part_id", routers.DeletePartHandler(db))

	// Tires
	auth.POST("/assignTireToCar/:chassis_number", routers.AssignTireToCarHandler(db))
	auth.GET("/getTiresByCar/:chassis_number", routers.GetTiresByCarHandler(db))
	auth.PUT("/updateTire/:tire_id", routers.UpdateTireHandler(db))
	auth.DELETE("/deleteTire/:tire_id", routers.DeleteTireHandler(db))

	// Championships
	auth.POST("/createChampionship", routers.CreateChampionshipHandler(db))
	auth.GET("/getChampionships", routers.GetChampionshipsHandler(db))
	auth.GET("/getChampionships/:team_id", routers.GetChampionshipsByTeamIDHandler(db))
	auth.PUT("/updateChampionship/:championship_id", routers.UpdateChampionshipHandler(db))
	auth.DELETE("/deleteChampionship/:championship_id", routers.DeleteChampionshipHandler(db))

	// Stages
	auth.POST("/createStage", routers.CreateStageHandler(db))
	auth.GET("/getStages", routers.GetStagesHandler(db))
	auth.GET("/getStage/:championship_id", routers.GetStagesByChampionshipIDHandler(db))
	auth.PUT("/updateStage/:stage_id", routers.UpdateStageHandler(db))
	auth.DELETE("/deleteStage/:stage_id", routers.DeleteStageHandler(db))

	// Sessions
	auth.POST("/createSession", routers.CreateSessionHandler(db))
	auth.GET("/getSessions", routers.GetSessionsHandler(db))
	auth.GET("/getSession/:stage_id", routers.GetSessionsByStageIDHandler(db))
	auth.PUT("/updateSession/:session_id", routers.UpdateSessionHandler(db))
	auth.DELETE("/deleteSession/:session_id", routers.DeleteSessionHandler(db))

	// Laps
	auth.POST("/createLap", routers.CreateLapHandler(db))
	auth.GET("/getLaps/:session_id", routers.GetLapsBySessionHandler(db))
	auth.PUT("/updateLap/:lap_id", routers.UpdateLapHandler(db))
	auth.DELETE("/deleteLap/:lap_id", routers.DeleteLapHandler(db))

	// Analysis
	auth.GET("/analyzeFastestLap/:session_id", routers.GetFastestLapHandler(db))
	auth.GET("/analyzeAverageLap/:session_id", routers.GetAverageLapHandler(db))
	auth.GET("/analyzePartsCount/:chassis_number", routers.GetPartsCountForCarHandler(db))
	auth.GET("/getTireCountForCar/:chassis_number", routers.GetTireCountForCarHandler(db))
	auth.GET("/getAverageTreadForCar/:chassis_number", routers.GetAverageTreadForCarHandler(db))
	auth.GET("/getTiresCountByTreadAndCompoundAndCar/:chassis_number/:tread/:compound", routers.GetTiresCountByTreadAndCompoundAndCarHandler(db))

	// Run the server on port 8080
	r.Run(":8080")
}
