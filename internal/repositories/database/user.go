package database

import "mysite/internal/models"

type UserDBRepository struct{}

func (u UserDBRepository) GetByEmail(email string) models.User {
	// SELECT email, password_hash, created_at, FROM users WHERE email = email
	panic("implement me")
}
