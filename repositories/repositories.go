package repositories

import "mysite/repositories/models"

type UserRepositories interface {
	GetByEmail(email string) models.User
}