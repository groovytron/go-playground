package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golangplayground/app"
	"golangplayground/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read .env
	loadEnvFile()

	app := app.New()

	// Connect to database
	// database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db_host, db_port, db_name, db_user, db_password := fetchDatabaseParametersFromEnv()
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", db_host, db_port, db_name, db_user, db_password)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	check(err)

	// Migrate automatically
	// TODO: only migrate when a flag is given
	database.AutoMigrate(&models.Todo{}, &models.Task{})

	// Generate data
	// TODO: only seed data when a flag is given
	// todo := models.Todo{Name: "My First Todo", Description: "This is amazing"}

	// database.Create(&todo)

	// taskOne := models.Task{Name: "Clean The Kitchen", TodoID: todo.ID}
	// taskTwo := models.Task{Name: "Clean The Kitchen", TodoID: todo.ID}

	// database.Create(&taskOne)
	// database.Create(&taskTwo)

	// TODO: migrate to PostgeSQL and create connection pool and close connections when done
	app.Database = database

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App server started")

	err = http.ListenAndServe(":9000", nil)

	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func fetchDatabaseParametersFromEnv() (string, string, string, string, string) {
	host, variableIsSet := os.LookupEnv("POSTGRES_HOST")

	if !variableIsSet {
		log.Fatal("POSTGRES_HOST variable is not set")
	}

	port, variableIsSet := os.LookupEnv("POSTGRES_PORT")

	if !variableIsSet {
		log.Fatal("POSTGRES_PORT variable is not set")
	}

	database, variableIsSet := os.LookupEnv("POSTGRES_DB")

	if !variableIsSet {
		log.Fatal("POSTGRES_DATABASE variable is not set")
	}

	user, variableIsSet := os.LookupEnv("POSTGRES_USER")

	if !variableIsSet {
		log.Fatal("POSTGRES_USER variable is not set")
	}

	password, variableIsSet := os.LookupEnv("POSTGRES_PASSWORD")

	if !variableIsSet {
		log.Fatal("POSTGRES_PASSWORD variable is not set")
	}

	return host, port, database, user, password
}
