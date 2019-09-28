package model

import (
	"account/constant"
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// HmacSampleSecret 密钥
var HmacSampleSecret = []byte("KYLE_WANG")

// UserLoginReturn 返回的用户信息
type UserLoginReturn struct {
	Token string
	UID   string
}

// UserInfo 用户表
type UserInfo struct {
	UID      string `gorm:"primary_key"`
	Username string
	Nickname string
	Password string
	Mobile   string
}

func createToken(UID string) (string, error) {
	jti, err := uuid.NewV4()
	if err != nil {
		return "", errors.New("ERROR")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  string(UID),
		Subject:   "APP",
		Issuer:    "MAIN",
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiresAt: time.Now().AddDate(0, 0, 30).UTC().Unix(),
		Id:        jti.String(),
		NotBefore: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(HmacSampleSecret)
	if err != nil {
		return "", errors.New("ERROR")
	}
	return tokenString, nil
}

// UserLogin 登录方法
func (srv *Service) UserLogin(username string, password string) (UserLoginReturn, error) {
	// 查询数据库
	var userinfo UserInfo
	srv.DB.First(&userinfo, "username = ?", username)

	if strings.EqualFold(userinfo.Username, username) {
		passwordMd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if strings.EqualFold(userinfo.Password, passwordMd5) {
			// 密码相同，签发token
			tokenString, err := createToken(string(userinfo.UID))

			if err != nil {
				return UserLoginReturn{}, errors.New("ERROR")
			}
			data := UserLoginReturn{Token: tokenString, UID: userinfo.UID}
			return data, nil

		}
	}
	return UserLoginReturn{}, errors.New("ERROR")
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// MobileLogin 手机登录
func (srv *Service) MobileLogin(mobile string, code string) (UserLoginReturn, *constant.Error) {
	// 查询redis
	reply, err := redis.String(srv.redisConn.Do("GET", "MobileCode:"+mobile))
	log.Println("reply", reply == "", err)
	if err == nil {
		if reply == code {
			// 验证码失效
			srv.redisConn.Do("DEL", "MobileCode:"+mobile)

			// 验证码正确，用手机号去查询用户

			var userinfo UserInfo
			srv.DB.First(&userinfo, "mobile = ?", mobile)
			// userinfo 存在则返回 不存在创建token

			if userinfo.UID == "" {
				// 不存在，创建一个
				var userinfo UserInfo

				userinfo.Mobile = mobile

				jti, _ := uuid.NewV4()
				userinfo.UID = jti.String()

				srv.DB.Create(&userinfo)
			}

			tokenString, _ := createToken(string(userinfo.UID))
			data := UserLoginReturn{Token: tokenString, UID: userinfo.UID}
			return data, nil
		}
	}
	// 验证码错误
	errData := new(constant.Error)
	errData.Code = constant.VERIFICATION_CODE
	return UserLoginReturn{}, errData
}
