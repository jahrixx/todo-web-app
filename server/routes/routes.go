package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jahrixx/todo-web-app/controllers"
	"github.com/jahrixx/todo-web-app/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// Protected task routes
	task := api.Group("/tasks", middleware.Protected())
	task.Get("/", controllers.GetTasks)
	task.Post("/", controllers.CreateTask)
	task.Put("/:id", controllers.UpdateTask)
	task.Delete("/:id", controllers.DeleteTask)
}
