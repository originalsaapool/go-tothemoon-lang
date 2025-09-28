package service

import (
	"todorestapi/internal/model"
	"todorestapi/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u model.User) model.User {
	return s.repo.Save(u)
}

func (s *UserService) GetUsers() []model.User {
	return s.repo.GetAll()
}

func (s *UserService) GetUser(id int) *model.User {
	return s.repo.GetByID(id)
}
