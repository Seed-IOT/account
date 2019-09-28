package model

import (
	"account/constant"
)

// VerificationCodeReturnData s
type VerificationCodeReturnData struct {
	Mobile string
	Code   string
}

// func typeof(v interface{}) string {
// 	return reflect.TypeOf(v).String()
// }

// GetCode 手机登录
func (srv *Service) GetCode(mobile string, codeType string) (VerificationCodeReturnData, *constant.Error) {
	// 发送验证码
	// apiKey := srv.config.Release.YunPianAppKey
	errData := new(constant.Error)
	// 检查是否超过60秒
	time, err := srv.redisConn.Do("TTL", "MobileCode:"+string(mobile))
	if err == nil {
		v, ok := time.(int64)
		if ok {
			if v == -2 {
				data := VerificationCodeReturnData{}
				data.Mobile = mobile
				data.Code = "1234"

				// 写入redis 设置有效期
				srv.redisConn.Do("SETEX", "MobileCode:"+string(mobile), 60, data.Code)
				return data, nil
			}
			// 存在提示请不要频繁发送验证码
			errData.Code = constant.VERFIICATION_CODE_EXISTS
			return VerificationCodeReturnData{}, errData
		}
	}
	errData.Code = constant.UNKNOWN_ERROR
	return VerificationCodeReturnData{}, errData
}
