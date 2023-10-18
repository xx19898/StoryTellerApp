package configuration

import (
	"StoryTellerAppBackend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ConfigureDatabaseForTest() {
	ConnectTestDb(&gorm.Config{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.User{})
	ResetEverythingAndPopulateRoles()
}

func ResetEverythingElseExceptRoles() {
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.User{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.Story{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.Comment{})
}

func ResetEverythingAndPopulateRoles() {
	DB.Exec("DELETE FROM user_roles;")
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.User{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Select(clause.Associations).Delete(&models.Role{})
	DB.CreateInBatches([]models.Role{{Name: "User"}, {Name: "Admin"}}, 2)
}
