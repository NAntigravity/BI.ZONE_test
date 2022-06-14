package auth

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
	"BI.ZONE_test/utils"
)

func RegisterUser(credentials models.RegisterCredentials) error {
	hashedPassword, err := utils.HashPassword(credentials.Password)
	if err != nil {
		return err
	}
	user := models.User{
		Username:   credentials.Username,
		Password:   hashedPassword,
		UserRoleID: credentials.Role,
		IsVerified: false,
	}
	err = db.GetDb().Create(&user).Error
	return err
}
