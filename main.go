package main

import (
	"fmt"
	"hydra/repository"
	"os"

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

	env := os.Getenv("GO_ENV")

	if env == "" {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Printf("An Error occured")
			return
		}
		return
	} 
	
	fmt.Printf("\n\nNo Env file current env is:: %s\n\n", env)
	
}