package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"unique_index;not null;"`
	Password   string `json:"-" gorm:"not null;"`
	UserRoleID uint   `json:"user_role_id"`
	IsVerified bool   `json:"isVerified" gorm:"not null;default:false;"`
}
