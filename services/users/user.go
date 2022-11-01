package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserService{
		Repository: repository,
	}
}

func (us *UserService) GetAll() []models.User {
	return us.Repository.GetAll()
}

func (us *UserService) GetByID(id int) models.User {
	return us.Repository.GetByID(id)
}

func (us *UserService) Create(userRequest models.UserRequest) models.User {
	return us.Repository.Create(userRequest)
}

func (us *UserService) Login(userRequest models.UserRequest) models.User {
	return us.Repository.Login(userRequest)
}

func (us *UserService) Update(id int, userRequest models.UserRequest) models.User {
	return us.Repository.Update(id, userRequest)
}

func (us *UserService) Delete(id int) bool {
	return us.Repository.Delete(id)
}
