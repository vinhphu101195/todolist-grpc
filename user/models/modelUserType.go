package models

import "time"

type (
	// UserModel describes a UserModel type
	UserModel struct {
		Userid   uint   `json:"userid" gorm:"primary_key;AUTO_INCREMENT"`
		Username string `json:"userName"`
		Password string `json:"password"`
		Email    string `json:"email"`

		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
	}
)
