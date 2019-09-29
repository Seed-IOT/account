package constant

import (
	"encoding/json"
	"fmt"
)

// Error x
type Error struct {
	Code    int
	Message string
}

var SUCCESS = 0
var UNKNOWN_ERROR = 8000
var VERFIICATION_CODE_EXISTS = 8001
var VERIFICATION_CODE = 8002
var USER_OR_PASSWORD_ERROR = 8003

// ErrorCodes x
var ErrorCodes = map[int]string{
	SUCCESS:                  "成功",
	UNKNOWN_ERROR:            "未知错误",
	VERFIICATION_CODE_EXISTS: "验证码已经存在",
	VERIFICATION_CODE:        "验证码错误",
	USER_OR_PASSWORD_ERROR:   "用户名或密码错误",
}

func (e Error) String() string {
	return ErrorCodes[e.Code]
}

// UnmarshalJSON x
func (e *Error) UnmarshalJSON() (Error, error) {
	errData := Error{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"code": %d, "message": "%s"}`, e.Code, e.String())), &errData)
	return errData, err
}
