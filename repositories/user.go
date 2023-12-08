package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct{}

func (ur *UserRepositoryImpl) GetAll() []models.User {
	var users []models.User

	database.DB.Preload(clause.Associations).Find(&users)

	return users
}

func (ur *UserRepositoryImpl) GetByID(id int) models.User {
	var user models.User

	database.DB.Preload(clause.Associations).First(&user, id)

	return user
}

func (ur *UserRepositoryImpl) Create(userRequest models.UserRequest) models.User {
	user := userRequest.ToDBForm()

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(password)

	rec := database.DB.Create(&user)

	rec.Preload(clause.Associations).Last(&user)

	return user
}

func (ur *UserRepositoryImpl) Login(userRequest models.UserRequest) models.User {
	user := userRequest.ToDBForm()

	database.DB.First(&user, "email = ?", user.Email)

	return user
}

func (ur *UserRepositoryImpl) Update(id int, userRequest models.UserRequest) models.User {
	user := ur.GetByID(id)

	if user.ID == 0 {
		return user
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)

	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = userRequest.Email
	user.Password = string(password)
	user.IsAdmin = userRequest.IsAdmin

	rec := database.DB.Save(&user)

	rec.Preload(clause.Associations).Last(&user)

	return user
}

func (ur *UserRepositoryImpl) Delete(id int) bool {
	user := ur.GetByID(id)

	rec := database.DB.Select("Blogs", "Likes", "Comments").Delete(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
