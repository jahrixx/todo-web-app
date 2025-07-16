package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jahrixx/todo-web-app/config"
	"github.com/jahrixx/todo-web-app/models"
	"github.com/jahrixx/todo-web-app/utils"
	"golang.org/x/oauth2"
)

func GitHubLogin(c *fiber.Ctx) error {
	url := config.GitHubOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func GitHubCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := config.GitHubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(500).SendString("Token exchange failed")
	}

	client := config.GitHubOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(500).SendString("Failed to fetch user data")
	}
	defer resp.Body.Close()

	var userData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userData)

	email := userData["email"]
	if email == nil {
		// Get email from separate endpoint
		emailResp, _ := client.Get("https://api.github.com/user/emails")
		defer emailResp.Body.Close()
		var emails []map[string]interface{}
		json.NewDecoder(emailResp.Body).Decode(&emails)
		for _, e := range emails {
			if e["primary"].(bool) {
				email = e["email"]
				break
			}
		}
	}

	username := userData["login"].(string)

	// ✅ Find or create user in DB
	user, err := models.FindOrCreateUserByEmail(fmt.Sprintf("%v", email), username)
	if err != nil {
		return c.Status(500).SendString("User error")
	}

	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(500).SendString("JWT generation failed")
	}

	return c.Redirect(fmt.Sprintf("%s/oauth/callback?token=%s", os.Getenv("FRONTEND_URL"), jwtToken))
}
