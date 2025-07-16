package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/models"
	"github.com/jahrixx/todo-web-app/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashedPassword)

	_, err := config.DB.Exec("INSERT INTO users (email, password, username, firstName, lastName, birthdate, gender, address, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", user.Email, user.Password, user.Username, user.FirstName, user.LastName, user.BirthDate, user.Gender, user.Address, user.Phone)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"message": "User Registered Successfully"})
}

func Login(c *fiber.Ctx) error {
	var userCredentials models.User
	if err := c.BodyParser(&userCredentials); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var storedUser models.User

	err := config.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", userCredentials.Email).Scan(&storedUser.ID, &storedUser.Password)

	if err == sql.ErrNoRows {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid Credentials"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userCredentials.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	token, err := utils.GenerateJWT(storedUser.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate JWT"})
	}

	return c.JSON(fiber.Map{"token": token})
}
