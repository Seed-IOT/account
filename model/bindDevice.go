package model

import (
	"account/constant"
	"fmt"
	"strings"
)

type bindDevicesParams struct {
	Snos []string
	UID  string
}

// BindDevice 绑定设备
func (srv *Service) BindDevice(snos []string, uid string) ([]Device, *constant.Error) {
	// 检验用户是否合法
	errorData := new(constant.Error)
	errorData.Code = constant.UNKNOWN_ERROR

	userInfo := UserInfo{}
	srv.DB.First(&userInfo, "uid = ?", uid)
	// 判断是否存在
	if !strings.EqualFold(userInfo.UID, uid) {
		// 不存在
		errorData := new(constant.Error)
		errorData.Code = constant.UNKNOWN_ERROR
		return nil, errorData
	}

	devices := []Device{}
	srv.DB.Where("sno in (?)", snos).Find(&devices)
	// 查数据库 是否存在
	deviceLen := len(devices)
	if deviceLen == 0 {
		// 不存在
		errorData := new(constant.Error)
		errorData.Code = constant.DEVICE_NOT_FOUND
		return []Device{}, errorData
	}

	// 创建关联
	srv.DB.Model(&userInfo).Association("Devices").Append(&devices)
	fmt.Println(userInfo)
	return devices, nil
}
