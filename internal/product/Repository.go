package product

import (
	"UsersAPI/internal/domain"
	"UsersAPI/package/store"
)

type IRepository interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error

	//Products methods
	GetProductById(id int) (*domain.Product, error)
	AddProduct(name string, price float64, userId int) (*domain.Product, error)
	GetProductsByUserId(idUser int) (*[]domain.ProductWithUser, error)
}

type Repository struct {
	Interface store.Interface
}

func (r *Repository) GetUserById(id int) (*domain.User, error) {
	user, err := r.Interface.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error){
	user, err := r.Interface.AddUser(name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(id int, user *domain.User) (*domain.User, error) {
	user, err := r.Interface.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(id int) error {
	err := r.Interface.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

// Products methods

func (r *Repository) GetProductById(id int) (*domain.Product, error) {
	product, err := r.Interface.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) AddProduct(name string, price float64, userId int) (*domain.Product, error) {
	product, err := r.Interface.AddProduct(name, price, userId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) GetProductsByUserId(userId int) (*[]domain.ProductWithUser, error) {
	productsWithUser, err := r.Interface.GetProductsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return productsWithUser, nil
}