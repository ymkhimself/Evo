package model

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)"`
	Desc        string `gorm:"type:varchar(255)"` // 题目描述
	AutoRefresh bool   //是否自动刷新flag
	Command     string `gorm:"type:varchar(255)"` //刷新flag时使用的shell命令
	Visible     bool   //是否可见
}
