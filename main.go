package main

import (
	"os"

	studenthandler "github.com/KarMint26/echo-student-api/controllers/StudentHandler"
	"github.com/KarMint26/echo-student-api/storages"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialization Echo Framework
	e := echo.New()

	// Loaded ENV
	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal("Can't load environment variables")
	}

	// Open Connection to DB
	dbConfig := &storages.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	storages.ConnectionDatabase(dbConfig)

	 // Create an instance of the Database type
	db := &studenthandler.Database{
        DB: storages.GetDB(),
    }
	
	// Routes
	e.GET("/api/v1/students", db.GetStudents)
	e.POST("/api/v1/students", db.CreateStudent)
	e.GET("/api/v1/student/:id", db.GetStudentById)
	e.PUT("/api/v1/student/:id", db.UpdateStudent)
	e.DELETE("/api/v1/student/:id", db.DeleteStudent)

	// Starting Server
	e.Start(":8001")
}