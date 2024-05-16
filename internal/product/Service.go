package product

import (
	"UsersAPI/internal/domain"
)

type IService interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error

	//Products methods
	GetProductById(id int) (*domain.Product, error)
	AddProduct(name string, price float64, userId int) (*domain.Product, error)
	GetProductsByUserId(userId int) (*[]domain.ProductWithUser, error)
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetUserById(id int) (*domain.User, error) {
	user, err := s.Repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error) {
	user, err := s.Repository.AddUser(name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(id int, user *domain.User) (*domain.User, error) {
	user, err := s.Repository.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(id int) error {
	err := s.Repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

// Products methods

func (s *Service) GetProductById(id int) (*domain.Product, error) {
	product, err := s.Repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) AddProduct(name string, price float64, userId int) (*domain.Product, error) {
	product, err := s.Repository.AddProduct(name, price, userId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProductsByUserId(userId int) (*[]domain.ProductWithUser, error) {
	productsWithUser, err := s.Repository.GetProductsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return productsWithUser, nil
}

func (s *Service) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	product, err := s.Repository.UpdateProduct(id, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}