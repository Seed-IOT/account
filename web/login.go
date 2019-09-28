package web

import (
	"account/constant"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (srv *server) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpError := new(constant.Error)

	username := c.PostForm("username")
	password := c.PostForm("password")
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
		returnJSON.Code = 500
		returnJSON.Message = "ERROR"
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
