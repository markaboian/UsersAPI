package store

import "UsersAPI/internal/domain"

type Interface interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error

	// Product methods
	GetProductById(id int) (*domain.Product, error)
	AddProduct(name string, price float64, userId int) (*domain.Product, error)
}