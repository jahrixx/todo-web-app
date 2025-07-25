package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(UserID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	fmt.Println("SIGNING WITH:", os.Getenv("JWT_SECRET"))
	return token.SignedString([]byte(secret))
}

func GenerateSecureToken(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
