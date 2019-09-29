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

// Login
// @Summary Login
// @Description 账号密码登录
// @ID Login
// @Accept  json
// @Produce  json
// @Param body body web.LoginParames true "用户登录"
// @Router /account/login [post]
func (srv *server) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)

	param := &LoginParames{}
	c.Bind(&param)

	username := param.Username
	password := param.Password

	var userInfo, err = srv.service.UserLogin(username, password)
	returnJSON := baseReturn{}
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

func (srv *server) MobileLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)

	mobile := c.PostForm("mobile")
	code := c.PostForm("code")
	var userInfo, err = srv.service.MobileLogin(mobile, code)
	returnJSON := baseReturn{}
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

func (srv *server) GetCode(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	codeType := c.PostForm("codeType")
	mobile := c.PostForm("mobile")
	var codeData, err = srv.service.GetCode(mobile, codeType)
	returnJSON := baseReturn{}
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

type baseReturn struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
