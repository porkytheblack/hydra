package repository

import (
	"hydra/db"
	"hydra/handlers"
	"hydra/middleware"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) InitRepo() {

	_db, err := db.NewConnection(&db.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("SSLMode"),
	})


	

	if err != nil {
		log.Fatalf("Unable to initialize DB %s", err)
	}

	// db.RunMigrations(_db)

	repo.DB = _db
}


func (repo *Repository) SetupRoutes(app *fiber.App)  {
	app.Use(logger.New())
	app.Use(middleware.AuthMiddleWare)
	api := app.Group("/v1")

	handlers.GenerateHandlers(&api, repo.DB)

}
