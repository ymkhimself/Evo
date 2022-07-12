package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Pwd  string `gorm:"type:varchar(255);not null"`
}
