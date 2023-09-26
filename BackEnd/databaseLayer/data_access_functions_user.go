package databaselayer

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
)

func CreateNewUser(name string, password string, email string, roles []models.Role) models.User {
	newUser := models.User{Name: name, Password: password, Email: email, Roles: roles}
	configuration.DB.Create(&newUser)
	return newUser
}

func FindUserByName(name string) models.User {
	var user models.User
	configuration.DB.Where(models.User{Name: name}).First(&user)
	return user
}

func FindUserPasswordHashByName(username string) (string, error) {
	user := FindUserByName(username)
	return user.Password, nil
}

func FindUserById(id int) models.User {
	var user models.User
	configuration.DB.Where(models.User{ID: uint(id)}).First(&user)
	return user
}
