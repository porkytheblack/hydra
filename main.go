package main

import (
	"fmt"
	"hydra/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main() {
	loadEnv()
	repo := &repository.Repository{}
	app := fiber.New()
	repo.InitRepo()
	repo.SetupRoutes(app)
	app.Listen("0.0.0.0:8080")


}	

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("An Error occured")
	}
}