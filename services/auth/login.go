package auth

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
	"BI.ZONE_test/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

func LoginUser(username, password string) (models.LoginResponse, error) {
	account := &models.User{}
	err := db.GetDb().Table("users").Where("username = ?", username).First(account).Error

	if err != nil {
		log.Warn(err)
		if err == gorm.ErrRecordNotFound {
			return models.LoginResponse{}, fmt.Errorf("Invalid login credentials ")
		}
		return models.LoginResponse{}, err
	}

	if !utils.CheckPasswordHash(password, account.Password) { // Password does not match!
		return models.LoginResponse{}, fmt.Errorf("Invalid login credentials ")
	}

	// check for admin verification
	if account.IsVerified != true {
		return models.LoginResponse{}, fmt.Errorf("Sorry, you are not verified ")
	}

	// create JWT token
	tk := &models.JWT{UserID: account.ID, Role: account.UserRoleID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	return models.LoginResponse{Token: tokenString}, nil
}
