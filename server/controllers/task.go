package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/models"
)

// GET /api/tasks - get all tasks of logged-in user
func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	rows, err := config.DB.Query("SELECT id, title, description, completed, user_id FROM tasks WHERE user_id = $1", userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.UserID); err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// POST /api/tasks - create task
func CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	task.ID = uuid.New().String()

	_, err := config.DB.Exec("INSERT INTO tasks (id, user_id, title, description, completed) VALUES ($1, $2, $3, $4, $5)",
		task.ID, userID, task.Title, task.Description, false)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(201).JSON(task)
}

// PUT /api/tasks/:id - update task
func UpdateTask(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	taskID := c.Params("id")

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	_, err := config.DB.Exec(`
		UPDATE tasks SET title = $1, description = $2, completed = $3
		WHERE id = $4 AND user_id = $5`,
		task.Title, task.Description, task.Completed, taskID, userID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.JSON(fiber.Map{"message": "Task updated"})
}

// DELETE /api/tasks/:id - delete task
func DeleteTask(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	taskID := c.Params("id")

	_, err := config.DB.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", taskID, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.JSON(fiber.Map{"message": "Task deleted"})
}
