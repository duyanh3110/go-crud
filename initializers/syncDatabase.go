package initializers

import "go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
