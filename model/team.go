package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(50);not null"`
	Pwd   string  `gorm:"type:varchar(255);not null"`
	Logo  string  `gorm:"type:varchar(50)"` //logo 的地址
	Score float64 `gorm:"not null"`
}
