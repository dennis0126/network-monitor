package service

import (
	"github.com/dennis0126/network-monitor/internal/model"
	"github.com/dennis0126/network-monitor/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (s UserService) ListUsers() ([]model.User, error) {
	return s.userRepo.ListUsers()
}
