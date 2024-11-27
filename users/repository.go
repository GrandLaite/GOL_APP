package users

import (
	"database/sql"
	"errors"
)

type UserRepository interface {
	Create(user User) (int, error)
	GetByID(userID int) (User, error)
	Update(userID int, user User) error
	Delete(userID int) error
	GetByUsername(username string) (User, error)
	GetPremiumStatus(userID int) (bool, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) Create(user User) (int, error) {
	var userID int
	query := "INSERT INTO Users (username, password_hash, is_premium) VALUES ($1, $2, $3) RETURNING id"
	err := ur.DB.QueryRow(query, user.Username, user.Password, user.IsPremium).Scan(&userID)
	return userID, err
}

func (ur *userRepository) GetByID(userID int) (User, error) {
	var user User
	query := "SELECT id, username, password_hash, is_premium FROM Users WHERE id = $1"
	err := ur.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password, &user.IsPremium)
	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	}
	return user, err
}

func (ur *userRepository) Update(userID int, user User) error {
	query := "UPDATE Users SET username = $1, is_premium = $2 WHERE id = $3"
	_, err := ur.DB.Exec(query, user.Username, user.IsPremium, userID)
	return err
}

func (ur *userRepository) Delete(userID int) error {
	query := "DELETE FROM Users WHERE id = $1"
	_, err := ur.DB.Exec(query, userID)
	return err
}

func (ur *userRepository) GetByUsername(username string) (User, error) {
	var user User
	query := "SELECT id, username, password_hash, is_premium FROM Users WHERE username = $1"
	err := ur.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.IsPremium)
	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	}
	return user, err
}

func (ur *userRepository) GetPremiumStatus(userID int) (bool, error) {
	var isPremium bool
	query := "SELECT is_premium FROM Users WHERE id = $1"
	err := ur.DB.QueryRow(query, userID).Scan(&isPremium)
	return isPremium, err
}
