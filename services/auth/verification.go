package auth

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
)

func Verify(userID int) error {
	return db.GetDb().Model(&models.User{}).Where("id = ?", userID).Update("is_verified", "true").Error
}
