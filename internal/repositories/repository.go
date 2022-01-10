package repositories

import "github.com/klishchov-bohdan/myrepos/internal/models"

type UserRepositories interface {
	GetByEmail(email string) (user *models.User, err error)
	Create(user *models.User) (*models.User, error)
	GetByID(id int) (user *models.User, err error)
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
