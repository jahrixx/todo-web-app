package models

import (
	"database/sql"
	"time"

	"github.com/jahrixx/todo-web-app/config"
)

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Username      string `json:"username"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	BirthDate     string `json:"birthdate"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	ResetToken    string `gorm:"size: 255"`
	ResetTokenExp time.Time
}

func FindOrCreateUserByEmail(email string, username string) (*User, error) {
	var user User

	err := config.DB.QueryRow("SELECT id, email, username FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Username)

	switch err {
	case nil:
		return &user, nil
	case sql.ErrNoRows:
		user = User{
			Email:    email,
			Username: username,
		}

		_, err := config.DB.Exec("INSERT INTO users (email, username) VALUES ($1, $2) RETURNING id", user.Email, user.Username)

		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
