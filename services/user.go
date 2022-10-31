package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{
		Repository: &repositories.UserRepositoryImpl{},
	}
}

func (cs *UserService) GetAll() []models.User {
	return cs.Repository.GetAll()
}

func (cs *UserService) GetByID(id int) models.User {
	return cs.Repository.GetByID(id)
}

func (cs *UserService) Create(userRequest models.UserRequest) models.User {
	return cs.Repository.Create(userRequest)
}

func (cs *UserService) Login(userRequest models.UserRequest) models.User {
	return cs.Repository.Login(userRequest)
}

func (cs *UserService) Update(id int, userRequest models.UserRequest) models.User {
	return cs.Repository.Update(id, userRequest)
}

func (cs *UserService) Delete(id int) bool {
	return cs.Repository.Delete(id)
}
