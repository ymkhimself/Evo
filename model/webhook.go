package model

import "gorm.io/gorm"

type Webhook struct {
	gorm.Model
	Url     string `gorm:"type:varchar(255)"`
	Type    string `gorm:"type:varchar(30)"`
	Retry   uint   //重试次数
	Timeout uint   //时限
}
