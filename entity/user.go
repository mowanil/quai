package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(32)" json:"username"`
	Password string `gorm:"column:password;type:varchar(1024)" json:"password"`
}
