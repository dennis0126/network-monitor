package service

import (
	"errors"
	"github.com/dennis0126/network-monitor/internal/model"
	"github.com/dennis0126/network-monitor/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrAuthFailed = errors.New("authentication failed")

type AuthService struct {
	sessionRepo repository.SessionRepository
	userService UserService
}

func NewAuthService(sessionRepo repository.SessionRepository, userService UserService) AuthService {
	return AuthService{sessionRepo: sessionRepo, userService: userService}
}

type LoginParam struct {
	Name      string
	Password  string
	IpAddress string
	UserAgent string
}

func (s AuthService) Login(param LoginParam) (model.Session, error) {
	user, err := s.userService.GetUserByName(param.Name)
	if err != nil {
		return model.Session{}, err
	}
	if user == nil {
		return model.Session{}, ErrAuthFailed
	}

	if !checkPasswordHash(param.Password, user.PasswordHash) {
		return model.Session{}, ErrAuthFailed
	}

	session, err := s.sessionRepo.CreateSession(user.ID, param.IpAddress, param.UserAgent)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s AuthService) Logout(sessionId string) error {
	err := s.sessionRepo.DeleteSessionById(sessionId)
	if err != nil {
		return err
	}
	return nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
