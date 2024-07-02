package service

import (
	"github.com/dennis0126/network-monitor/internal/model"
	"github.com/dennis0126/network-monitor/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (s UserService) CreateUser(name string, password string) (model.User, error) {
	id := uuid.New().String()
	passwordHash, err := hashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	return s.userRepo.CreateUser(model.User{ID: id, Name: name, PasswordHash: passwordHash, CreatedAt: time.Now(), UpdatedAt: time.Now()})
}

func (s UserService) ListUsers() ([]model.User, error) {
	return s.userRepo.ListUsers()
}

func (s UserService) GetUserByName(name string) (*model.User, error) {
	return s.userRepo.GetUserByName(name)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
