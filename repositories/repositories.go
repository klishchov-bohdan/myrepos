package repositories

import "mysite/repositories/models"

type UserRepositories interface {
	GetByEmail(email string) *models.User
}

type SupplierRepositories interface {
	GetAll() ([]*models.Supplier, error)
}
