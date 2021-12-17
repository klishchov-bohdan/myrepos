package repositories

import "mysite/repositories/models"

type UserRepositories interface {
	GetByEmail(email string) *models.User
}

type SupplierRepositories interface {
	GetAll() ([]*models.Supplier, error)
}

type ProductsRepositories interface {
	GetByCategory(category string) ([]*models.Product, error)
}

type OrdersRepositories interface {
	GetByID(ID uint64) (*models.Order, error)
}
