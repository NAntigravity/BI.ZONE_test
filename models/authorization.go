package models

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	Username string `json:"username" gorm:"unique_index;not null;"`
	Role     uint   `json:"role"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}
