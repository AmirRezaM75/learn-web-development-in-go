package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type UserService struct {
	DB *sql.DB
}

func (us UserService) Create(email, password string) (*User, error) {
	password = hash(password)

	if len(password) == 0 {
		err := errors.New("something goes wrong with hashing password")
		return nil, err
	}

	user := User{
		Email:    email,
		Password: password,
	}

	row := us.DB.QueryRow(`
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`, email, password)

	err := row.Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func hash(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hashedBytes)
}

func (us UserService) GetUserById(id string) (User, error) {
	var user User

	row := us.DB.QueryRow(`
		SELECT id, email, password FROM users
		WHERE id = $1
	`, id)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	return user, err
}

func (us UserService) getUserByEmail(email string) (User, error) {
	var user User

	row := us.DB.QueryRow(`
		SELECT id, email, password FROM users
		WHERE email = $1
	`, email)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	return user, err
}

func (us UserService) Get(limit int) ([]User, error) {
	rows, err := us.DB.Query(`
		SELECT * FROM users LIMIT $1
	`, limit)

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (us UserService) Login(email, password string) (*User, error) {
	user, err := us.getUserByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}
