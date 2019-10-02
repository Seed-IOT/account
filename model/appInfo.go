package model

// AppInfo 表 用来存储app
type AppInfo struct {
	AppID      uint `gorm:"primary_key"`
	AppSecret  string
	UserDomain string
}
