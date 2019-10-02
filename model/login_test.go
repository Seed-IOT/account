package model

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateToken(t *testing.T) {
	uid := "asdasdasdas"
	data, err := CreateToken(uid)
	// 用解析方法解出来看是不是正确的
	tokenData, err := ParseToken(data)
	if err != nil {
		t.FailNow()
	}
	if !strings.EqualFold(tokenData["aud"].(string), uid) {
		// 不通过
		t.FailNow()
	}
	fmt.Println("TestCreateToken Success")
}
