package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(user RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     user.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
	return err == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour).Unix()

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "secret"

	tokenString, _ := token.SignedString([]byte(secret))
	
	return tokenString
}
