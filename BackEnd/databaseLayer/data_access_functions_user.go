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

func FindUserByName(name string) (models.User, error) {
	var user models.User
	err := configuration.DB.Where(models.User{Name: name}).First(&user)

	if err != nil {
		return user, error(err.Error)
	}

	return user, nil
}

func FindUserPasswordHashByName(username string) (string, error) {
	user, err := FindUserByName(username)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}
