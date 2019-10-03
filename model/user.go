package model

import (
	"time"
)

// UserInfo 用户表
type UserInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UID       string `gorm:"primary_key"`
	Username  string
	Nickname  string
	Password  string
	Mobile    string
	Devices   []*Device `gorm:"many2many:user_devices;"`
}

// Device 设备表
type Device struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Sno       string `gorm:"primary_key"`
	Mac       string
	UserInfo  []*UserInfo `gorm:"many2many:user_devices;"`
}
