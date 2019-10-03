package web

import (
	"account/constant"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// LoginParames 登录参数
type LoginParames struct {
	Password string `json:"password" xml:"password" binding:"required"`
	Username string `json:"username" xml:"username" binding:"required"`
}

// MobileLoginParames 登录参数
type MobileLoginParames struct {
	Mobile string `json:"mobile" xml:"mobile" binding:"required"`
	Code   string `json:"code" xml:"code" binding:"required"`
}

// GetCodeParam 登录参数
type GetCodeParam struct {
	Mobile   string `json:"mobile" xml:"mobile" binding:"required"`
	CodeType string `json:"codeType" xml:"codeType" binding:"required"`
}

// Login
// @Summary Login
// @Description 账号密码登录
// @ID Login
// @Accept  json
// @Produce  json
// @Param body body web.LoginParames true "用户登录"
// @Success 200 {object} constant.BaseReturn "ok"
// @Router /account/login [post]
func (srv *server) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)

	param := LoginParames{}
	c.Bind(&param)

	username := param.Username
	password := param.Password

	var userInfo, err = srv.service.UserLogin(username, password)
	returnJSON := constant.BaseReturn{}
	if err == nil {
		jsonBytes, _ := json.Marshal(userInfo)

		httpError.Code = constant.SUCCESS
		httpErrorData, _ := httpError.UnmarshalJSON()

		returnJSON.Code = httpErrorData.Code
		returnJSON.Message = httpErrorData.Message

		returnJSON.Data = jsonBytes
	} else {
		httpErrorData, _ := err.UnmarshalJSON()
		returnJSON.Code = httpErrorData.Code
		returnJSON.Message = httpErrorData.Message
	}
	c.JSON(200, returnJSON)
}

// Mobile Login
// @Summary Mobile Login
// @Description 验证码登录
// @ID Mobile Login
// @Accept  json
// @Produce  json
// @Param body body web.MobileLoginParames true "用户登录"
// @Success 200 {object} constant.BaseReturn "ok"
// @Router /account/mobileLogin [post]
func (srv *server) MobileLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)

	param := MobileLoginParames{}

	c.Bind(&param)

	mobile := param.Mobile
	code := param.Code
	var userInfo, err = srv.service.MobileLogin(mobile, code)
	returnJSON := constant.BaseReturn{}
	if err == nil {
		jsonBytes, _ := json.Marshal(userInfo)

		httpError.Code = constant.SUCCESS
		httpErrorData, _ := httpError.UnmarshalJSON()

		returnJSON.Code = httpErrorData.Code
		returnJSON.Message = httpErrorData.Message
		returnJSON.Data = jsonBytes
		c.JSON(200, returnJSON)

	} else {
		httpErrorData, _ := err.UnmarshalJSON()

		returnJSON.Code = httpErrorData.Code
		returnJSON.Message = httpErrorData.Message
		c.JSON(200, returnJSON)
	}
}

// Get Code
// @Summary Get Code
// @Description 获取验证码
// @ID Get Code
// @Accept  json
// @Produce  json
// @Param body body web.GetCodeParam true "获取验证码"
// @Success 200 {object} constant.BaseReturn "ok"
// @Router /account/getCode [post]
func (srv *server) GetCode(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	param := GetCodeParam{}

	c.Bind(&param)

	codeType := param.CodeType
	mobile := param.Mobile
	var codeData, err = srv.service.GetCode(mobile, codeType)
	returnJSON := constant.BaseReturn{}
	if err == nil {
		jsonBytes, jsonErr := json.Marshal(codeData)
		if jsonErr == nil {

			httpError := new(constant.Error)
			httpError.Code = constant.SUCCESS

			returnJSON.Code = httpError.Code
			returnJSON.Message = httpError.Message
			returnJSON.Data = jsonBytes
			c.JSON(200, returnJSON)
			return
		}
	}

	errData, _ := err.UnmarshalJSON()
	returnJSON.Code = errData.Code
	returnJSON.Message = errData.Message

	c.JSON(200, returnJSON)
}
