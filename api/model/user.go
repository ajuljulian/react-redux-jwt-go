package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name      string
	UserRefer uint
}

type User struct {
	gorm.Model
	Username string `json:"username" xml:"username" form:"username" query:"username"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Roles    []Role `gorm:"foreignKey:UserRefer"`
	Token    string `json:"token,omitempty"`
}
