package model

// TypeApp 表 用来存储app
type TypeApp struct {
	AppID      uint `gorm:"primary_key"`
	AppSecret  string
	UserDomain string
}
