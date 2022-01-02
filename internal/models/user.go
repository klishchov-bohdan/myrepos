package models

import (
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func NewUser(name string, Email string, password string) *User {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		ID:        rand.Intn(math.MaxInt32),
		Name:      name,
		Email:     Email,
		Password:  string(pwd),
		CreatedAt: time.Now().Format("01-02-2006 15:04:05"),
	}
}
