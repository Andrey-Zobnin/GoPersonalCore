
package service

import (
    "persona-core/model"
    "persona-core/repository"
)

type UserService interface {
    CreateUser(user *model.User) (*model.User, error)
    GetUser(id string) (*model.User, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
    return &userService{repo: r}
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
    return s.repo.Create(user)
}

func (s *userService) GetUser(id string) (*model.User, error) {
    return s.repo.GetByID(id)
}
