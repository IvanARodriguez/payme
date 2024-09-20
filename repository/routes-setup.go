package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

// SetupRoutes method should be part of the repository package
func (repo *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/businesses", repo.CreateBusiness)
	api.Get("/businesses", repo.GetBusinesses)
}
