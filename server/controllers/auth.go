package controllers

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/sendgrid/sendgrid-go"
	// "github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/crypto/bcrypt"

	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/models"
	"github.com/jahrixx/todo-web-app/utils"
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

// ForgotPassword handles the password reset request and sends email
func ForgotPassword(c *fiber.Ctx) error {
	type Request struct {
		Email string `json:"email"`
	}
	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, email FROM users WHERE email = $1", body.Email).Scan(&user.ID, &user.Email)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// Generate token and expiration
	token := utils.GenerateSecureToken(32)
	exp := time.Now().Add(1 * time.Hour)

	_, err = config.DB.Exec("UPDATE users SET reset_token = $1, reset_token_exp = $2 WHERE id = $3", token, exp, user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save reset token"})
	}

	resetURL := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("FRONTEND_URL"), token)

	// Send reset email
	err = sendResetEmail(user.Email, resetURL)
	if err != nil {
		fmt.Println("Email sending error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to send reset email"})
	}

	return c.JSON(fiber.Map{"message": "Reset link sent to email"})
}

// ResetPassword updates the user's password using the token
func ResetPassword(c *fiber.Ctx) error {
	type Request struct {
		Token       string `json:"token"`
		NewPassword string `json:"password"`
	}
	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, reset_token_exp FROM users WHERE reset_token = $1", body.Token).Scan(&user.ID, &user.ResetTokenExp)
	if err == sql.ErrNoRows {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or expired token"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if time.Now().After(user.ResetTokenExp) {
		return c.Status(400).JSON(fiber.Map{"error": "Token expired"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 14)

	_, err = config.DB.Exec("UPDATE users SET password = $1, reset_token = NULL, reset_token_exp = NULL WHERE id = $2", string(hashed), user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password updated successfully"})
}

func sendResetEmail(toEmail, resetURL string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	addr := host + ":" + port

	subject := "Subject: Password Reset\n"
	body := fmt.Sprintf("Click this link to reset your password:\n\n%s", resetURL)
	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, host)
	return smtp.SendMail(addr, auth, from, []string{toEmail}, msg)
}

// sendResetEmail sends the password reset link using SendGrid
// func sendResetEmail(toEmail, resetURL string) error {
// 	from := mail.NewEmail("Your App", os.Getenv("FROM_EMAIL"))
// 	to := mail.NewEmail("", toEmail)
// 	subject := "Password Reset"
// 	plainTextContent := "Click this link to reset your password: " + resetURL
// 	htmlContent := "<p>Click <a href=\"" + resetURL + "\">here</a> to reset your password.</p>"

// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	_, err := client.Send(message)
// 	return err
// }
