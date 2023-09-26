package databaselayer

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
)

func FindRoleByName(name string) models.Role {
	var role models.Role
	configuration.DB.Where(models.Role{Name: name})
	return role
}
