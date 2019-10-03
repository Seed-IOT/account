package web

import (
	"account/constant"
	"account/model"

	"github.com/gin-gonic/gin"
)

// BindDevicesParames 登录参数
type BindDevicesParames struct {
	UID  string   `json:"uid" xml:"uid" binding:"required"`
	Snos []string `json:"snos" xml:"snos" binding:"required"`
}

type bindReturn struct {
	constant.BaseReturn
	Data []model.Device `json:"data"`
}

// BindDevices
// @Summary BindDevices
// @Description 绑定设备
// @ID BindDevices
// @Accept  json
// @Produce  json
// @Param body body web.BindDevicesParames true "用户登录"
// @Success 200 {object} constant.BaseReturn "ok"
// @Router /account/bindings [post]
func (srv *server) BindDevices(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)
	httpError.Code = constant.SUCCESS

	param := BindDevicesParames{}
	c.Bind(&param)

	snos := param.Snos
	uid := param.UID

	var data, err = srv.service.BindDevice(snos, uid)
	if err == nil {
		// 绑定成功
	} else {
		// 绑定失败
	}

	returnData := bindReturn{}
	returnData.Code = constant.SUCCESS
	returnData.Data = data
	c.JSON(200, returnData)
}
