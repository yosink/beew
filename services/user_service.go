package services

import (
	"beew/models"
	"beew/repositories"
)

type IUserService interface {
	GetUserByPhone(string) (models.User, error)
	ExistsByID(int) (bool, error)
}

type UserService struct {
	repo repositories.IUserRepository
}

func (u UserService) ExistsByID(id int) (bool, error) {
	return u.repo.ExistsByID(id)
}

func (u UserService) GetUserByPhone(phone string) (models.User, error) {
	return u.repo.GetByPhone(phone)
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &UserService{repo: repo}
}
