package configuration

import (
	"StoryTellerAppBackend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ConfigureDatabaseForTest() {
	ConnectTestDb(&gorm.Config{})
	ResetEverythingAndPopulateRoles()
}

func ResetEverythingElseExceptRoles() {
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.User{Name: "testuser"})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.Story{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.Comment{})
}

func ResetEverythingAndPopulateRoles() {
	DB.Exec("DELETE FROM user_roles;")
	DB.Exec("DELETE FROM comments")
	DB.Exec("DELETE FROM stories;")
	DB.Exec("DELETE FROM roles;")
	DB.Exec("DELETE FROM users;")
	DB.CreateInBatches([]models.Role{{Name: "User"}, {Name: "Admin"}}, 2)
}
