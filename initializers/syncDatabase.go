package initializers

import "github.com/sureshchandak1/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
