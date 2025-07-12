package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/routes"
)

func main() {
	godotenv.Load(".env")
	config.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
