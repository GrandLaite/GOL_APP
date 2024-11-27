package users

import (
	"errors"
	"gol_messenger/auth"
)

type UserService interface {
	AuthenticateUser(username, password string) (int, bool, error)
	GetUser(userID int) (User, error)
	UpdateUser(userID int, updatedUser User) error
	DeleteUser(userID int) error
	RegisterUser(user User) (int, error)
	GetUserPremiumStatus(userID int) (bool, error)
	GenerateToken(userID int, isPremium bool) (string, error)
}

type userService struct {
	UserRepository UserRepository
	TokenService   auth.TokenService
}

func NewUserService(userRepository UserRepository, tokenService auth.TokenService) UserService {
	return &userService{
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
}

func (us *userService) RegisterUser(user User) (int, error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword

	return us.UserRepository.Create(user)
}

func (us *userService) AuthenticateUser(username, password string) (int, bool, error) {
	user, err := us.UserRepository.GetByUsername(username)
	if err != nil || !auth.CheckPasswordHash(password, user.Password) {
		return 0, false, errors.New("неверное имя пользователя или пароль")
	}

	return user.ID, user.IsPremium, nil
}

func (us *userService) GenerateToken(userID int, isPremium bool) (string, error) {
	return us.TokenService.GenerateToken(userID, isPremium)
}

func (us *userService) GetUser(userID int) (User, error) {
	return us.UserRepository.GetByID(userID)
}

func (us *userService) UpdateUser(userID int, updatedUser User) error {
	return us.UserRepository.Update(userID, updatedUser)
}

func (us *userService) DeleteUser(userID int) error {
	return us.UserRepository.Delete(userID)
}

func (us *userService) GetUserPremiumStatus(userID int) (bool, error) {
	return us.UserRepository.GetPremiumStatus(userID)
}
