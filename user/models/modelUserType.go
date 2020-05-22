package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// UserModel describes a UserModel type
	UserModel struct {
		gorm.Model
		Username string `json:"userName"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
)
