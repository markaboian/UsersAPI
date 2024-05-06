package store

import (
	"UsersAPI/internal/domain"
	"database/sql"
)

type Sql struct {
	DB *sql.DB
}

func (s *Sql) GetUserById(id int) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE id = ?"

	row := s.DB.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.LastName, &user.PlaceOfBirth, &user.DateOfBirth, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func(s *Sql) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error) {

	query := "INSERT INTO users (name, lastName, place_of_birth, date_of_birth, email) VALUES (?, ?, ?, ?, ?)"

	result, err := s.DB.Exec(query, name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Id: int(id),
		Name: name,
		LastName: lastName,
		PlaceOfBirth: placeOfBirth,
		DateOfBirth: dateOfBirth,
		Email: email,
	}

	return user, nil
}

func (s *Sql) UpdateUser(id int, user *domain.User) (*domain.User, error) {

	query := "UPDATE users SET "

	var values []interface{}

	if user.Name != "" {
		query += "name = ?, "
		values = append(values, user.Name)
	}

	if user.LastName != "" {
		query += "lastName = ?, "
		values = append(values, user.LastName)
	}

	if user.PlaceOfBirth != "" {
		query += "place_of_birth = ?, "
		values = append(values, user.PlaceOfBirth)
	}

	if user.DateOfBirth != "" {
		query += "date_of_birth = ?, "
		values = append(values, user.DateOfBirth)
	}

	if user.Email != "" {
		query += "email = ?"
		values = append(values, user.Email)
	}

	query += " WHERE id = ?;"
	values = append(values, id)

	_, err := s.DB.Exec(query, values...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Sql) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// Products methods

func (s *Sql) GetProductById(id int) (*domain.Product, error) {
	var product domain.Product
	query := "SELECT * FROM products WHERE id = ?"

	
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.UserId)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *Sql) AddProduct(name string, price float64, userId int) (*domain.Product, error) {
	query := "INSERT INTO products (name, price, user_id) VALUES (?, ?, ?)"

	result, err := s.DB.Exec(query, name, price, userId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product := &domain.Product{
		Id: int(id),
		Name: name,
		Price: price,
		UserId: userId,
	}

	return product, nil
}