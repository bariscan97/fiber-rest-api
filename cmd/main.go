package main

import (
	"log"
	"os"
	database "todo_api/db"
	"todo_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	pool := database.Pool()

	routers := routes.NewRouters(app, pool)

	routers.InitRouter()

	routers.Start(":" + os.Getenv("PORT"))
}
