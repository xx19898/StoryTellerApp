package helpers

import "StoryTellerAppBackend/models"

func RolesToString(roles []models.Role) []string {
	stringArray := []string{}

	for i := 0; i < len(roles); i++ {
		stringArray = append(stringArray, roles[i].Name)
	}

	return stringArray
}
