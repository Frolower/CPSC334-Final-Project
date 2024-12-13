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

	// Tires
	//auth.POST("/assignTireToCar/:chassis_number", routers.AssignTireToCarHandler(db))
	//auth.POST("/assignTireUsageToSession", routers.AssignTireInSessionHandler(db))

	// Championship
	//auth.POST("/createChampionship", routers.CreateChampionshipHandler(db))
	//auth.POST("/assignStageToChampionship/:championship_id", routers.CreateStageHandler(db))
	//auth.POST("/assignSessionToStage/:stage_id", routers.CreateSessionHandler(db))
	//auth.POST("/assignLapToSession/:session_id", routers.CreateLapHandler(db))

	// Pilots
	auth.POST("/addPilotToTeam/:team_id", routers.AddPilotToTeamHandler(db))
	auth.POST("/assignPilotToCar", routers.AssignPilotToCarHandler(db))
	auth.POST("/assignPilotToChampionship", routers.AssignPilotToChampionshipHandler(db))
	auth.POST("/assignPilotToStage", routers.AssignPilotToStageHandler(db))
	auth.POST("/assignPilotToSession", routers.AssignPilotToSessionHandler(db))

	// Staff
	//auth.POST("/addStaffToTeam/:team_id", routers.AddStaffToTeamHandler(db))
	//auth.POST("/assignStaffToCar/:chassis_number", routers.AssignStaffToCarHandler(db))

	// Setup
	//auth.POST("/assignSetupToCar/:chassis_number", routers.CreateSetupHandler(db))
	//auth.POST("/assignSetupToSession/:session_id", routers.AssignSetupToSessionHandler(db))
	//auth.POST("/addParameterToSetup/:setup_id", routers.AddParameterToSetupHandler(db))
	//auth.PUT("/updateParameter/:setup_id", routers.UpdateParameterHandler(db))

	// Delete endpoints
	auth.DELETE("/deleteTeam/:team_id", routers.DeleteTeamHandler(db))
	auth.DELETE("/deleteCar/:chassis_number", routers.DeleteCarHandler(db))
	auth.DELETE("/deletePart/:part_id", routers.DeletePartHandler(db))
	auth.DELETE("/deleteTire/:tire_id", routers.DeleteTireHandler(db))
	auth.DELETE("/deleteChampionship/:championship_id", routers.DeleteChampionshipHandler(db))
	auth.DELETE("/deleteStage/:stage_id", routers.DeleteStageHandler(db))
	auth.DELETE("/deleteSession/:session_id", routers.DeleteSessionHandler(db))
	auth.DELETE("/deleteLap/:session_id/:lap_number", routers.DeleteLapHandler(db))
	auth.DELETE("/deletePilot/:document_number/:first_name/:last_name", routers.DeletePilotHandler(db))
	auth.DELETE("/deleteStaff/:document_number/:first_name/:last_name", routers.DeleteStaffHandler(db))
	//auth.DELETE("/deleteSetup/:setup_id", routers.DeleteSetupHandler(db))
	//auth.DELETE("/deleteParameter/:setup_id/:attribute", routers.DeleteParameterHandler(db))

	// Get data
	auth.GET("/getCarsByUser", routers.GetCarsByUserHandler(db))
	auth.GET("/getPartsByUser", routers.GetPartsByUserHandler(db))
	auth.GET("/getTiresByUser", routers.GetTiresByUserHandler(db))
	auth.GET("/getChampionshipsByUser", routers.GetChampionshipsByUserHandler(db))
	auth.GET("/getStagesByUser", routers.GetStagesByUserHandler(db))
	auth.GET("/getSessionsByUser", routers.GetSessionsByUserHandler(db))
	auth.GET("/getLapsByUser", routers.GetLapsByUserHandler(db))
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

	// Run the server on port 8080
	r.Run(":8080")
}
