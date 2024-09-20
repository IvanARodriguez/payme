package main

import (
	"log"
	"os"

	"github.com/IvanARodriguez/payme/models"
	"github.com/IvanARodriguez/payme/repository"
	"github.com/IvanARodriguez/payme/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}

	database, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = models.RunMigration(database)

	if err != nil {
		log.Fatal("Could not migrate partners to the database")
	}

	repo := &repository.Repository{
		Database: database,
	}

	app := fiber.New()

	repo.SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
