package repository

import (
	"net/http"
	"time"

	"github.com/IvanARodriguez/payme/models"
	"github.com/gofiber/fiber/v2"
)

type CreateBusinessDto struct {
	Name string `json:"name"`
}

func (r *Repository) CreateBusiness(ctx *fiber.Ctx) error {
	business := CreateBusinessDto{}
	err := ctx.BodyParser(&business)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid business data"})
		return err
	}
	newBusiness := models.Business{
		Name:      business.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.Database.Create(&newBusiness).Error; err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Failed to create business"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&newBusiness)
	return nil
}

func (r *Repository) GetBusinesses(ctx *fiber.Ctx) error {
	var businesses []models.Business
	err := r.Database.Find(&businesses).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Failed to retrieve businesses"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(businesses)
	return nil
}
