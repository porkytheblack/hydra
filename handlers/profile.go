package handlers

import (
	"hydra/db/models"
	"hydra/dtos"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProfile(ctx *fiber.Ctx, db *gorm.DB) error {
	auth_0_user := &dtos.Auth0User{}
	err := ctx.BodyParser(auth_0_user)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "Unable to parse request body",
			"data": nil,
			"status": 400,
		})
	}

	db.Model(&models.ProfileModel{}).Create(&models.ProfileModel{
		FirstName: &auth_0_user.GivenName,
		LastName: &auth_0_user.FamilyName,
		Email: &auth_0_user.Email,
		PicUrl: &auth_0_user.Picture,
		Auth0ID: &auth_0_user.UserID,
	})

	return nil
}