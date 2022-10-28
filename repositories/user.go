package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct{}

func (cr *UserRepositoryImpl) GetAll() []models.User {
	var users []models.User

	database.DB.Preload(clause.Associations).Find(&users)

	return users
}

func (cr *UserRepositoryImpl) GetByID(id int) models.User {
	var user models.User

	database.DB.Preload(clause.Associations).First(&user, id)

	return user
}

func (cr *UserRepositoryImpl) Register(userRequest models.UserRequest) models.User {
	user := userRequest.ToDBForm()

	rec := database.DB.Create(&user)

	rec.Last(&user)

	return user
}

func (cr *UserRepositoryImpl) Login(userRequest models.UserRequest) models.User {
	user := userRequest.ToDBForm()

	database.DB.First(&user, "email = ?", user.Email)

	return user
}

func (cr *UserRepositoryImpl) Update(id int, userRequest models.UserRequest) models.User {
	user := cr.GetByID(id)

	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.RoleID = userRequest.RoleID

	rec := database.DB.Save(&user)

	rec.Last(&user)

	return user
}

func (cr *UserRepositoryImpl) Delete(id int) bool {
	user := cr.GetByID(id)

	rec := database.DB.Delete(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
