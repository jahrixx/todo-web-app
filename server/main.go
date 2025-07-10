package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/routes"
)

func main() {
	godotenv.Load(".env")
	config.Connect()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
